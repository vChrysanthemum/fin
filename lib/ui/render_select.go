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

	uiBuffer := termui.NewList()
	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border
	uiBuffer.BorderFg = node.BorderFg

	if node.Width < 0 {
		if true == node.Border {
			node.Width = nodeSelect.ChildrenMaxStringWidth + 2
		} else {
			node.Width = nodeSelect.ChildrenMaxStringWidth
		}
	}
	uiBuffer.Width = node.Width

	if node.Height < 0 {
		if true == node.Border {
			node.Height = len(nodeSelect.Children) + 2
		} else {
			node.Height = len(nodeSelect.Children)
		}
	}
	uiBuffer.Height = node.Height

	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY

	node.uiBuffer = uiBuffer

	nodeSelect.refreshUiBufferItems()

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	return
}
