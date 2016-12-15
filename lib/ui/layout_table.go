package ui

func (p *Page) _layoutBodyTableAlignChild(node *Node) {
	for _node := node.FirstChild; nil != _node; _node = _node.NextSibling {
		p._layoutBodyTableAlignChild(_node)
	}

	if nil != node.UIBlock && true == *node.Display {
		node.UIBlock.Align()
	}
}

func (p *Page) layoutBodyTable(node *Node) (isFallthrough bool) {
	isFallthrough = false
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

	p._layoutBodyTableAlignChild(node)

	nodeDataTable.Body.Align()
	return
}
