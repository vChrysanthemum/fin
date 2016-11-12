package ui

import (
	"log"

	"golang.org/x/net/html"

	lua "github.com/yuin/gopher-lua"
)

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

func (p *Script) luaFuncNodeSetAttribute(L *lua.LState) int {
	if L.GetTop() < 3 {
		return 0
	}

	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		L.Push(lua.LNil)
		return 0
	}
	isUIChange, isNeedRerenderPage := node.ParseAttribute([]html.Attribute{
		html.Attribute{Key: L.ToString(2), Val: L.ToString(3)},
	})
	if false == isUIChange {
		return 0
	}
	if true == isNeedRerenderPage {
		p.page.Rerender()
	} else {
		node.uiRender()
	}
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
	node.uiRender()
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
		isNeedRerenderPage := node.SetText(text)
		if true == isNeedRerenderPage {
			p.page.Rerender()
		} else {
			node.uiRender()
		}
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

func (p *Script) luaFuncNodeRegisterKeyPressEnterHandler(L *lua.LState) int {
	if L.GetTop() < 2 {
		L.Push(lua.LNil)
		return 1
	}

	lu := L.ToUserData(1)
	callback := L.ToFunction(2)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		L.Push(lua.LNil)
		return 1
	}

	key := node.RegisterKeyPressEnterHandler(func(_node *Node, args ...interface{}) {
		_L := args[0].(*lua.LState)
		_callback := args[1].(*lua.LFunction)
		luaNode := _L.NewUserData()
		luaNode.Value = node
		if err := _L.CallByParam(lua.P{
			Fn:      _callback,
			NRet:    0,
			Protect: true,
		}, luaNode); err != nil {
			panic(err)
		}
	}, L, callback)

	L.Push(lua.LString(key))
	return 1
}

func (p *Script) luaFuncNodeRemoveKeyPressEnterHandler(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	lu := L.ToUserData(1)
	key := L.ToString(2)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		return 0
	}

	log.Println(key, node.HtmlData)
	node.RemoveKeyPressEnterHandler(key)
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
