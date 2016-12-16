package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyInputText(node *Node) {
	uiBuffer := node.Data.(*NodeInputText).Editor

	p.normalRenderNodeBlock(node)

	uiBuffer.TextFgColor = termui.ColorBlue

	p.BufferersAppend(node, uiBuffer)

	return
}
