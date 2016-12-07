package ui

import "github.com/gizak/termui"

func (p *Page) normalRenderNodeBlock(node *Node) {
	if nil == node.UIBlock {
		return
	}

	if true == node.isShouldCalculateWidth {
		node.UIBlock.Width = termui.TermWidth()
	}

	node.UIBlock.X = p.renderingX
	node.UIBlock.Y = p.renderingY
}
