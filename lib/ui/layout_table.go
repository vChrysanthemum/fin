package ui

func (p *Page) _layoutBodyTableChild(node *Node, isParentDeclareAvailWorkNode bool) {
	for _node := node.FirstChild; nil != _node; _node = _node.NextSibling {
		p._layoutBodyTableChild(_node, isParentDeclareAvailWorkNode)
	}

	if nil != node.UIBlock && true == *node.Display {
		node.UIBlock.Align()

		if true == node.isWorkNode && true == isParentDeclareAvailWorkNode {
			p.pushWorkingNode(node)
		}
	}
}

func (p *Page) layoutBodyTable(
	node *Node, isParentDeclareAvailWorkNode bool,
) (isFallthrough, isChildNodesAvailWorkNode bool) {

	isFallthrough = false
	isChildNodesAvailWorkNode = isParentDeclareAvailWorkNode
	nodeDataTable := node.Data.(*NodeTable)

	var prevSiblingNode *Node
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nil != _node.UIBlock && true == *_node.Display {
			prevSiblingNode = _node
			break
		}
	}

	if nil != prevSiblingNode {
		nodeDataTable.Body.Y = prevSiblingNode.UIBlock.Y + prevSiblingNode.UIBlock.Height
	} else {
		nodeDataTable.Body.Y = 0
	}

	nodeDataTable.Body.Align()
	p._layoutBodyTableChild(node, isChildNodesAvailWorkNode)
	nodeDataTable.Body.Align()

	return
}
