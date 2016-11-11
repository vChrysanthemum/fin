package ui

import "github.com/gizak/termui"

type NodeGauge struct {
	*Node
}

func (p *Node) InitNodeGauge() *NodeGauge {
	nodeGauge := new(NodeGauge)
	nodeGauge.Node = p
	p.Data = nodeGauge
	return nodeGauge
}

func (p *NodeGauge) SetPercent(percent int) {
	uiBuffer := p.Node.uiBuffer.(*termui.Gauge)
	uiBuffer.Percent = percent
	uiRender(uiBuffer)
}
