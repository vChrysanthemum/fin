package ui

func (p *Page) renderBodyCanvas(node *Node) {
	uiBuffer := node.Data.(*NodeCanvas).Canvas

	p.normalRenderNodeBlock(node)

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	return
}
