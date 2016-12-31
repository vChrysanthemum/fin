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

func (p *Page) _layoutBodyTableGetPrevSiblingNode(node *Node) *Node {
	var prevSiblingNode *Node
	for _node := node.PrevSibling; nil != _node; _node = _node.PrevSibling {
		if nil != _node.UIBlock && true == node.CheckIfDisplay() {
			prevSiblingNode = _node
			break
		}
	}

	if nil != prevSiblingNode {
		return prevSiblingNode
	}

	if nil != node.Parent {
		if nil != node.Parent.UIBlock {
			return node.Parent
		} else {
			return p._layoutBodyTableGetPrevSiblingNode(node.Parent)
		}
	}

	return nil
}

func (p *Page) layoutBodyTable(node *Node) (isFallthrough bool) {

	isFallthrough = false
	nodeDataTable := node.Data.(*NodeTable)

	var prevSiblingNode *Node = p._layoutBodyTableGetPrevSiblingNode(node)
	if nil != prevSiblingNode {
		nodeDataTable.Body.Y = prevSiblingNode.UIBlock.Y + prevSiblingNode.UIBlock.Height
	} else {
		nodeDataTable.Body.Y = 0
	}

	nodeDataTable.Body.Align()
	p._layoutBodyTableChild(node)
	nodeDataTable.Body.Align()

	return
}
