package script

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) WriteContentToFile(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	filepathStr := filepath.Join(GlobalOption.ProjectPath, L.ToString(1))
	content := L.ToString(2)
	err := ioutil.WriteFile(filepathStr, []byte(content), os.FileMode(0777))
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	return 0
}

func (p *Script) ReadContentFromFile(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	filepathStr := filepath.Join(GlobalOption.ProjectPath, L.ToString(1))
	content, err := ioutil.ReadFile(filepathStr)
	if nil != err && false == strings.Contains(err.Error(), "no such file or directory") {
		log.Println(err.Error(), filepathStr)
		return 0
	}

	L.Push(lua.LString(content))
	return 1
}
