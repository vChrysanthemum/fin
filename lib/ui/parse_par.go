package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyPar(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	ret.InitNodePar()

	if nil != htmlNode.FirstChild {
		uiBuffer := ret.UIBuffer.(*termui.Par)
		uiBuffer.Text = htmlNode.FirstChild.Data
	}

	return
}
