package iron

import (
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"

	"golang.org/x/net/http2"
)

type Server struct {
	ListenStr string
	ServeType string
	BaseDir   string
	ViewDir   string
	RunMode   string

	NotFoundHandler   Handler
	isTMPLAutoRefresh bool
	views             map[string]*View
	httpMux           *ServeMux
	Hooker            Hook

	BaseUrl              string
	StaticBaseUrl        string
	StaticBasePath       string
	StaticUploadBaseUrl  string
	StaticUploadBasePath string

	ImgExts []string

	ConfFiles []ConfFile

	isClosedAfterHandle bool

	LogPath string
	Log     *os.File

	httpServer *http.Server
}

func NewServer() *Server {
	p := &Server{}
	p.views = make(map[string]*View)
	p.httpMux = p.NewServeMux()

	p.ImgExts = []string{"jpeg", "gif", "png", "jpg"}

	initEncoder()

	p.NotFoundHandler = HandlerFunc(Handle404)
	p.Hooker.BeforeServeRequest = make([]HookBeforeServeRequest, 0)
	p.Hooker.BeforeHttpHandles = make([]HookBeforeHttpHandle, 0)
	p.Hooker.ErrorRecovers = make([]HookErrorRecover, 0)
	p.Hooker.AfterHttpHandles = make([]HookAfterHttpHandle, 0)
	p.Hooker.UrlRewrite = make([]HookUrlRewrite, 0)

	p.httpServer = &http.Server{
		Handler:        p.httpMux,
		ReadTimeout:    90 * time.Second,
		WriteTimeout:   90 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http2.ConfigureServer(p.httpServer, &http2.Server{})

	return p
}

func (p *Server) LoadConfFile(confFile ConfFile, prefix string) error {
	var ok bool
	var confStr string

	p.BaseDir, _ = os.Getwd()

	confStr, ok = confFile.Get(prefix, "site_view_dir")
	if ok {
		p.ViewDir = confStr
	} else {
		p.ViewDir = p.BaseDir + "/view"
	}

	p.ConfFiles = append(p.ConfFiles, confFile)

	confStr, ok = confFile.Get(prefix, "listen_str")
	if ok {
		p.ListenStr = confStr
	}

	confStr, ok = confFile.Get(prefix, "serve_type")
	if ok {
		p.ServeType = confStr
	} else {
		p.ServeType = "server"
	}

	confStr, ok = confFile.Get(prefix, "runmode")
	if ok {
		p.RunMode = confStr
	}

	switch p.RunMode {
	case "dev":
		p.isTMPLAutoRefresh = true
	case "test":
		p.isTMPLAutoRefresh = true
	case "proc":
		p.isTMPLAutoRefresh = false
	}

	confStr, ok = confFile.Get(prefix, "site_baseurl")
	if ok {
		p.BaseUrl = confStr
	}

	confStr, ok = confFile.Get(prefix, "site_static_baseurl")
	if ok {
		p.StaticBaseUrl = confStr
	}

	confStr, ok = confFile.Get(prefix, "site_static_basepath")
	if ok {
		p.StaticBasePath = confStr
	}

	confStr, ok = confFile.Get(prefix, "site_static_upload_baseurl")
	if ok {
		p.StaticUploadBaseUrl = confStr
	}

	confStr, ok = confFile.Get(prefix, "site_static_upload_basepath")
	if ok {
		p.StaticUploadBasePath = confStr
	}

	confStr, ok = confFile.Get(prefix, "log")
	if ok {
		if nil != p.Log {
			p.Log.Close()
		}
		p.LogPath = confStr
		p.Log, _ = os.OpenFile(p.LogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0755)
		log.SetOutput(p.Log)
	}

	return nil
}

func (p *Server) LoadConfFileValue(k1, k2 string) string {
	var ret, tmp string
	var ok bool

	for _, f := range p.ConfFiles {
		tmp, ok = f.Get(k1, k2)
		if ok {
			ret = tmp
		}
	}
	return ret
}

func (p *Server) SetCloseAfterHandle() {
	p.isClosedAfterHandle = true
}

func (p *Server) Run() error {
	log.Println("Server started")

	p.isClosedAfterHandle = false

	switch p.ServeType {
	case "fcgi":
		listener, err := net.Listen("tcp", p.ListenStr)
		if nil != err {
			return err
		}

		fcgi.Serve(listener, p.httpMux)

	case "server":
		p.httpServer.Addr = p.ListenStr
		//err := p.httpServer.ListenAndServeTLS("./cert.pem", "./key.pem")
		err := p.httpServer.ListenAndServe()
		if nil != err {
			return err
		}

	default:
		log.Println("请选择 fcgi 或 server 模式")
	}

	log.Println("Server closed.")

	return nil

}
