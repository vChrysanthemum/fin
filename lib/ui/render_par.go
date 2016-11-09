package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	nodePar := node.Data.(*NodePar)

	var uiBuffer *termui.Par
	if nil != node.uiBuffer {
		uiBuffer = node.uiBuffer.(*termui.Par)
	} else {
		uiBuffer = termui.NewPar(nodePar.Text)
	}

	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border

	uiBuffer.Width = node.Width

	if node.Height < 0 {
		if true == node.Border {
			node.Height = 3
		} else {
			node.Height = 1
		}
	}
	uiBuffer.Height = node.Height

	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY
	if "" != node.ColorFg {
		uiBuffer.TextFgColor = ColorToTermuiAttribute(node.ColorFg, termui.ColorDefault)
	}
	if "" != node.ColorBg {
		uiBuffer.TextBgColor = ColorToTermuiAttribute(node.ColorBg, termui.ColorDefault)
	}

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
