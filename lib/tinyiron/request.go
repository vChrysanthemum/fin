package tinyiron

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	server   *Server
	RemoteIp string
	W        http.ResponseWriter
	R        *http.Request
	V        map[string]interface{}
	ViewData map[string]interface{}
	Now      int64
}

func (p *Request) Init(w http.ResponseWriter, r *http.Request) {
	p.W = w
	p.R = r
	p.RemoteIp = strings.Split(r.RemoteAddr, ":")[0]
	p.Now = time.Now().Local().Unix()
}

func (p *Request) Redirect(url string) {
	http.Redirect(p.W, p.R, url, 302)
}

func (p *Request) ApiOutput(data interface{}, errno int, errmsg string) {
	p.W.Header().Add("Server", "tinyiron")
	p.W.Header().Add("Content-Type", "application/json")
	ret := map[string]interface{}{"data": data, "errno": errno, "errmsg": errmsg}
	res, _ := json.Marshal(ret)
	p.W.Write(res)
}

func (p *Request) Render(path string) {
	p.server.Render(path, p.W, p.R, p.ViewData)
}

func (p *Request) MustFormString(key string, defaultRet string) (ret string) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return defaultRet
	}
	v[0] = strings.TrimSpace(v[0])
	return v[0]
}

func (p *Request) FormString(key string) (ret string, err error) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return "", errors.New(fmt.Sprintf("%v not exists", key))
	}
	v[0] = strings.TrimSpace(v[0])
	return v[0], nil
}

func (p *Request) FormFloat64(key string) (ret float64, err error) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return -1, errors.New(fmt.Sprintf("%v not exists", key))
	}

	ret, err = strconv.ParseFloat(v[0], 64)
	if nil != err {
		return -1, errors.New(fmt.Sprintf("%v not int", key))
	}
	return ret, nil
}

func (p *Request) FormInt(key string) (ret int, err error) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return -1, errors.New(fmt.Sprintf("%v not exists", key))
	}

	_v, err := strconv.ParseInt(v[0], 10, 64)
	ret = int(_v)
	if nil != err {
		return -1, err
	}
	return ret, nil
}

func (p *Request) MustFormInt(key string, defaultRet int) (ret int) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return defaultRet
	}

	_v, err := strconv.ParseInt(v[0], 10, 64)
	ret = int(_v)
	if nil != err {
		return defaultRet
	}
	return ret
}

func (p *Request) MustFormInt64(key string, defaultRet int64) (ret int64) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return defaultRet
	}

	_v, err := strconv.ParseInt(v[0], 10, 64)
	ret = int64(_v)
	if nil != err {
		return defaultRet
	}
	return ret
}

func (p *Request) FormInt64(key string) (ret int64, err error) {
	v := p.R.Form[key]
	if 0 == len(v) {
		return -1, errors.New(fmt.Sprintf("%v not exists", key))
	}

	_v, err := strconv.ParseInt(v[0], 10, 64)
	ret = int64(_v)
	if nil != err {
		return -1, errors.New(fmt.Sprintf("%v not int", key))
	}
	return ret, nil
}

func (p *Request) SaveFile(key, newFileRelativePath string) (err error) {
	file, _, err := p.R.FormFile(key)
	if err != nil {
		return err
	}
	defer file.Close()
	f, err := os.Create(fmt.Sprintf("%s/%s", p.server.StaticUploadBasePath, newFileRelativePath))
	defer f.Close()
	io.Copy(f, file)
	return nil
}
