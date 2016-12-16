package ui

import "github.com/gizak/termui/extra"

func (p *Page) layoutBodyTabpane(
	node *Node, isParentDeclareAvailWorkNode bool,
) (isFallthrough, isChildNodesAvailWorkNode bool) {

	isFallthrough = true
	isChildNodesAvailWorkNode = true
	var prevSiblingNode *Node
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nil != _node.UIBlock && true == *_node.Display {
			prevSiblingNode = _node
			break
		}
	}

	if nil != prevSiblingNode {
		node.UIBlock.Y = prevSiblingNode.UIBlock.Y + prevSiblingNode.UIBlock.Height
	} else {
		node.UIBlock.Y = 0
	}

	if true == node.UIBlock.Border {
		node.UIBlock.Height = 3
	} else {
		node.UIBlock.Height = 1
	}

	p.layoutingY = node.UIBlock.Y + node.UIBlock.Height
	node.UIBlock.Align()

	return
}

func (p *Page) layoutBodyTabpaneTab(
	node *Node, isParentDeclareAvailWorkNode bool,
) (isFallthrough, isChildNodesAvailWorkNode bool) {

	isFallthrough = true

	p.layoutingY = node.Parent.UIBlock.Y + node.Parent.UIBlock.Height

	parentUiBuffer := node.Parent.uiBuffer.(*extra.Tabpane)
	nodeDataTab := node.Data.(*NodeTabpaneTab)

	if nodeDataTab.Index == parentUiBuffer.GetActiveIndex() {
		isChildNodesAvailWorkNode = true
	} else {
		isChildNodesAvailWorkNode = false
	}

	return
}
