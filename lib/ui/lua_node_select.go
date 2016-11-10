package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeSelectPointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeSelect {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node       *Node
		nodeSelect *NodeSelect
		ok         bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeSelect, ok = node.Data.(*NodeSelect); false == ok {
		return nil
	}

	return nodeSelect
}

func (p *Script) luaFuncNodeSelectAppendOption(L *lua.LState) int {
	if L.GetTop() < 3 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeSelect := p._getNodeSelectPointerFromUserData(L, lu)
	if nil == nodeSelect {
		return 0
	}

	nodeSelect.AppendOption(L.ToString(2), L.ToString(3))

	p.page.Rerender()
	return 0
}

func (p *Script) luaFuncNodeSelectClearOptions(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeSelect := p._getNodeSelectPointerFromUserData(L, lu)
	if nil == nodeSelect {
		return 0
	}

	nodeSelect.ClearOptions()

	p.page.Rerender()
	return 0
}
