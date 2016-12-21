package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui/extra"

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

	ret.InitNodeTabpaneTab()

	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "label":
			uiBuffer := ret.uiBuffer.(*extra.Tab)
			uiBuffer.SetLabel(attr.Val)
		}
	}

	ret.isShouldTermuiRenderChild = true

	return
}

func (p *NodeTabpane) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	uiBuffer := p.Node.uiBuffer.(*extra.Tabpane)

	for _, v := range attr {
		switch v.Key {
		case "tabpanefg":
			isUIChange = true
			uiBuffer.TabpaneFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TABPANE_FG)

		case "tabpanebg":
			isUIChange = true
			uiBuffer.TabpaneBg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TABPANE_BG)

		case "tabfg":
			isUIChange = true
			uiBuffer.TabFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TAB_FG)

		case "tabbg":
			isUIChange = true
			uiBuffer.TabBg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TAB_BG)

		case "activetabfg":
			isUIChange = true
			uiBuffer.ActiveTabFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_ACTIVE_TAB_FG)

		case "activetabbg":
			isUIChange = true
			uiBuffer.ActiveTabBg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_ACTIVE_TAB_BG)
		}
	}

	return
}
