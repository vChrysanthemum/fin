package ui

import (
	"fin/ui/utils"

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

	ret.InitNodeTabpaneTab(parentNode)

	nodeTabpaneTab := ret.Data.(*NodeTabpaneTab)
	uiBuffer := ret.UIBuffer.(*extra.Tab)
	for index := 0; index < len(nodeTabpaneTab.NodeTabpane.Tabs); index++ {
		if uiBuffer == nodeTabpaneTab.NodeTabpane.Tabs[index] {
			nodeTabpaneTab.Index = index
			break
		}
	}

	ret.Tab = nodeTabpaneTab

	return
}

func (p *NodeTabpane) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	uiBuffer := p.Node.UIBuffer.(*extra.Tabpane)

	for _, v := range attr {
		switch v.Key {
		case "hidemenu":
			if "true" == v.Val {
				if false == uiBuffer.IsHideMenu {
					isUIChange = true
				}
				uiBuffer.IsHideMenu = true
			} else {
				if true == uiBuffer.IsHideMenu {
					isUIChange = true
				}
				uiBuffer.IsHideMenu = false
			}

		case "tabpanefg":
			isUIChange = true
			uiBuffer.TabpaneFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabpaneFg)

		case "tabpanebg":
			isUIChange = true
			uiBuffer.TabpaneBg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabpaneBg)

		case "tabfg":
			isUIChange = true
			uiBuffer.TabFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabFg)

		case "tabbg":
			isUIChange = true
			uiBuffer.TabBg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabBg)

		case "activetabfg":
			isUIChange = true
			uiBuffer.ActiveTabFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultActiveTabFg)

		case "activetabbg":
			isUIChange = true
			uiBuffer.ActiveTabBg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultActiveTabBg)
		}
	}

	return
}

func (p *NodeTabpaneTab) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	uiBuffer := p.Node.UIBuffer.(*extra.Tab)

	for _, v := range attr {
		switch v.Key {
		case "label":
			uiBuffer.SetLabel(v.Val)
			isUIChange = true

		case "name":
			if nil == p.NodeTabpane || p.Name == v.Val {
				continue
			}
			delete(p.NodeTabpane.TabsNameToIndex, p.Name)
			p.Name = v.Val
			p.NodeTabpane.TabsNameToIndex[p.Name] = p.Index
		}
	}

	return
}
