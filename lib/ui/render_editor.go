package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyEditor(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeEditor).Editor
	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border
	uiBuffer.BorderFg = node.BorderFg
	uiBuffer.Border = node.Border
	uiBuffer.BorderFg = node.BorderFg
	uiBuffer.Width = node.Width
	uiBuffer.Height = node.Height
	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY
	uiBuffer.TextFgColor = termui.ColorBlue

	node.uiBuffer = uiBuffer

	p.bufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	return
}
