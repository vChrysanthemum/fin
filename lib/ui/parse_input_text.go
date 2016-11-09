package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyInputText(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	inputText := ret.InitNodeInputText()

	ret.Data = inputText
	ret.Width = termui.TermWidth()
	ret.Height = 1

	return
}
