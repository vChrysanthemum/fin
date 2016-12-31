package ui

func (p *Page) _layoutBodyTableChild(node *Node) {
	for _node := node.FirstChild; nil != _node; _node = _node.NextSibling {
		p._layoutBodyTableChild(_node)
	}

	if nil != node.UIBlock && true == node.CheckIfDisplay() {
		node.UIBlock.Align()

		if true == node.isWorkNode {
			p.pushWorkingNode(node)
		}
	}
}

func (p *Page) _layoutBodyTableGetPrevSiblingNodeBottomY(node *Node) int {
	var (
		prevSiblingNode *Node
		nodeDataTable   *NodeTable
		ok              bool
	)
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nodeDataTable, ok = _node.Data.(*NodeTable); true == ok {
			return nodeDataTable.Body.Y + nodeDataTable.Body.Height
		}
		if nil != _node.UIBlock && true == node.CheckIfDisplay() {
			prevSiblingNode = _node
			break
		}
	}

	if nil != prevSiblingNode {
		return prevSiblingNode.UIBlock.Y + prevSiblingNode.UIBlock.Height
	}

	if nil != node.Parent {
		if nil != node.Parent.UIBlock {
			return node.Parent.UIBlock.Y + node.Parent.UIBlock.Height
		} else {
			return p._layoutBodyTableGetPrevSiblingNodeBottomY(node.Parent)
		}
	}

	return 0
}

func (p *Page) layoutBodyTable(node *Node) (isFallthrough bool) {

	isFallthrough = false
	nodeDataTable := node.Data.(*NodeTable)

	nodeDataTable.Body.Y = p._layoutBodyTableGetPrevSiblingNodeBottomY(node)

	nodeDataTable.Body.Align()
	p._layoutBodyTableChild(node)
	nodeDataTable.Body.Align()

	return
}
