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
	p.UIBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Height = 3

	uiBuffer.BarColor = ColorDefaultGaugeBarcolor
	uiBuffer.PercentColor = ColorDefaultGaugePercentcolor
	uiBuffer.PercentColorHighlighted = ColorDefaultGaugePercentColorHighlighted

	return
}
