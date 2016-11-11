package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func (p *Page) parseBodyEditor(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeEditor := ret.InitNodeEditor()

	ret.Data = nodeEditor
	ret.Width = termui.TermWidth()
	ret.Height = 10

	ret.uiBuffer = nodeEditor.Editor

	return
}
