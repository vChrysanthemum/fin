package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

func (p *Page) renderBodyPar(node *Node) {
	uiBuffer := node.UIBuffer.(*termui.Par)

	p.normalRenderNodeBlock(node)

	if true == node.isShouldCalculateHeight {
		if true == node.UIBlock.Border {
			node.UIBlock.Height = utils.CalculateTextHeight(uiBuffer.Text, node.UIBlock.Width) + 2
		} else {
			node.UIBlock.Height = utils.CalculateTextHeight(uiBuffer.Text, node.UIBlock.Width)
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
		uiBuffer.TextFgColor = utils.ColorToTermuiAttribute(node.ColorFg, utils.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = utils.ColorToTermuiAttribute(node.ColorBg, utils.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	return
}
