package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

func (p *Page) renderBodyPar(node *Node) {
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

	if true == node.isShouldCalculateWidth {
		if true == node.UIBlock.Border {
			node.UIBlock.Width = rw.StringWidth(uiBuffer.Text) + 2
		} else {
			node.UIBlock.Width = rw.StringWidth(uiBuffer.Text)
		}
	}

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = uiutils.ColorToTermuiAttribute(node.ColorFg, uiutils.COLOR_DEFAULT)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = uiutils.ColorToTermuiAttribute(node.ColorBg, uiutils.COLOR_DEFAULT)
	}

	p.BufferersAppend(node, uiBuffer)

	return
}
