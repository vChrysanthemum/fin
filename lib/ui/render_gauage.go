package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyGauge(node *Node) (isFallthrough bool) {
	isFallthrough = false
	uiBuffer := node.uiBuffer.(*termui.Gauge)

	node.uiBlock = &uiBuffer.Block
	p.normalRenderNodeBlock(node)

	p.BufferersAppend(node, uiBuffer)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
