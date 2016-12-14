package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyGauge(node *Node) {
	uiBuffer := node.uiBuffer.(*termui.Gauge)

	p.normalRenderNodeBlock(node)

	p.BufferersAppend(node, uiBuffer)

	return
}
