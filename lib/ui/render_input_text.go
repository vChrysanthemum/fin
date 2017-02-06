package ui

import "fin/ui/utils"

func (p *Page) renderBodyInputText(node *Node) {
	uiBuffer := node.Data.(*NodeInputText).Terminal

	p.normalRenderNodeBlock(node)

	uiBuffer.TextFgColor = utils.ColorBlue

	p.BufferersAppend(node, uiBuffer)

	return
}
