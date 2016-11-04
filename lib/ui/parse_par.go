package ui

import "golang.org/x/net/html"

func (p *Page) parseBodyPar(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = false

	nodePar := ret.InitNodePar()

	if nil != htmlNode.FirstChild {
		nodePar.Text = htmlNode.FirstChild.Data
	}

	ret.Data = nodePar
	ret.Height = -1

	return
}
