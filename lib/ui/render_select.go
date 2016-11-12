package ui

import (
	. "in/ui/utils"

	"github.com/gizak/termui"
)

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

	p.normalRenderNodeBlock(node)

	if true == node.isShouldCalculateWidth {
		if true == node.uiBlock.Border {
			node.uiBlock.Width = nodeSelect.ChildrenMaxStringWidth + 2
		} else {
			node.uiBlock.Width = nodeSelect.ChildrenMaxStringWidth
		}
	}

	var height int
	if true == node.uiBlock.Border {
		height = len(nodeSelect.Children) + 2
	} else {
		height = len(nodeSelect.Children)
	}
	if true == node.isShouldCalculateHeight {
		node.uiBlock.Height = height
	} else {
		node.uiBlock.Height = MaxInt(node.uiBlock.Height, height)
	}

	nodeSelect.refreshUiBufferItems()

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
