package ui

import (
	uiutils "in/ui/utils"

	"github.com/gizak/termui"
)

func (p *Node) refreshUiBufferItems() {
	nodeSelect := p.Data.(*NodeSelect)

	items := make([]string, 0)
	var str string
	for index, nodeOption := range nodeSelect.Children {
		if index < nodeSelect.DisplayLinesRange[0] {
			continue
		}
		if index >= nodeSelect.DisplayLinesRange[1] {
			continue
		}

		str = uiutils.FormatStringWithWidth(nodeOption.Data, nodeSelect.ChildrenMaxStringWidth)
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

func (p *Page) renderBodySelect(node *Node) {
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

	nodeSelect.DisplayLinesRange[0] = 0

	if true == node.isShouldCalculateHeight {
		node.UIBlock.Height = height
		nodeSelect.DisplayLinesRange[1] = len(nodeSelect.Children)
	} else {
		if 0 == nodeSelect.SelectedOptionIndex {
			if true == node.UIBlock.Border {
				nodeSelect.DisplayLinesRange[1] = node.UIBlock.Height - 2
			} else {
				nodeSelect.DisplayLinesRange[1] = node.UIBlock.Height
			}
		}
	}

	nodeSelect.refreshUiBufferItems()

	p.BufferersAppend(node, uiBuffer)

	p.pushWorkingNode(node)

	p.renderingY = uiBuffer.Y + uiBuffer.Height

	return
}
