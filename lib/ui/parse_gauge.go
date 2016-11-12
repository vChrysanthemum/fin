package ui

import (
	"golang.org/x/net/html"
)

func (p *Page) parseBodyGauge(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	nodeGauge := ret.InitNodeGauge()
	nodeGauge.ParseAttribute(htmlNode.Attr)

	return
}
