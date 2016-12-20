package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

func (p *Page) renderBodyTerminal(node *Node) {
	uiBuffer := node.Data.(*NodeTerminal).Editor

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = uiutils.ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = uiutils.ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	return
}
