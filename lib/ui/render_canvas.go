package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyCanvas(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeCanvas).Canvas
	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border
	uiBuffer.BorderFg = node.BorderFg
	uiBuffer.Width = node.Width
	uiBuffer.Height = node.Height
	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY
	if "" != node.ColorFg {
		uiBuffer.ItemFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.ItemBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	p.pushWorkingNode(node)

	return
}
