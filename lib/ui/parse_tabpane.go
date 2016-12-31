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

	ret.InitNodeTabpaneTab(parentNode)

	nodeTabpaneTab := ret.Data.(*NodeTabpaneTab)
	uiBuffer := ret.uiBuffer.(*extra.Tab)
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

	uiBuffer := p.Node.uiBuffer.(*extra.Tabpane)

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

func (p *NodeTabpaneTab) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	uiBuffer := p.Node.uiBuffer.(*extra.Tab)

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
