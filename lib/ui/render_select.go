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
		if true == node.UIBlock.Border {
			node.UIBlock.Width = nodeSelect.ChildrenMaxStringWidth + 2
		} else {
			node.UIBlock.Width = nodeSelect.ChildrenMaxStringWidth
		}
		node.UIBlock.Width += node.UIBlock.PaddingLeft
		node.UIBlock.Width += node.UIBlock.PaddingRight
	}

	var height int
	if true == node.UIBlock.Border {
		height = len(nodeSelect.Children) + 2
	} else {
		height = len(nodeSelect.Children)
	}
	height += node.UIBlock.PaddingTop
	height += node.UIBlock.PaddingBottom

	nodeSelect.refreshUiBufferItems()

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
