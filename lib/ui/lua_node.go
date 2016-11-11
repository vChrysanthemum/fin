package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodePointerFromUserData(L *lua.LState, lu *lua.LUserData) *Node {
	if nil == lu || nil == lu.Value {
		return nil
	}

	node, ok := lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	return node
}

func (p *Script) luaFuncGetNodePointer(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	nodeId := L.ToString(1)

	var (
		node *Node
		ok   bool
	)

	node, ok = p.page.IdToNodeMap[nodeId]

	if true == ok {
		luaNode := L.NewUserData()
		luaNode.Value = node
		L.Push(luaNode)
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (p *Script) luaFuncNodeUIRender(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		return 0
	}

	node.uiRender()

	return 0
}

func (p *Script) luaFuncNodeSetActive(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		L.Push(lua.LNil)
		return 0
	}
	p.page.SetActiveNode(node)
	return 0
}

func (p *Script) luaFuncNodeGetHtmlData(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(node.HtmlData))

	return 1
}

func (p *Script) luaFuncNodeSetText(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	text := L.ToString(2)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		return 0
	}

	if nil != node.SetText {
		node.SetText(text)
	}

	return 0
}

func (p *Script) luaFuncNodeGetValue(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node || nil == node.GetValue {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(node.GetValue()))
	return 1
}

func (p *Script) luaFuncNodeOnKeyPressEnter(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	callback := L.ToFunction(2)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node || nil == node.OnKeyPressEnter {
		return 0
	}

	go func(_L *lua.LState, _node *Node, _callback *lua.LFunction) {
		_node.OnKeyPressEnter()
		luaNode := _L.NewUserData()
		luaNode.Value = node
		if err := _L.CallByParam(lua.P{
			Fn:      _callback,
			NRet:    0,
			Protect: true,
		}, luaNode); err != nil {
			panic(err)
		}
	}(L, node, callback)

	return 0
}

func (p *Script) luaFuncNodeRemove(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		return 0
	}

	p.page.RemoveNode(node)

	return 0
}
