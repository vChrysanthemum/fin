package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyInputText(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeInputText).Editor
	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border
	uiBuffer.BorderFg = node.BorderFg
	uiBuffer.Width = node.Width
	uiBuffer.Height = node.Height
	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY
	uiBuffer.TextFgColor = termui.ColorBlue

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	p.pushWorkingNode(node)

	return
}
