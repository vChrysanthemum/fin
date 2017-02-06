package ui

import (
	"github.com/gizak/termui"
	lua "github.com/yuin/gopher-lua"
)

func (p *Script) luaFuncUIReRender(L *lua.LState) int {
	p.page.SetActiveNode(nil)
	p.page.ReRender()
	return 0
}

func (p *Script) luaFuncWindowWidth(L *lua.LState) int {
	L.Push(lua.LNumber(termui.TermWidth()))
	return 1
}

func (p *Script) luaFuncWindowHeight(L *lua.LState) int {
	L.Push(lua.LNumber(termui.TermHeight()))
	return 1
}

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	content := L.ToString(1)
	callback := L.ToFunction(2)
	page, err := Parse(content)
	if nil != err {
		return 0
	}

	err = page.Render()
	if nil != err {
		return 0
	}

	nodeSelect, _ := page.IDToNodeMap["SelectConfirm"]
	if nil == nodeSelect {
		return 0
	}

	nodeDataSelect := nodeSelect.Data.(*NodeSelect)

	p.page.ClearActiveNode()
	page.uiRender()
	p.page.SetActiveNode(nodeSelect)

	nodeDataSelect.DisableQuit = true

	key := nodeSelect.RegisterKeyPressEnterHandler(func(_node *Node, args ...interface{}) {
		_L := args[0].(*lua.LState)
		_callback := args[1].(*lua.LFunction)
		_page := args[2].(*Page)
		_mainPage := args[3].(*Page)
		_key := args[4].(string)
		luaNode := _L.NewUserData()
		luaNode.Value = _node

		_ret, _ok := _node.Data.(*NodeSelect).NodeDataGetValue()
		var _paramData lua.LValue
		if false == _ok {
			_paramData = lua.LNil
		} else {
			_paramData = lua.LString(_ret)
		}

		_page.Clear()
		_mainPage.SetActiveNode(nil)
		_mainPage.ReRender()
		_node.RemoveKeyPressEnterHandler(_key)

		if err := p.Script.LuaCallByParam(_L, lua.P{
			Fn:      _callback,
			NRet:    0,
			Protect: true,
		}, _paramData); err != nil {
			panic(err)
		}

	}, L, callback, page, p.page)
	_job := nodeSelect.KeyPressEnterHandlers[key]
	_job.Args = append(_job.Args, key)
	nodeSelect.KeyPressEnterHandlers[key] = _job

	L.Push(lua.LString(key))

	return 1
}
