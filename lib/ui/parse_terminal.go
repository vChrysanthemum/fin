package ui

import (
	. "in/ui/utils"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyTerminal(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	ret.InitNodeTerminal()

	return
}

func (p *NodeTerminal) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool) {
	isUIChange = false
	isNeedRerenderPage = false

	for _, v := range attr {
		switch v.Key {
		case "active_borderfg":
			isUIChange = true
			p.ActiveModeBorderColor = ColorToTermuiAttribute(v.Val, COLOR_ACTIVE_MODE_BORDERFG)
		}
	}

	return
}
