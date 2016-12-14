package ui

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

func (p *Page) _renderBodyTabpaneOneTab(nodeTab *Node) {
	var (
		nodeDataTab  *NodeTabpaneTab
		nodeTabChild *Node
	)

	nodeDataTab = nodeTab.Data.(*NodeTabpaneTab)

	for nodeTabChild = nodeTab.FirstChild; nodeTabChild != nil; nodeTabChild = nodeTabChild.NextSibling {
		if true == nodeTabChild.isShouldHide {
			continue
		}

		nodeDataTab.Tab.AddBlocks(nodeTabChild.uiBuffer.(termui.Bufferer))
	}
}

func (p *Page) renderBodyTabpane(node *Node) {
	nodeDataTabpane := node.Data.(*NodeTabpane)

	uiBuffer := node.uiBuffer.(*extra.Tabpane)

	p.normalRenderNodeBlock(node)
	node.UIBlock.X = 0
	node.UIBlock.Y = 0

	nodeDataTabpane.Tabs = []extra.Tab{}
	for nodeTab := node.FirstChild; nodeTab != nil; nodeTab = nodeTab.NextSibling {
		p._renderBodyTabpaneOneTab(nodeTab)
		nodeDataTabpane.Tabs = append(nodeDataTabpane.Tabs, *(nodeTab.Data.(*NodeTabpaneTab).Tab))
	}

	uiBuffer.SetTabs(nodeDataTabpane.Tabs...)

	p.BufferersAppend(node, node.uiBuffer.(termui.Bufferer))

	p.pushWorkingNode(node)

	return
}
