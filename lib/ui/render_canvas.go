package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyCanvas(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeCanvas).Canvas

	node.uiBlock = &uiBuffer.Block
	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.ItemFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.ItemBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
