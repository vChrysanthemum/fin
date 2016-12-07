package ui

import (
	"unicode/utf8"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyTabpane(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	ret.InitNodeTabpane()

	return
}

func (p *Page) parseBodyTabpaneTab(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTabpaneTab := ret.InitNodeTabpaneTab()

	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "label":
			nodeTabpaneTab.Tab.Label = attr.Val
			nodeTabpaneTab.Tab.RuneLen = utf8.RuneCount([]byte(attr.Val))
		}
	}

	ret.isShouldTermuiRenderChild = true

	return
}
