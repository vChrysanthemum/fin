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
	defer func() {
		if len(p.Node.KeyPressHandlers) > 0 {
			for _, v := range p.Node.KeyPressHandlers {
				v.Args = append(v.Args, e)
				v.Handler(p.Node, v.Args...)
			}
		}
	}()

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
			p.Node.refreshUiBufferItems()
			p.Node.uiRender()
		}
		return
	}

	if true == IsVimKeyPressDown(keyStr) {
		if p.SelectedOptionIndex+1 >= len(p.Children) {
		} else {
			p.SelectedOptionIndex++
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

func (p *NodeSelect) NodeDataGetValue() (string, bool) {
	if p.SelectedOptionIndex < 0 || p.SelectedOptionIndex > len(p.Children) {
		return "", false
	} else {
		return p.Children[p.SelectedOptionIndex].Value, true
	}
}

func (p *NodeSelect) NodeDataSetValue(value string) {
	for k, option := range p.Children {
		if option.Value == value {
			p.SelectedOptionIndex = k
			p.Node.refreshUiBufferItems()
			p.Node.uiRender()
			return
		}
	}

	p.SelectedOptionIndex = -1
	p.Node.refreshUiBufferItems()
	p.Node.uiRender()
}

func (p *NodeSelect) AppendOption(value, data string) {
	p.Children = append(p.Children, NodeSelectOption{Value: value, Data: data})
	width := rw.StringWidth(data)
	if width > p.ChildrenMaxStringWidth {
		p.ChildrenMaxStringWidth = width
	}
	p.Node.refreshUiBufferItems()
	p.Node.uiRender()
}

func (p *NodeSelect) ClearOptions() {
	p.SelectedOptionIndex = 0
	p.Children = []NodeSelectOption{}
	p.ChildrenMaxStringWidth = 0
	p.Node.refreshUiBufferItems()
	p.Node.uiRender()
}

func (p *NodeSelect) SetOptionData(value, newData string) {
	for k, v := range p.Children {
		if value == v.Value {
			v.Data = newData
			p.Children[k] = v
			p.Node.refreshUiBufferItems()
			p.Node.uiRender()
			break
		}
	}
}

func (p *NodeSelect) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeSelect) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Node.uiRender()
	}
}
