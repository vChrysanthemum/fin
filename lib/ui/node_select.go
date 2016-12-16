package ui

import (
	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

type NodeSelect struct {
	*Node
	SelectedOptionColorFg  string
	SelectedOptionColorBg  string
	SelectedOptionIndex    int
	Children               []NodeSelectOption
	ChildrenMaxStringWidth int
	DisplayLinesRange      [2]int
	DisableQuit            bool
}

func (p *Node) InitNodeSelect() {
	nodeSelect := new(NodeSelect)
	nodeSelect.Node = p
	nodeSelect.Children = make([]NodeSelectOption, 0)
	p.Data = nodeSelect
	p.KeyPress = nodeSelect.KeyPress

	nodeSelect.SelectedOptionColorFg = COLOR_SELECTED_OPTION_COLORFG
	nodeSelect.SelectedOptionColorBg = COLOR_SELECTED_OPTION_COLORBG

	nodeSelect.DisplayLinesRange = [2]int{0, 0}

	p.Data = nodeSelect

	uiBuffer := termui.NewList()
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = true
	uiBuffer.Border = true
	uiBuffer.BorderFg = COLOR_DEFAULT_BORDER_FG

	p.isWorkNode = true

	return
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

	if true == IsVimKeyPressUp(keyStr) {
		if p.SelectedOptionIndex-1 < 0 {
		} else {
			p.SelectedOptionIndex--
			if p.SelectedOptionIndex < p.DisplayLinesRange[0] {
				p.DisplayLinesRange[0] = p.DisplayLinesRange[0] - 1
				p.DisplayLinesRange[1] = p.DisplayLinesRange[1] - 1
			}
			p.Node.refreshUiBufferItems()
			p.Node.uiRender()
		}
		return
	}

	if true == IsVimKeyPressDown(keyStr) {
		if p.SelectedOptionIndex+1 >= len(p.Children) {
		} else {
			p.SelectedOptionIndex++
			if p.SelectedOptionIndex >= p.DisplayLinesRange[1] {
				p.DisplayLinesRange[1] = p.DisplayLinesRange[1] + 1
				p.DisplayLinesRange[0] = p.DisplayLinesRange[0] + 1
			}
			p.Node.refreshUiBufferItems()
			p.Node.uiRender()
		}
		return
	}

	if "<enter>" == keyStr {
		if len(p.Node.KeyPressEnterHandlers) > 0 {
			for _, v := range p.Node.KeyPressEnterHandlers {
				v.Handler(p.Node, v.Args...)
			}
		}
		return
	}
}

func (p *NodeSelect) NodeDataGetValue() string {
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

func (p *NodeSelect) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Node.uiRender()
	}
}
