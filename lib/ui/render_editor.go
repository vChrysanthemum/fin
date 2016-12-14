package ui

import (
	uiutils "in/ui/utils"

	"github.com/gizak/termui"
)

func (p *Page) renderBodyEditor(node *Node) {
	uiBuffer := node.Data.(*NodeEditor).Editor

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = uiutils.ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = uiutils.ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	return
}
