package ui

import (
	"fin/ui/utils"
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
			uiBuffer.BarColor = utils.ColorToTermuiAttribute(v.Val, ColorDefaultGaugeBarcolor)

		case "percentcolor":
			isUIChange = true
			uiBuffer.PercentColor = utils.ColorToTermuiAttribute(v.Val, ColorDefaultGaugePercentcolor)

		case "percentcolor_highlighted":
			isUIChange = true
			uiBuffer.PercentColorHighlighted =
				utils.ColorToTermuiAttribute(v.Val, ColorDefaultGaugePercentColorHighlighted)

		case "percent":
			isUIChange = true
			uiBuffer.Percent, _ = strconv.Atoi(v.Val)
		}
	}

	return
}
