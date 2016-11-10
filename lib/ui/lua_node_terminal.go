package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeTerminalPointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeTerminal {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node         *Node
		nodeTerminal *NodeTerminal
		ok           bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeTerminal, ok = node.Data.(*NodeTerminal); false == ok {
		return nil
	}

	return nodeTerminal
}

func (p *Script) luaFuncNodeTerminalPopNewCommand(L *lua.LState) int {
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	lu := L.ToUserData(1)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(string(nodeTerminal.PopNewCommand())))
	return 1
}

func (p *Script) luaFuncNodeTerminalWriteLine(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		return 0
	}

	nodeTerminal.WriteLine(L.ToString(2))
	return 1
}
