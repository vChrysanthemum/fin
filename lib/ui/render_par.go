package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.uiBuffer.(*termui.Par)

	p.normalRenderNodeBlock(node)

	if node.uiBlock.Height < 0 {
		if true == node.uiBlock.Border {
			node.uiBlock.Height = 3
		} else {
			node.uiBlock.Height = 1
		}
	}

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
