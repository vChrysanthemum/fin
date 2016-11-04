package ui

import (
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

type Script struct {
	page     *Page
	luaDocs  []string
	luaState *lua.LState
}

func (p *Page) prepareScript() {
	var err error
	script := new(Script)

	script.page = p

	script.luaDocs = make([]string, 0)

	script.luaState = lua.NewState()

	luaBase := script.luaState.NewTable()
	script.luaState.SetGlobal("base", luaBase)
	script.luaState.SetField(luaBase, "ResBaseDir", lua.LString(GlobalOption.LuaResBaseDir))
	script.luaState.SetField(luaBase, "GetNodePointer", script.luaState.NewFunction(script.luaFuncGetNodePointer))
	script.luaState.SetField(luaBase, "GetNodeHtmlData", script.luaState.NewFunction(script.luaFuncGetNodeHtmlData))

	err = script.luaState.DoFile(filepath.Join(GlobalOption.LuaResBaseDir, "ui/main.lua"))
	if nil != err {
		panic(err)
	}

	p.script = script
}

func (p *Script) appendDoc(doc, docType string) {
	if "text/lua" != docType {
		return
	}

	if err := p.luaState.DoString(doc); nil != err {
		panic(err)
	}
	p.luaDocs = append(p.luaDocs, doc)
}

func (p *Page) GetLuaDocs(index int) string {
	return p.script.luaDocs[index]
}

func (p *Page) AppendScript(doc, docType string) {
	p.script.appendDoc(doc, docType)
}
