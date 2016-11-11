package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeGaugePointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeGauge {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node      *Node
		nodeGauge *NodeGauge
		ok        bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeGauge, ok = node.Data.(*NodeGauge); false == ok {
		return nil
	}

	return nodeGauge
}

func (p *Script) luaFuncNodeGaugeSetPercent(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeGauge := p._getNodeGaugePointerFromUserData(L, lu)
	if nil == nodeGauge {
		return 0
	}

	nodeGauge.SetPercent(L.ToInt(2))
	return 0
}
