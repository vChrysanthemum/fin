package ui

import "github.com/gizak/termui"

type NodeSelect struct {
	*Node
	SelectedOptionColorFg  string
	SelectedOptionColorBg  string
	SelectedOptionIndex    int
	Children               []NodeSelectOption
	ChildrenMaxStringWidth int
}

func (p *Node) InitNodeSelect() *NodeSelect {
	nodeSelect := new(NodeSelect)
	nodeSelect.Node = p
	nodeSelect.Children = make([]NodeSelectOption, 0)
	p.Data = nodeSelect
	p.KeyPress = nodeSelect.KeyPress

	nodeSelect.SelectedOptionColorFg = COLOR_SELECTED_OPTION_COLORFG
	nodeSelect.SelectedOptionColorBg = COLOR_SELECTED_OPTION_COLORBG
	nodeSelect.Border = true
	nodeSelect.Width = -1
	nodeSelect.Height = -1

	return nodeSelect
}

type NodeSelectOption struct {
	Data  string
	Value string
}

func (p *NodeSelect) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.page.ActiveNode = nil
		return
	}

	if "<up>" == keyStr {
		p.SelectedOptionIndex--
		if p.SelectedOptionIndex < 0 {
			p.SelectedOptionIndex = len(p.Children) - 1
		}
		p.Node.refreshUiBufferItems()
		termui.Render(p.Node.uiBuffer.(termui.Bufferer))
		return
	}

	if "<down>" == keyStr {
		p.SelectedOptionIndex += 1
		if p.SelectedOptionIndex >= len(p.Children) {
			p.SelectedOptionIndex = 0
		}
		p.Node.refreshUiBufferItems()
		termui.Render(p.Node.uiBuffer.(termui.Bufferer))
		return
	}

}
