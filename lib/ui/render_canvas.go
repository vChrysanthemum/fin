package ui

import (
	. "in/ui/utils"

	"github.com/gizak/termui"
)

func (p *Page) renderBodyCanvas(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.Data.(*NodeCanvas).Canvas

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.ItemFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.ItemBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
