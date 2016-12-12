package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeModalPointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeModal {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node      *Node
		nodeModal *NodeModal
		ok        bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeModal, ok = node.Data.(*NodeModal); false == ok {
		return nil
	}

	return nodeModal
}

func (p *Script) luaFuncNodeModalDoString(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeModal := p._getNodeModalPointerFromUserData(L, lu)
	if nil == nodeModal {
		return 0
	}

	callback := L.ToString(2)

	script := nodeModal.page.Script
	err := script.luaState.DoString(callback)

	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	} else {
		return 0
	}
}

func (p *Script) luaFuncNodeModalShow(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeModal := p._getNodeModalPointerFromUserData(L, lu)
	if nil == nodeModal {
		return 0
	}

	p.page.ClearActiveNode()
	uiClear()
	nodeModal.page.Refresh()
	go nodeModal.page.Script.Run()

	p.page.CurrentModal = nodeModal

	return 0
}

func (p *Script) luaFuncModalClose(L *lua.LState) int {
	if nil == p.page.MainPage {
		return 0
	}

	p.page.ClearActiveNode()
	uiClear()
	p.page.MainPage.Refresh()

	return 0
}

func (p *Script) luaFuncMainPageDoString(L *lua.LState) int {
	if nil == p.page.MainPage {
		return 0
	}

	if L.GetTop() < 1 {
		return 0
	}

	callback := L.ToString(1)

	script := p.page.MainPage.Script
	err := script.luaState.DoString(callback)

	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	} else {
		return 0
	}
}
