package ui

import (
	. "in/ui/utils"
	"strconv"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

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

func (p *NodeGauge) ParseAttribute(attr []html.Attribute) (isNeedRerenderPage bool) {
	isNeedRerenderPage = false
	uiBuffer := p.Node.uiBuffer.(*termui.Gauge)

	for _, v := range attr {
		switch v.Key {
		case "barcolor":
			uiBuffer.BarColor = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_BARCOLOR)
		case "percentcolor":
			uiBuffer.PercentColor = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR)
		case "percentcolor_highlighted":
			uiBuffer.PercentColorHighlighted =
				ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED)
		case "percent":
			uiBuffer.Percent, _ = strconv.Atoi(v.Val)
		}
	}

	return
}
