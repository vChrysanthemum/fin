package ui

import (
	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

type NodeSelect struct {
	DisableQuit bool
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
	p.GetValue = nodeSelect.GetValue
	p.FocusMode = nodeSelect.FocusMode
	p.UnFocusMode = nodeSelect.UnFocusMode
	p.ActiveMode = nodeSelect.ActiveMode
	p.UnActiveMode = nodeSelect.UnActiveMode

	nodeSelect.SelectedOptionColorFg = COLOR_SELECTED_OPTION_COLORFG
	nodeSelect.SelectedOptionColorBg = COLOR_SELECTED_OPTION_COLORBG
	nodeSelect.Border = true
	nodeSelect.BorderFg = COLOR_DEFAULT_BORDER_FG
	nodeSelect.Width = -1
	nodeSelect.Height = -1

	return nodeSelect
}

type NodeSelectOption struct {
	Value string
	Data  string
}

func (p *NodeSelect) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr && false == p.DisableQuit {
		p.Node.QuitActiveMode()
		return
	}

	if 0 == len(p.Children) {
		return
	}

	if "<up>" == keyStr {
		p.SelectedOptionIndex--
		if p.SelectedOptionIndex < 0 {
			p.SelectedOptionIndex = len(p.Children) - 1
		}
		p.Node.refreshUiBufferItems()
		p.Node.uiRender()
		return
	}

	if "<down>" == keyStr {
		p.SelectedOptionIndex += 1
		if p.SelectedOptionIndex >= len(p.Children) {
			p.SelectedOptionIndex = 0
		}
		p.Node.refreshUiBufferItems()
		p.Node.uiRender()
		return
	}

	if "<enter>" == keyStr {
		if len(p.Node.KeyPressEnterHandlers) > 0 {
			p.Node.JobHanderLocker.RLock()
			defer p.Node.JobHanderLocker.RUnlock()
			for _, v := range p.Node.KeyPressEnterHandlers {
				v.Handler(p.Node, v.Args...)
			}
		}
		return
	}
}

func (p *NodeSelect) GetValue() string {
	nodeSelectOption := p.Children[p.SelectedOptionIndex]
	return nodeSelectOption.Value
}

func (p *NodeSelect) AppendOption(value, data string) {
	p.Children = append(p.Children, NodeSelectOption{Value: value, Data: data})
	width := rw.StringWidth(data)
	if width > p.ChildrenMaxStringWidth {
		p.ChildrenMaxStringWidth = width
	}
	p.Node.refreshUiBufferItems()
}

func (p *NodeSelect) ClearOptions() {
	p.SelectedOptionIndex = 0
	p.Children = []NodeSelectOption{}
	p.ChildrenMaxStringWidth = 0
}

func (p *NodeSelect) FocusMode() {
	p.Node.uiBuffer.(*termui.List).Border = true
	p.Node.uiBuffer.(*termui.List).BorderFg = COLOR_FOCUS_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeSelect) UnFocusMode() {
	p.Node.uiBuffer.(*termui.List).Border = p.Node.Border
	p.Node.uiBuffer.(*termui.List).BorderFg = p.Node.BorderFg
	p.Node.uiRender()
}

func (p *NodeSelect) ActiveMode() {
	p.Node.uiBuffer.(*termui.List).BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeSelect) UnActiveMode() {
	p.Node.uiBuffer.(*termui.List).BorderFg = p.Node.BorderFg
	p.Node.uiRender()
}
