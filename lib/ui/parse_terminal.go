package ui

import "golang.org/x/net/html"

func (p *Page) parseBodyTerminal(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTerminal := ret.InitNodeTerminal()

	for _, v := range htmlNode.Attr {
		switch v.Key {
		case "active_borderfg":
			nodeTerminal.ActiveModeBorderColor = ColorToTermuiAttribute(v.Val, COLOR_ACTIVE_MODE_BORDERFG)
		}
	}

	return
}
