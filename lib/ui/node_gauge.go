package ui

import (
	"github.com/gizak/termui"
)

type NodeGauge struct {
	*Node
}

func (p *Node) InitNodeGauge() {
	nodeGauge := new(NodeGauge)
	nodeGauge.Node = p
	p.Data = nodeGauge

	p.Data = nodeGauge

	uiBuffer := termui.NewGauge()
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Height = 3

	uiBuffer.BarColor = COLOR_DEFAULT_GAUGE_BARCOLOR
	uiBuffer.PercentColor = COLOR_DEFAULT_GAUGE_PERCENTCOLOR
	uiBuffer.PercentColorHighlighted = COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED

	return
}
