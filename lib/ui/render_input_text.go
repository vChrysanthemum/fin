package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyInputText(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeInputText).Editor

	p.normalRenderNodeBlock(node)

	uiBuffer.TextFgColor = termui.ColorBlue

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Block.Y + uiBuffer.Block.Height

	return
}
