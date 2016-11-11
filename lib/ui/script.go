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
	script.luaState.SetField(luaBase, "Log", script.luaState.NewFunction(script.LuaFuncLog))
	script.luaState.SetField(luaBase, "WindowConfirm", script.luaState.NewFunction(script.luaFuncWindowConfirm))

	script.luaState.SetField(luaBase, "GetNodePointer", script.luaState.NewFunction(script.luaFuncGetNodePointer))
	script.luaState.SetField(luaBase, "NodeRender", script.luaState.NewFunction(script.luaFuncNodeRender))
	script.luaState.SetField(luaBase, "NodeSetActive", script.luaState.NewFunction(script.luaFuncNodeSetActive))
	script.luaState.SetField(luaBase, "NodeGetHtmlData", script.luaState.NewFunction(script.luaFuncNodeGetHtmlData))
	script.luaState.SetField(luaBase, "NodeSetText", script.luaState.NewFunction(script.luaFuncNodeSetText))
	script.luaState.SetField(luaBase, "NodeGetValue", script.luaState.NewFunction(script.luaFuncNodeGetValue))
	script.luaState.SetField(luaBase, "NodeOnKeyPressEnter",
		script.luaState.NewFunction(script.luaFuncNodeOnKeyPressEnter))
	script.luaState.SetField(luaBase, "NodeRemove", script.luaState.NewFunction(script.luaFuncNodeRemove))

	script.luaState.SetField(luaBase, "NodeCanvasSet", script.luaState.NewFunction(script.luaFuncNodeCanvasSet))
	script.luaState.SetField(luaBase, "NodeCanvasDraw", script.luaState.NewFunction(script.luaFuncNodeCanvasDraw))

	script.luaState.SetField(luaBase, "NodeSelectAppendOption",
		script.luaState.NewFunction(script.luaFuncNodeSelectAppendOption))
	script.luaState.SetField(luaBase, "NodeSelectClearOptions",
		script.luaState.NewFunction(script.luaFuncNodeSelectClearOptions))

	script.luaState.SetField(luaBase, "NodeTerminalPopNewCommand",
		script.luaState.NewFunction(script.luaFuncNodeTerminalPopNewCommand))
	script.luaState.SetField(luaBase, "NodeTerminalWriteNewLine",
		script.luaState.NewFunction(script.luaFuncNodeTerminalWriteNewLine))

	err = script.luaState.DoFile(filepath.Join(GlobalOption.LuaResBaseDir, "ui/core.lua"))
	if nil != err {
		panic(err)
	}

	p.script = script
}

func (p *Script) appendDoc(doc, docType string) {
	if "text/lua" != docType {
		return
	}
	p.luaDocs = append(p.luaDocs, doc)
}

func (p *Page) GetLuaDocs(index int) string {
	return p.script.luaDocs[index]
}

func (p *Page) AppendScript(doc, docType string) {
	p.script.appendDoc(doc, docType)
}

func (p *Script) Run() {
	for _, doc := range p.luaDocs {
		if err := p.luaState.DoString(doc); nil != err {
			panic(err)
		}
	}
}
