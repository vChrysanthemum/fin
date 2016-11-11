package ui

import (
	"strconv"

	"github.com/gizak/termui"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyGauge(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	ret.InitNodeGauge()

	uiBuffer := ret.uiBuffer.(*termui.Gauge)

	for _, v := range htmlNode.Attr {
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
