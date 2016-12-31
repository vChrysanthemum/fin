package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeTabpanePointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeTabpane {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node        *Node
		nodeTabpane *NodeTabpane
		ok          bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeTabpane, ok = node.Data.(*NodeTabpane); false == ok {
		return nil
	}

	return nodeTabpane
}

func (p *Script) luaFuncNodeTabpaneSetActiveTab(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	nodeTabpane := p._getNodeTabpanePointerFromUserData(L, L.ToUserData(1))
	name := L.ToString(2)
	if nil == nodeTabpane {
		return 0
	}

	nodeTabpane.SetActiveTab(name)
	return 0
}
