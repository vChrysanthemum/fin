package ui

func (p *Page) normalLayoutNodeBlock(
	node *Node, isParentDeclareAvailWorkNode bool,
) (isFallthrough, isChildNodesAvailWorkNode bool) {
	isFallthrough = true
	isChildNodesAvailWorkNode = true
	if nil == node.UIBlock {
		return
	}

	node.UIBlock.X = p.layoutingX
	node.UIBlock.Y = p.layoutingY

	if nil != node.UIBlock {
		node.UIBlock.Align()
	}

	p.layoutingY = node.UIBlock.Y + node.UIBlock.Height

	return
}
