package ui

import (
	uiutils "fin/ui/utils"
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

func (p *NodeGauge) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false
	uiBuffer := p.Node.UIBuffer.(*termui.Gauge)

	for _, v := range attr {
		switch v.Key {
		case "barcolor":
			isUIChange = true
			uiBuffer.BarColor = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_BARCOLOR)

		case "percentcolor":
			isUIChange = true
			uiBuffer.PercentColor = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR)

		case "percentcolor_highlighted":
			isUIChange = true
			uiBuffer.PercentColorHighlighted =
				uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED)

		case "percent":
			isUIChange = true
			uiBuffer.Percent, _ = strconv.Atoi(v.Val)
		}
	}

	return
}
