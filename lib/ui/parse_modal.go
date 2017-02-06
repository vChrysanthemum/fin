package ui

import (
	"log"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyModal(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = nil
	isFallthrough = false

	pageFilepath := ""
	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "src":
			pageFilepath = attr.Val
		}
	}

	if "" == pageFilepath {
		return
	}

	content, err := GetFileContent(pageFilepath)
	if nil != err {
		log.Println(err)
		return
	}

	node := p.newNode(htmlNode)
	err = node.InitNodeModal(string(content))
	if nil != err {
		log.Println(err)
		return
	}

	ret = node
	parentNode.addChild(ret)

	p.Modals[node.ID] = node.Data.(*NodeModal)

	return
}
