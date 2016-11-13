package ui

import (
	"log"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

type Script struct {
	page     *Page
	luaDocs  []string
	luaState *lua.LState

	CancelSigs map[string](chan bool)
}

func (p *Page) prepareScript() {
	var err error
	script := new(Script)

	script.CancelSigs = make(map[string]chan bool, 0)

	script.page = p

	script.luaDocs = make([]string, 0)

	script.luaState = lua.NewState()

	luaBase := script.luaState.NewTable()
	script.luaState.SetGlobal("base", luaBase)

	script.luaState.SetField(luaBase, "ResBaseDir", lua.LString(
		filepath.Join(GlobalOption.ResBaseDir, "lua/"),
	))

	script.luaState.SetField(luaBase, "Log", script.luaState.NewFunction(script.LuaFuncLog))

	script.luaState.SetField(luaBase, "SetInterval", script.luaState.NewFunction(script.LuaFuncSetInterval))
	script.luaState.SetField(luaBase, "SetTimeout", script.luaState.NewFunction(script.LuaFuncSetTimeout))
	script.luaState.SetField(luaBase, "SendCancelSig", script.luaState.NewFunction(script.LuaFuncSendCancelSig))

	script.luaState.SetField(luaBase, "WindowConfirm", script.luaState.NewFunction(script.luaFuncWindowConfirm))

	script.luaState.SetField(luaBase, "GetNodePointer", script.luaState.NewFunction(script.luaFuncGetNodePointer))

	script.luaState.SetField(luaBase, "NodeSetAttribute", script.luaState.NewFunction(script.luaFuncNodeSetAttribute))
	script.luaState.SetField(luaBase, "NodeSetActive", script.luaState.NewFunction(script.luaFuncNodeSetActive))

	script.luaState.SetField(luaBase, "NodeGetHtmlData", script.luaState.NewFunction(script.luaFuncNodeGetHtmlData))
	script.luaState.SetField(luaBase, "NodeSetText", script.luaState.NewFunction(script.luaFuncNodeSetText))
	script.luaState.SetField(luaBase, "NodeGetValue", script.luaState.NewFunction(script.luaFuncNodeGetValue))

	script.luaState.SetField(luaBase, "NodeSetCursor", script.luaState.NewFunction(script.luaFuncNodeSetCursor))
	script.luaState.SetField(luaBase, "NodeResumeCursor", script.luaState.NewFunction(script.luaFuncNodeResumeCursor))
	script.luaState.SetField(luaBase, "NodeHideCursor", script.luaState.NewFunction(script.luaFuncNodeHideCursor))

	script.luaState.SetField(luaBase, "NodeRegisterKeyPressHandler",
		script.luaState.NewFunction(script.luaFuncNodeRegisterKeyPressHandler))
	script.luaState.SetField(luaBase, "NodeRegisterKeyPressEnterHandler",
		script.luaState.NewFunction(script.luaFuncNodeRegisterKeyPressEnterHandler))
	script.luaState.SetField(luaBase, "NodeRemoveKeyPressEnterHandler",
		script.luaState.NewFunction(script.luaFuncNodeRemoveKeyPressEnterHandler))

	script.luaState.SetField(luaBase, "NodeRemove", script.luaState.NewFunction(script.luaFuncNodeRemove))

	script.luaState.SetField(luaBase, "NodeCanvasSet", script.luaState.NewFunction(script.luaFuncNodeCanvasSet))
	script.luaState.SetField(luaBase, "NodeCanvasDraw", script.luaState.NewFunction(script.luaFuncNodeCanvasDraw))

	script.luaState.SetField(luaBase, "NodeSelectAppendOption",
		script.luaState.NewFunction(script.luaFuncNodeSelectAppendOption))
	script.luaState.SetField(luaBase, "NodeSelectClearOptions",
		script.luaState.NewFunction(script.luaFuncNodeSelectClearOptions))

	script.luaState.SetField(luaBase, "NodeTerminalRegisterCommandHandle",
		script.luaState.NewFunction(script.luaFuncNodeTerminalRegisterCommandHandle))
	script.luaState.SetField(luaBase, "NodeTerminalRemoveCommandHandle",
		script.luaState.NewFunction(script.luaFuncNodeTerminalRemoveCommandHandle))
	script.luaState.SetField(luaBase, "NodeTerminalWriteNewLine",
		script.luaState.NewFunction(script.luaFuncNodeTerminalWriteNewLine))
	script.luaState.SetField(luaBase, "NodeTerminalClearLines",
		script.luaState.NewFunction(script.luaFuncNodeTerminalClearLines))

	err = script.luaState.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/ui/core.lua"))
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
			log.Println(err)
			panic(err)
		}
	}
}
