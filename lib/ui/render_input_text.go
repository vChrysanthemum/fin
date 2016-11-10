package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyInputText(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeInputText).Editor

	node.uiBlock = &uiBuffer.Block
	p.normalRenderNodeBlock(node)

	uiBuffer.TextFgColor = termui.ColorBlue

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
