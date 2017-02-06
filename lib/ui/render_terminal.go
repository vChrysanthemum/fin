package ui

import "fin/ui/utils"

func (p *Page) renderBodyTerminal(node *Node) {
	uiBuffer := node.Data.(*NodeTerminal).Terminal

	p.normalRenderNodeBlock(node)

	if "" != node.ColorFg {
		uiBuffer.TextFgColor = utils.ColorToTermuiAttribute(node.ColorFg, utils.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = utils.ColorToTermuiAttribute(node.ColorBg, utils.ColorDefault)
	}

	p.BufferersAppend(node, uiBuffer)

	return
}
