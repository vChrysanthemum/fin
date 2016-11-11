package ui

import (
	"in/ui/editor"

	"github.com/gizak/termui"
)

type NodeInputText struct {
	*Node
	*editor.Editor
}

func (p *Node) InitNodeInputText() *NodeInputText {
	nodeInputText := new(NodeInputText)
	nodeInputText.Node = p
	nodeInputText.Editor = editor.NewEditor()
	nodeInputText.Editor.Border = true
	nodeInputText.Editor.BorderTop = false
	nodeInputText.Editor.BorderLeft = false
	nodeInputText.Editor.BorderRight = false
	nodeInputText.Editor.BorderBottom = true

	p.Data = nodeInputText
	p.KeyPress = nodeInputText.KeyPress
	p.GetValue = nodeInputText.GetValue
	p.FocusMode = nodeInputText.FocusMode
	p.UnFocusMode = nodeInputText.UnFocusMode
	p.ActiveMode = nodeInputText.ActiveMode
	p.UnActiveMode = nodeInputText.UnActiveMode

	p.uiBuffer = nodeInputText.Editor
	p.uiBlock = &nodeInputText.Editor.Block

	p.uiBlock.Width = termui.TermWidth()
	p.uiBlock.Height = 1

	return nodeInputText
}

func (p *NodeInputText) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
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

	p.Editor.Write(keyStr)
	p.Node.uiRender()
}

func (p *NodeInputText) GetValue() string {
	if len(p.Editor.Lines) > 0 {
		return string(p.Editor.Lines[0].Data)
	} else {
		return ""
	}
}

func (p *NodeInputText) FocusMode() {
	p.Node.tmpBorder = p.Node.uiBlock.Border
	p.Node.tmpBorderFg = p.Node.uiBlock.BorderFg
	p.Node.uiBlock.Border = true
	p.Node.uiBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeInputText) UnFocusMode() {
	p.Node.uiBlock.Border = p.Node.tmpBorder
	p.Node.uiBlock.BorderFg = p.Node.tmpBorderFg
	p.Node.uiRender()
}

func (p *NodeInputText) ActiveMode() {
	p.Node.tmpBorderFg = p.Node.uiBlock.BorderFg
	p.Node.uiBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeInputText) UnActiveMode() {
	p.Node.uiBlock.BorderFg = p.Node.tmpBorderFg
	p.Editor.UnActiveMode()
	p.Node.uiRender()
}
