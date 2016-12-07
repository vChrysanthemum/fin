package ui

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

func (p *Page) _renderBodyTabpaneOneTab(nodeTab *Node) {
	var (
		nodeTabData  *NodeTabpaneTab
		nodeTabChild *Node
	)

	nodeTabData = nodeTab.Data.(*NodeTabpaneTab)

	for nodeTabChild = nodeTab.FirstChild; nodeTabChild != nil; nodeTabChild = nodeTabChild.NextSibling {
		if true == nodeTabChild.isShouldHide {
			continue
		}

		nodeTabData.Tab.AddBlocks(nodeTabChild.uiBuffer.(termui.Bufferer))
	}
}

func (p *Page) renderBodyTabpane(node *Node) {
	nodeTabpaneData := node.Data.(*NodeTabpane)

	uiBuffer := node.uiBuffer.(*extra.Tabpane)

	p.normalRenderNodeBlock(node)

	nodeTabpaneData.Tabs = []extra.Tab{}
	for nodeTab := node.FirstChild; nodeTab != nil; nodeTab = nodeTab.NextSibling {
		p._renderBodyTabpaneOneTab(nodeTab)
		nodeTabpaneData.Tabs = append(nodeTabpaneData.Tabs, *(nodeTab.Data.(*NodeTabpaneTab).Tab))
	}

	uiBuffer.SetTabs(nodeTabpaneData.Tabs...)

	p.BufferersAppend(node, node.uiBuffer.(termui.Bufferer))

	p.pushWorkingNode(node)

	return
}

func (p *Page) renderBodyTabpaneTab(node *Node) {
	uiBuffer := node.Parent.uiBuffer.(*extra.Tabpane)
	p.renderingY = uiBuffer.Y
}
