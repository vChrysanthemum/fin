package ui

func (p *Page) normalLayoutNodeBlock(node *Node) {
	if nil == node.UIBlock {
		return
	}

	node.UIBlock.X = p.layoutingX
	node.UIBlock.Y = p.layoutingY

	p.layoutingY = node.UIBlock.Y + node.UIBlock.Height
}
