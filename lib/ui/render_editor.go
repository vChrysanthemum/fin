package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyEditor(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeEditor).Editor

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Block.Y + uiBuffer.Block.Height

	return
}
