package ui

import (
	"in/script"
	"log"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

type Script struct {
	script.Script
	page     *Page
	luaDocs  []string
	luaState *lua.LState
}

func (p *Page) prepareScript() {
	var err error
	s := new(Script)

	s.CancelSigs = make(map[string]chan bool, 0)

	s.page = p

	s.luaDocs = make([]string, 0)

	s.luaState = lua.NewState()

	luaBase := s.luaState.NewTable()
	s.luaState.SetGlobal("base", luaBase)

	s.luaState.SetField(luaBase, "ResBaseDir", lua.LString(
		filepath.Join(GlobalOption.ResBaseDir, "lua/"),
	))

	s.Script.RegisterInLuaTable(s.luaState, luaBase)

	s.luaState.SetField(luaBase, "WindowWidth", s.luaState.NewFunction(s.luaFuncWindowWidth))
	s.luaState.SetField(luaBase, "WindowHeight", s.luaState.NewFunction(s.luaFuncWindowHeight))
	s.luaState.SetField(luaBase, "WindowConfirm", s.luaState.NewFunction(s.luaFuncWindowConfirm))

	s.luaState.SetField(luaBase, "GetNodePointer", s.luaState.NewFunction(s.luaFuncGetNodePointer))

	s.luaState.SetField(luaBase, "NodeSetAttribute", s.luaState.NewFunction(s.luaFuncNodeSetAttribute))
	s.luaState.SetField(luaBase, "NodeSetActive", s.luaState.NewFunction(s.luaFuncNodeSetActive))

	s.luaState.SetField(luaBase, "NodeGetHtmlData", s.luaState.NewFunction(s.luaFuncNodeGetHtmlData))
	s.luaState.SetField(luaBase, "NodeSetText", s.luaState.NewFunction(s.luaFuncNodeSetText))
	s.luaState.SetField(luaBase, "NodeGetValue", s.luaState.NewFunction(s.luaFuncNodeGetValue))

	s.luaState.SetField(luaBase, "NodeSetCursor", s.luaState.NewFunction(s.luaFuncNodeSetCursor))
	s.luaState.SetField(luaBase, "NodeResumeCursor", s.luaState.NewFunction(s.luaFuncNodeResumeCursor))
	s.luaState.SetField(luaBase, "NodeHideCursor", s.luaState.NewFunction(s.luaFuncNodeHideCursor))

	s.luaState.SetField(luaBase, "NodeRegisterKeyPressHandler",
		s.luaState.NewFunction(s.luaFuncNodeRegisterKeyPressHandler))
	s.luaState.SetField(luaBase, "NodeRegisterKeyPressEnterHandler",
		s.luaState.NewFunction(s.luaFuncNodeRegisterKeyPressEnterHandler))
	s.luaState.SetField(luaBase, "NodeRemoveKeyPressEnterHandler",
		s.luaState.NewFunction(s.luaFuncNodeRemoveKeyPressEnterHandler))

	s.luaState.SetField(luaBase, "NodeRemove", s.luaState.NewFunction(s.luaFuncNodeRemove))

	s.luaState.SetField(luaBase, "NodeCanvasSet", s.luaState.NewFunction(s.luaFuncNodeCanvasSet))
	s.luaState.SetField(luaBase, "NodeCanvasDraw", s.luaState.NewFunction(s.luaFuncNodeCanvasDraw))

	s.luaState.SetField(luaBase, "NodeSelectAppendOption",
		s.luaState.NewFunction(s.luaFuncNodeSelectAppendOption))
	s.luaState.SetField(luaBase, "NodeSelectClearOptions",
		s.luaState.NewFunction(s.luaFuncNodeSelectClearOptions))

	s.luaState.SetField(luaBase, "NodeTerminalRegisterCommandHandle",
		s.luaState.NewFunction(s.luaFuncNodeTerminalRegisterCommandHandle))
	s.luaState.SetField(luaBase, "NodeTerminalRemoveCommandHandle",
		s.luaState.NewFunction(s.luaFuncNodeTerminalRemoveCommandHandle))
	s.luaState.SetField(luaBase, "NodeTerminalWriteNewLine",
		s.luaState.NewFunction(s.luaFuncNodeTerminalWriteNewLine))
	s.luaState.SetField(luaBase, "NodeTerminalClearLines",
		s.luaState.NewFunction(s.luaFuncNodeTerminalClearLines))

	err = s.luaState.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/script/core.lua"))
	if nil != err {
		log.Println(err)
		panic(err)
	}
	err = s.luaState.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/ui/core.lua"))
	if nil != err {
		log.Println(err)
		panic(err)
	}

	p.Script = s
}

func (p *Script) appendDoc(doc, docType string) {
	if "text/lua" != docType {
		return
	}
	p.luaDocs = append(p.luaDocs, doc)
}

func (p *Page) GetLuaDocs(index int) string {
	return p.Script.luaDocs[index]
}

func (p *Page) AppendScript(doc, docType string) {
	p.Script.appendDoc(doc, docType)
}

func (p *Script) Run() {
	for _, doc := range p.luaDocs {
		if err := p.luaState.DoString(doc); nil != err {
			log.Println(err)
			panic(err)
		}
	}
}
