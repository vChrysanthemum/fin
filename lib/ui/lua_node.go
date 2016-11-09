package ui

import (
	"unicode/utf8"

	"github.com/gizak/termui"
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

func (p *Script) _getNodeCanvasPointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeCanvas {

	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node       *Node
		nodeCanvas *NodeCanvas
		ok         bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeCanvas, ok = node.Data.(*NodeCanvas); false == ok {
		return nil
	}

	return nodeCanvas
}

func (p *Script) luaFuncGetNodePointer(L *lua.LState) int {
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

func (p *Script) luaFuncNodeSetActive(L *lua.LState) int {
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
	lu := L.ToUserData(1)
	node := p._getNodePointerFromUserData(L, lu)
	if nil == node {
		return 0
	}

	p.page.RemoveNode(node)

	return 0
}

func (p *Script) luaFuncNodeCanvasSet(L *lua.LState) int {
	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	if nil == nodeCanvas {
		return 0
	}

	params := L.GetTop()
	if params < 4 {
		return 0
	}

	ch, _ := utf8.DecodeRuneInString(L.ToString(4))
	colorFg := termui.ColorDefault
	colorBg := termui.ColorBlue
	if params >= 5 {
		colorFg = ColorToTermuiAttribute(L.ToString(5), termui.ColorBlue)
	}
	if params >= 6 {
		colorBg = ColorToTermuiAttribute(L.ToString(6), termui.ColorDefault)
	}
	nodeCanvas.Canvas.Set(L.ToInt(2), L.ToInt(3), &termui.Cell{ch, colorFg, colorBg})
	return 0
}

func (p *Script) luaFuncNodeCanvasDraw(L *lua.LState) int {
	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	if nil == nodeCanvas {
		return 0
	}
	uirender(nodeCanvas.Canvas)
	return 0
}
