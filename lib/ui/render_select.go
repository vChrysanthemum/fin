package ui

import "github.com/gizak/termui"

func (p *Page) renderBodySelect(node *Node) (isFallthrough bool) {
	isFallthrough = false
	items := make([]string, 0)

	var (
		str        string
		nodeSelect = node.Data.(*NodeSelect)
	)
	for _, nodeOption := range nodeSelect.Children {
		str = nodeOption.Data
		if "" != node.ColorFg {
			str = "[" + str + "]" + "(fg-" + node.ColorFg + ")"
		}
		items = append(items, str)

	}

	uiBuffer := termui.NewList()
	uiBuffer.BorderLabel = node.BorderLabel
	uiBuffer.Border = node.Border

	if node.Width < 0 {
		if true == node.Border {
			node.Width = nodeSelect.ChildrenMaxStringWidth + 3
		} else {
			node.Width = nodeSelect.ChildrenMaxStringWidth
		}
	}
	uiBuffer.Width = node.Width

	if node.Height < 0 {
		if true == node.Border {
			node.Height = len(items) + 2
		} else {
			node.Height = len(items)
		}
	}
	uiBuffer.Height = node.Height

	uiBuffer.X = p.renderingX
	uiBuffer.Y = p.renderingY
	uiBuffer.Items = items

	uiBuffer.ItemFgColor = termui.ColorCyan

	node.uiBuffer = uiBuffer

	p.bufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	return
}
