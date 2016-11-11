package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyCanvas(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	nodeCanvas := ret.InitNodeCanvas()

	ret.Data = nodeCanvas
	ret.Width = termui.TermWidth()
	ret.Height = 10

	ret.uiBuffer = nodeCanvas.Canvas

	return
}
