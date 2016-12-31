package ui

import "github.com/gizak/termui/extra"

func (p *Page) layoutBodyTabpane(node *Node) (isFallthrough bool) {

	isFallthrough = true
	var prevSiblingNode *Node
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nil != _node.UIBlock && true == _node.CheckIfDisplay() {
			prevSiblingNode = _node
			break
		}
	}

	if nil != prevSiblingNode {
		node.UIBlock.Y = prevSiblingNode.UIBlock.Y + prevSiblingNode.UIBlock.Height
	} else {
		node.UIBlock.Y = 0
	}

	uiBuffer := node.uiBuffer.(*extra.Tabpane)
	if true == uiBuffer.IsHideMenu {
		node.UIBlock.Height = 0
	} else {
		if true == node.UIBlock.Border {
			node.UIBlock.Height = 3
		} else {
			node.UIBlock.Height = 1
		}
	}

	p.layoutingY = node.UIBlock.Y + node.UIBlock.Height
	node.UIBlock.Align()

	return
}

func (p *Page) layoutBodyTabpaneTab(node *Node) (isFallthrough bool) {

	isFallthrough = true

	p.layoutingY = node.Parent.UIBlock.Y + node.Parent.UIBlock.Height

	return
}
