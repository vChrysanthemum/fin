package ui

import uiutils "fin/ui/utils"

func (p *Page) renderBodyEditor(node *Node) {
	uiBuffer := node.Data.(*NodeEditor).Editor

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = uiutils.ColorToTermuiAttribute(node.ColorFg, uiutils.COLOR_DEFAULT)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = uiutils.ColorToTermuiAttribute(node.ColorBg, uiutils.COLOR_DEFAULT)
	}

	p.BufferersAppend(node, uiBuffer)

	return
}
