package ui

func (p *Page) layoutBodyTable(node *Node) {
	nodeDataTable := node.Data.(*NodeTable)

	var prevSiblingNode *Node
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nil != _node.UIBlock && false == _node.isShouldHide {
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
}
