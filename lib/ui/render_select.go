package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

func (p *Node) refreshUIBufferItems() {
	nodeSelect := p.Data.(*NodeSelect)

	if nodeSelect.SelectedOptionIndex < nodeSelect.DisplayLinesRange[0] {
		nodeSelect.DisplayLinesRange[1] -= nodeSelect.DisplayLinesRange[0] - nodeSelect.SelectedOptionIndex
		nodeSelect.DisplayLinesRange[0] = nodeSelect.SelectedOptionIndex
	} else if nodeSelect.SelectedOptionIndex >= nodeSelect.DisplayLinesRange[1] {
		nodeSelect.DisplayLinesRange[0] += nodeSelect.SelectedOptionIndex - nodeSelect.DisplayLinesRange[1] + 1
		nodeSelect.DisplayLinesRange[1] = nodeSelect.SelectedOptionIndex + 1
	}

	var items []string
	var str string
	for index, nodeOption := range nodeSelect.Children {
		if nodeSelect.DisplayLinesRange[0] <= index && index < nodeSelect.DisplayLinesRange[1] {
			str = utils.FormatStringWithWidth(nodeOption.Data, nodeSelect.ChildrenMaxStringWidth)
			// str = fmt.Sprintf("%-"+strconv.Itoa(p.UIBlock.InnerArea.Dx())+"s", str)
			if index == nodeSelect.SelectedOptionIndex {
				str = "[" + str + "]" +
					"(fg-" + nodeSelect.SelectedOptionColorFg +
					",bg-" + nodeSelect.SelectedOptionColorBg + ")"
			} else {
				str = "[" + str + "]" + "(fg-" + p.ColorFg + ")"
			}
			items = append(items, str)
		}
	}

	p.UIBuffer.(*termui.List).Items = items
}

func (p *Page) renderBodySelect(node *Node) {
	nodeSelect := node.Data.(*NodeSelect)

	uiBuffer := node.UIBuffer.(*termui.List)

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
		if true == node.UIBlock.Border {
			nodeSelect.DisplayLinesRange[1] = node.UIBlock.Height - 2
		} else {
			nodeSelect.DisplayLinesRange[1] = node.UIBlock.Height
		}
	}

	nodeSelect.refreshUIBufferItems()

	p.BufferersAppend(node, uiBuffer)

	return
}
