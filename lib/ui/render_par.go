package ui

import (
	uiutils "in/ui/utils"

	"github.com/gizak/termui"
)

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.uiBuffer.(*termui.Par)

	p.normalRenderNodeBlock(node)

	if true == node.isShouldCalculateHeight {
		if true == node.UIBlock.Border {
			node.UIBlock.Height = uiutils.CalculateTextHeight(uiBuffer.Text, node.UIBlock.Width) + 2
		} else {
			node.UIBlock.Height = uiutils.CalculateTextHeight(uiBuffer.Text, node.UIBlock.Width)
		}
		node.UIBlock.Height += node.UIBlock.PaddingTop
		node.UIBlock.Height += node.UIBlock.PaddingBottom
	}

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = uiutils.ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = uiutils.ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
