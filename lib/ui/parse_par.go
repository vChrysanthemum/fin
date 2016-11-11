package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyPar(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	nodePar := ret.InitNodePar()
	uiBuffer := ret.uiBuffer.(*termui.Par)

	if nil != htmlNode.FirstChild {
		nodePar.Text = htmlNode.FirstChild.Data
		uiBuffer.Text = nodePar.Text
	}

	return
}
