package ui

import "github.com/gizak/termui"

func (p *Page) renderBodyLineChart(node *Node) {
	uiBuffer := node.uiBuffer.(*termui.LineChart)

	p.normalRenderNodeBlock(node)

	p.BufferersAppend(node, uiBuffer)

	return
}
