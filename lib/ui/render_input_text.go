package ui

import "fin/ui/utils"

func (p *Page) renderBodyInputText(node *Node) {
	uiBuffer := node.Data.(*NodeInputText).Editor

	p.normalRenderNodeBlock(node)

	uiBuffer.TextFgColor = utils.COLOR_BLUE

	p.BufferersAppend(node, uiBuffer)

	return
}
