package ui

import "github.com/gizak/termui"

type NodeGauge struct {
	*Node
}

func (p *Node) InitNodeGauge() *NodeGauge {
	nodeGauge := new(NodeGauge)
	nodeGauge.Node = p
	p.Data = nodeGauge

	p.Data = nodeGauge

	uiBuffer := termui.NewGauge()
	p.uiBuffer = uiBuffer
	p.uiBlock = &uiBuffer.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.uiBlock.Height = 3

	uiBuffer.BarColor = COLOR_DEFAULT_GAUGE_BARCOLOR
	uiBuffer.PercentColor = COLOR_DEFAULT_GAUGE_PERCENTCOLOR
	uiBuffer.PercentColorHighlighted = COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED

	return nodeGauge
}

func (p *NodeGauge) SetPercent(percent int) {
	uiBuffer := p.Node.uiBuffer.(*termui.Gauge)
	uiBuffer.Percent = percent
	p.Node.uiRender()
}
