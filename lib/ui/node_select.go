package ui

import (
	"github.com/gizak/termui"
)

type NodeSelect struct {
	DisableQuit bool
	*Node
	SelectedOptionColorFg  string
	SelectedOptionColorBg  string
	SelectedOptionIndex    int
	Children               []NodeSelectOption
	ChildrenMaxStringWidth int
	WaitKeyPressEnterChans []chan bool
}

func (p *Node) InitNodeSelect() *NodeSelect {
	nodeSelect := new(NodeSelect)
	nodeSelect.Node = p
	nodeSelect.Children = make([]NodeSelectOption, 0)
	nodeSelect.WaitKeyPressEnterChans = make([]chan bool, 0)
	p.Data = nodeSelect
	p.KeyPress = nodeSelect.KeyPress
	p.GetValue = nodeSelect.GetValue
	p.OnKeyPressEnter = nodeSelect.OnKeyPressEnter
	p.FocusMode = nodeSelect.FocusMode
	p.UnFocusMode = nodeSelect.UnFocusMode
	p.ActiveMode = nodeSelect.ActiveMode
	p.UnActiveMode = nodeSelect.UnActiveMode

	nodeSelect.SelectedOptionColorFg = COLOR_SELECTED_OPTION_COLORFG
	nodeSelect.SelectedOptionColorBg = COLOR_SELECTED_OPTION_COLORBG
	nodeSelect.Border = true
	nodeSelect.BorderFg = COLOR_DEFAULT_BORDERFG
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
	if "<escape>" == keyStr && false == p.DisableQuit {
		p.Node.QuitActiveMode()
		return
	}

	if "<up>" == keyStr {
		p.SelectedOptionIndex--
		if p.SelectedOptionIndex < 0 {
			p.SelectedOptionIndex = len(p.Children) - 1
		}
		p.Node.refreshUiBufferItems()
		uirender(p.Node.uiBuffer.(termui.Bufferer))
		return
	}

	if "<down>" == keyStr {
		p.SelectedOptionIndex += 1
		if p.SelectedOptionIndex >= len(p.Children) {
			p.SelectedOptionIndex = 0
		}
		p.Node.refreshUiBufferItems()
		uirender(p.Node.uiBuffer.(termui.Bufferer))
		return
	}

	if "<enter>" == keyStr && len(p.WaitKeyPressEnterChans) > 0 {
		p.Node.QuitActiveMode()
		for _, c := range p.WaitKeyPressEnterChans {
			c <- true
			close(c)
		}
		p.WaitKeyPressEnterChans = make([]chan bool, 0)
		return
	}

}

func (p *NodeSelect) GetValue() string {
	nodeSelectOption := p.Children[p.SelectedOptionIndex]
	return nodeSelectOption.Value
}

func (p *NodeSelect) OnKeyPressEnter() {
	c := make(chan bool, 0)
	p.WaitKeyPressEnterChans = append(p.WaitKeyPressEnterChans, c)
	<-c
}

func (p *NodeSelect) FocusMode() {
	p.Node.uiBuffer.(*termui.List).Border = true
	p.Node.uiBuffer.(*termui.List).BorderFg = COLOR_FOCUS_MODE_BORDERFG
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeSelect) UnFocusMode() {
	p.Node.uiBuffer.(*termui.List).Border = p.Node.Border
	p.Node.uiBuffer.(*termui.List).BorderFg = p.Node.BorderFg
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeSelect) ActiveMode() {
	p.Node.uiBuffer.(*termui.List).BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeSelect) UnActiveMode() {
	p.Node.uiBuffer.(*termui.List).BorderFg = p.Node.BorderFg
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}
