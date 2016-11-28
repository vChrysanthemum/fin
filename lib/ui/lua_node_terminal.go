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
		L.Push(lua.LNil)
		return 1
	}

	lu := L.ToUserData(1)
	callback := L.ToFunction(2)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		L.Push(lua.LNil)
		return 1
	}

	key := nodeTerminal.Node.RegisterKeyPressEnterHandler(func(_node *Node, args ...interface{}) {
		_L := args[0].(*lua.LState)
		_callback := args[1].(*lua.LFunction)
		_nodeTerminal := _node.Data.(*NodeTerminal)
		luaNode := _L.NewUserData()
		luaNode.Value = _node
		if err := luaCallByParam(_L, lua.P{
			Fn:      _callback,
			NRet:    0,
			Protect: true,
		}, luaNode, lua.LString(_nodeTerminal.PopNewCommand())); err != nil {
			panic(err)
		}
	}, L, callback)

	L.Push(lua.LString(key))
	return 1
}

func (p *Script) luaFuncNodeTerminalRemoveCommandHandle(L *lua.LState) int {
	if L.GetTop() < 2 {
		L.Push(lua.LNil)
		return 1
	}

	lu := L.ToUserData(1)
	key := L.ToString(2)
	nodeTerminal := p._getNodeTerminalPointerFromUserData(L, lu)
	if nil == nodeTerminal {
		L.Push(lua.LNil)
		return 1
	}

	nodeTerminal.Node.RemoveKeyPressEnterHandler(key)
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
