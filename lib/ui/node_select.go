package ui

import "github.com/gizak/termui"

type NodeSelect struct {
	*Node
	Children               []NodeSelectOption
	ChildrenMaxStringWidth int
}

func (p *Node) InitNodeSelect() *NodeSelect {
	nodeSelect := new(NodeSelect)
	nodeSelect.Node = p
	nodeSelect.Children = make([]NodeSelectOption, 0)
	p.Data = nodeSelect
	p.KeyPress = nodeSelect.KeyPress
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

	termui.Render(p.Node.uiBuffer.(termui.Bufferer))
}
