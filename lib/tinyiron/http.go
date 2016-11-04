package tinyiron

import (
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"strings"
	"sync"
)

type ServeMux struct {
	server *Server
	mu     sync.RWMutex
	m      map[string]muxEntry
	hosts  bool // whether any patterns contain hostnames
}

type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}

type Handler interface {
	ServeHTTP(*Server, http.ResponseWriter, *http.Request)
}

type HookBeforeServeRequest func(*Request) bool
type HookBeforeHttpHandle func(*Request) bool
type HookErrorRecover func(*Request, interface{}) bool
type HookAfterHttpHandle func(*Request) bool
type HookUrlRewrite func(*Request) bool

type Hook struct {
	BeforeServeRequest []HookBeforeServeRequest
	BeforeHttpHandles  []HookBeforeHttpHandle
	ErrorRecovers      []HookErrorRecover
	AfterHttpHandles   []HookAfterHttpHandle
	UrlRewrite         []HookUrlRewrite
}

// Redirect to a fixed URL
type redirectHandler struct {
	url  string
	code int
}

func (rh *redirectHandler) ServeHTTP(server *Server, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, rh.url, rh.code)
}

// RedirectHandler returns a request handler that redirects
// each request it receives to the given url using the given
// status code.
//
// The provided code should be in the 3xx range and is usually
// StatusMovedPermanently, StatusFound or StatusSeeOther.
func RedirectHandler(url string, code int) Handler {
	return &redirectHandler{url, code}
}

func Handle404(ir *Request) {
	ir.ApiOutput(nil, -1, "command not found")
}

// Does path match pattern?
func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		// should not happen
		return false
	}
	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}

// Return the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

// NewServeMux allocates and returns a new ServeMux.
func (p *Server) NewServeMux() *ServeMux {
	ret := &ServeMux{server: p, m: make(map[string]muxEntry)}
	return ret
}

func (p *Server) Router(path string, handler func(*Request)) {
	p.httpMux.HandleFunc(path, handler)
}

// Find a handler on a handler map given a path string
// Most-specific (longest) pattern wins
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}

func (mux *ServeMux) Handler(r *http.Request) (h Handler, pattern string) {
	if r.Method != "CONNECT" {
		if p := cleanPath(r.URL.Path); p != r.URL.Path {
			_, pattern = mux.handler(r.Host, p)
			url := *r.URL
			url.Path = p
			return RedirectHandler(url.String(), http.StatusMovedPermanently), pattern
		}
	}

	return mux.handler(r.Host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if mux.hosts {
		h, pattern = mux.match(host + path)
	}
	if h == nil {
		h, pattern = mux.match(path)
	}
	if h == nil {
		h, pattern = mux.server.NotFoundHandler, ""
	}
	return
}

func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseForm()
	r.ParseMultipartForm(32 << 20)

	var ir *Request = &Request{server: mux.server}
	ir.Init(w, r)
	for _, h := range mux.server.Hooker.UrlRewrite {
		if !h(ir) {
			return
		}
	}

	h, _ := mux.Handler(r)
	h.ServeHTTP(mux.server, w, r)
}

func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if mux.m[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

	if pattern[0] != '/' {
		mux.hosts = true
	}

	// Helpful behavior:
	// If pattern is /tree/, insert an implicit permanent redirect for /tree.
	// It can be overridden by an explicit registration.
	n := len(pattern)
	if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
		// If pattern contains a host name, strip it and use remaining
		// path for redirect.
		path := pattern
		if pattern[0] != '/' {
			// In pattern, at least the last character is a '/', so
			// strings.Index can't be -1.
			path = pattern[strings.Index(pattern, "/"):]
		}
		mux.m[pattern[0:n-1]] = muxEntry{h: RedirectHandler(path, http.StatusMovedPermanently), pattern: pattern}
	}
}

type HandlerFunc func(*Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(server *Server, w http.ResponseWriter, r *http.Request) {
	var ok bool
	var ir *Request = &Request{server: server}
	ir.Init(w, r)

	for _, h := range server.Hooker.BeforeServeRequest {
		if !h(ir) {
			return
		}
	}

	ir.ViewData = make(map[string]interface{})
	ir.ViewData["StaticBaseUrl"] = server.StaticBaseUrl

	_, ok = server.httpMux.m[ir.R.URL.Path]
	if !ok {
		if len(ir.R.URL.Path) > 8 && "/static/" == ir.R.URL.Path[0:8] {
			file := path.Join(server.StaticBasePath, ir.R.URL.Path[8:])
			if FileExists(file) {
				http.ServeFile(ir.W, ir.R, file)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		Handle404(ir)
		return
	}

	ir.V = make(map[string]interface{})

	for _, h := range server.Hooker.BeforeHttpHandles {
		if !h(ir) {
			return
		}
	}

	defer func() {
		err := recover()
		if nil == err {
			return
		}
		log.Println(string(debug.Stack()))
		log.Println(err)
		for _, h := range server.Hooker.ErrorRecovers {
			h(ir, err)
		}
	}()
	f(ir)

	for _, h := range server.Hooker.AfterHttpHandles {
		if !h(ir) {
			return
		}
	}

	if server.isClosedAfterHandle {
		log.Println("Server closed.")
		os.Exit(0)
	}
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(*Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}
