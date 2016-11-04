package ui

import "golang.org/x/net/html"

func (p *Page) parseBodyEditor(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeEditor := ret.InitNodeEditor()

	ret.Data = nodeEditor
	ret.Width = 30
	ret.Height = 30

	return
}
