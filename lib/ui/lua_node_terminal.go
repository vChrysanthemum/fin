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

func (p *Script) luaFuncNodeTerminalRegisterCommandHandle(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	callback := L.ToFunction(2)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		return 0
	}

	go func(_L *lua.LState, _node *Node, _callback *lua.LFunction) {
		_node.OnKeyPressEnter()
		_nodeTerminal := _node.Data.(*NodeTerminal)
		luaNode := _L.NewUserData()
		luaNode.Value = _node
		if err := _L.CallByParam(lua.P{
			Fn:      _callback,
			NRet:    0,
			Protect: true,
		}, luaNode, lua.LString(_nodeTerminal.PopNewCommand())); err != nil {
			panic(err)
		}
		_nodeTerminal.PrepareNewCommand()
		_node.uiRender()
	}(L, nodeTerminal.Node, callback)

	return 0
}

func (p *Script) luaFuncNodeTerminalWriteNewLine(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		return 0
	}

	nodeTerminal.WriteNewLine(L.ToString(2))
	nodeTerminal.Node.uiRender()
	return 1
}

func (p *Script) luaFuncNodeTerminalClearLines(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		return 0
	}

	nodeTerminal.ClearLines()
	nodeTerminal.Node.uiRender()
	return 0
}
