package script

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

type UnixSockClient struct {
	Sock       string
	Transport  *http.Transport
	HttpClient *http.Client
}

func (p *UnixSockClient) GetDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", p.Sock)
}

func (p *UnixSockClient) Get(urlStr string) (resp *http.Response, err error) {
	return p.HttpClient.Get(urlStr)
}

func (p *Script) _getUnixSockPointerFromUserData(L *lua.LState, lu *lua.LUserData) *UnixSockClient {
	if nil == lu || nil == lu.Value {
		return nil
	}

	unixSockClient, ok := lu.Value.(*UnixSockClient)
	if false == ok || nil == unixSockClient {
		return nil
	}

	return unixSockClient
}

func (p *Script) NewUnixSockClient(L *lua.LState) int {
	var ret UnixSockClient

	ret.Sock = L.ToString(1)
	ret.Transport = &http.Transport{Dial: ret.GetDial}
	ret.HttpClient = &http.Client{Transport: ret.Transport}

	luaUnixSockClient := L.NewUserData()
	luaUnixSockClient.Value = &ret
	L.Push(luaUnixSockClient)
	return 1
}

func (p *Script) UnixSockGet(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	unixSockClient := p._getUnixSockPointerFromUserData(L, L.ToUserData(1))
	urlStr := L.ToString(2)

	resp, err := unixSockClient.Get(urlStr)
	if nil != err {
		log.Println(err)
		return 0
	}
	defer resp.Body.Close()

	ret, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return 0
	}

	L.Push(lua.LString(ret))
	return 1
}
