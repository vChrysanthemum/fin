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
