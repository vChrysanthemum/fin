package ui

import (
	. "in/ui/utils"

	"github.com/gizak/termui"
)

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.uiBuffer.(*termui.Par)

	p.normalRenderNodeBlock(node)

	if true == node.isShouldCalculateHeight {
		if true == node.uiBlock.Border {
			node.uiBlock.Height = CalculateTextHeight(uiBuffer.Text, node.uiBlock.Width) + 2
		} else {
			node.uiBlock.Height = CalculateTextHeight(uiBuffer.Text, node.uiBlock.Width)
		}
	}

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
