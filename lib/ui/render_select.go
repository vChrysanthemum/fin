package ui

import "github.com/gizak/termui"

func (p *Node) refreshUiBufferItems() {
	nodeSelect := p.Data.(*NodeSelect)

	items := make([]string, 0)
	var str string
	for index, nodeOption := range nodeSelect.Children {
		str = FormatStringWithWidth(nodeOption.Data, nodeSelect.ChildrenMaxStringWidth)
		if index == nodeSelect.SelectedOptionIndex {
			str = "[" + str + "]" +
				"(fg-" + nodeSelect.SelectedOptionColorFg +
				",bg-" + nodeSelect.SelectedOptionColorBg + ")"
		} else {
			str = "[" + str + "]" + "(fg-" + p.ColorFg + ")"
		}
		items = append(items, str)
	}

	p.uiBuffer.(*termui.List).Items = items
}

func (p *Page) renderBodySelect(node *Node) (isFallthrough bool) {
	isFallthrough = false

	nodeSelect := node.Data.(*NodeSelect)

	uiBuffer := node.uiBuffer.(*termui.List)

	node.uiBlock = &uiBuffer.Block
	p.normalRenderNodeBlock(node)

	if node.Width < 0 {
		if true == node.Border {
			node.Width = nodeSelect.ChildrenMaxStringWidth + 2
		} else {
			node.Width = nodeSelect.ChildrenMaxStringWidth
		}
	}
	uiBuffer.Width = node.Width

	var height int
	if true == node.Border {
		height = len(nodeSelect.Children) + 2
	} else {
		height = len(nodeSelect.Children)
	}
	if node.Height < 0 {
		uiBuffer.Height = height
	} else {
		uiBuffer.Height = maxint(node.Height, height)
	}

	nodeSelect.refreshUiBufferItems()

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
