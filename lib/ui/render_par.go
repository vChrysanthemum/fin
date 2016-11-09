package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyPar(node *Node) (isFallthrough bool) {
	isFallthrough = false
	nodePar := node.Data.(*NodePar)

	uiBuffer := termui.NewPar(nodePar.RenderText())
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

	node.uiBuffer = uiBuffer

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
