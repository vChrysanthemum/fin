package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.uiBuffer.(*termui.Par)

	node.uiBlock = &uiBuffer.Block
	p.normalRenderNodeBlock(node)

	if node.Height < 0 {
		if true == node.Border {
			node.Height = 3
		} else {
			node.Height = 1
		}
	}
	uiBuffer.Height = node.Height

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
