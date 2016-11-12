package ui

import (
	. "in/ui/utils"
	"strconv"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyGauge(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	ret.InitNodeGauge()

	return
}

func (p *NodeGauge) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool) {
	isUIChange = false
	isNeedRerenderPage = false
	uiBuffer := p.Node.uiBuffer.(*termui.Gauge)

	for _, v := range attr {
		switch v.Key {
		case "barcolor":
			isUIChange = true
			uiBuffer.BarColor = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_BARCOLOR)

		case "percentcolor":
			isUIChange = true
			uiBuffer.PercentColor = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR)

		case "percentcolor_highlighted":
			isUIChange = true
			uiBuffer.PercentColorHighlighted =
				ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED)

		case "percent":
			isUIChange = true
			uiBuffer.Percent, _ = strconv.Atoi(v.Val)
		}
	}

	return
}
