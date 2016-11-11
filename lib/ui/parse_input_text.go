package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyInputText(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeInputText := ret.InitNodeInputText()

	ret.Data = nodeInputText
	ret.Width = termui.TermWidth()
	ret.Height = 1

	ret.uiBuffer = nodeInputText.Editor

	return
}
