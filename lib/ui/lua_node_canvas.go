package ui

import (
	"fin/ui/utils"
	"unicode/utf8"

	"github.com/gizak/termui"
	lua "github.com/yuin/gopher-lua"
)

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

func (p *Script) luaFuncNodeCanvasClean(L *lua.LState) int {
	params := L.GetTop()
	if params < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	nodeCanvas.Canvas.Clean()
	return 0
}

func (p *Script) luaFuncNodeCanvasUnSet(L *lua.LState) int {
	params := L.GetTop()
	if params < 3 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	nodeCanvas.Canvas.UnSet(L.ToInt(2), L.ToInt(3))
	return 0
}

func (p *Script) luaFuncNodeCanvasSet(L *lua.LState) int {
	params := L.GetTop()
	if params < 4 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	if nil == nodeCanvas {
		return 0
	}

	ch, _ := utf8.DecodeRuneInString(L.ToString(4))
	colorFg := utils.ColorDefault
	colorBg := utils.ColorBlue
	if params >= 5 {
		colorFg = utils.ColorToTermuiAttribute(L.ToString(5), utils.ColorBlue)
	}
	if params >= 6 {
		colorBg = utils.ColorToTermuiAttribute(L.ToString(6), utils.ColorDefault)
	}
	nodeCanvas.Canvas.Set(L.ToInt(2), L.ToInt(3), &termui.Cell{ch, colorFg, colorBg, 0, 0, 0})
	return 0
}

func (p *Script) luaFuncNodeCanvasDraw(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	lu := L.ToUserData(1)
	nodeCanvas := p._getNodeCanvasPointerFromUserData(L, lu)
	if nil == nodeCanvas {
		return 0
	}
	nodeCanvas.Node.UIRender()
	return 0
}
