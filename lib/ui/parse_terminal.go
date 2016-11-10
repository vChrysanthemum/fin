package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyTerminal(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTerminal := ret.InitNodeTerminal()

	ret.Data = nodeTerminal
	ret.Width = termui.TermWidth()
	ret.Height = 10

	return
}
