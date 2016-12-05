package ui

import (
	"in/script"
	"log"
	"path/filepath"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

type ScriptDoc struct {
	DataType string
	Data     string
}

type Script struct {
	Script               *script.Script
	page                 *Page
	luaDocs              []ScriptDoc
	luaState             *lua.LState
	LuaCallByParamLocker *sync.RWMutex
}

func (p *Page) prepareScript() {
	var err error
	s := new(Script)
	s.LuaCallByParamLocker = new(sync.RWMutex)
	s.Script = script.NewScript(s.LuaCallByParamLocker)

	s.page = p

	s.luaDocs = make([]ScriptDoc, 0)

	s.luaState = lua.NewState()

	luaBase := s.luaState.NewTable()

	s.Script.RegisterScript(s.luaState)

	s.luaState.SetGlobal("base", luaBase)

	s.Script.RegisterBaseTable(s.luaState, luaBase)

	s.luaState.SetField(luaBase, "UIRerender", s.luaState.NewFunction(s.luaFuncUIRerender))

	s.luaState.SetField(luaBase, "WindowWidth", s.luaState.NewFunction(s.luaFuncWindowWidth))
	s.luaState.SetField(luaBase, "WindowHeight", s.luaState.NewFunction(s.luaFuncWindowHeight))
	s.luaState.SetField(luaBase, "WindowConfirm", s.luaState.NewFunction(s.luaFuncWindowConfirm))

	s.luaState.SetField(luaBase, "GetNodePointer", s.luaState.NewFunction(s.luaFuncGetNodePointer))

	s.luaState.SetField(luaBase, "NodeWidth", s.luaState.NewFunction(s.luaFuncNodeWidth))
	s.luaState.SetField(luaBase, "NodeHeight", s.luaState.NewFunction(s.luaFuncNodeHeight))
	s.luaState.SetField(luaBase, "NodeGetAttribute", s.luaState.NewFunction(s.luaFuncNodeGetAttribute))
	s.luaState.SetField(luaBase, "NodeSetAttribute", s.luaState.NewFunction(s.luaFuncNodeSetAttribute))
	s.luaState.SetField(luaBase, "NodeSetActive", s.luaState.NewFunction(s.luaFuncNodeSetActive))

	s.luaState.SetField(luaBase, "NodeGetHtmlData", s.luaState.NewFunction(s.luaFuncNodeGetHtmlData))
	s.luaState.SetField(luaBase, "NodeSetText", s.luaState.NewFunction(s.luaFuncNodeSetText))
	s.luaState.SetField(luaBase, "NodeGetValue", s.luaState.NewFunction(s.luaFuncNodeGetValue))

	s.luaState.SetField(luaBase, "NodeSetCursor", s.luaState.NewFunction(s.luaFuncNodeSetCursor))
	s.luaState.SetField(luaBase, "NodeResumeCursor", s.luaState.NewFunction(s.luaFuncNodeResumeCursor))
	s.luaState.SetField(luaBase, "NodeHideCursor", s.luaState.NewFunction(s.luaFuncNodeHideCursor))

	s.luaState.SetField(luaBase, "NodeRegisterLuaActiveModeHandler",
		s.luaState.NewFunction(s.luaFuncNodeRegisterLuaActiveModeHandler))
	s.luaState.SetField(luaBase, "NodeRemoveLuaActiveModeHandler",
		s.luaState.NewFunction(s.luaFuncNodeRemoveLuaActiveModeHandler))
	s.luaState.SetField(luaBase, "NodeRegisterKeyPressHandler",
		s.luaState.NewFunction(s.luaFuncNodeRegisterKeyPressHandler))
	s.luaState.SetField(luaBase, "NodeRegisterKeyPressEnterHandler",
		s.luaState.NewFunction(s.luaFuncNodeRegisterKeyPressEnterHandler))
	s.luaState.SetField(luaBase, "NodeRemoveKeyPressEnterHandler",
		s.luaState.NewFunction(s.luaFuncNodeRemoveKeyPressEnterHandler))

	s.luaState.SetField(luaBase, "NodeRemove", s.luaState.NewFunction(s.luaFuncNodeRemove))

	s.luaState.SetField(luaBase, "NodeCanvasClean", s.luaState.NewFunction(s.luaFuncNodeCanvasClean))
	s.luaState.SetField(luaBase, "NodeCanvasUnSet", s.luaState.NewFunction(s.luaFuncNodeCanvasUnSet))
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
		panic(err)
	}
	err = s.luaState.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/ui/core.lua"))
	if nil != err {
		panic(err)
	}

	p.Script = s
}

func (p *Script) appendDoc(doc ScriptDoc) {
	p.luaDocs = append(p.luaDocs, doc)
}

func (p *Page) GetLuaDocs(index int) ScriptDoc {
	return p.Script.luaDocs[index]
}

func (p *Page) AppendScript(doc ScriptDoc) {
	p.Script.appendDoc(doc)
}

func (p *Script) Run() {
	var err error
	for _, doc := range p.luaDocs {
		switch doc.DataType {
		case "file":
			err = p.luaState.DoFile(
				filepath.Join(GlobalOption.ResBaseDir, "project", GlobalOption.ProjectName, doc.Data))
		case "string":
			err = p.luaState.DoString(doc.Data)
		}
		if nil != err {
			panic(err)
		}
	}
}

func luaCallByParam(L *lua.LState, cp lua.P, args ...lua.LValue) error {
	defer func() {
		if rcv := recover(); nil != rcv {
			log.Println(rcv)
		}
	}()
	return L.CallByParam(cp, args...)
}
