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
	inputText := new(NodeInputText)
	inputText.Node = p
	inputText.Editor = editor.NewEditor()
	inputText.Editor.Border = true
	inputText.Editor.BorderTop = false
	inputText.Editor.BorderLeft = false
	inputText.Editor.BorderRight = false
	inputText.Editor.BorderBottom = true
	p.Border = true
	p.Data = inputText
	p.KeyPress = inputText.KeyPress
	p.GetValue = inputText.GetValue
	p.FocusMode = inputText.FocusMode
	p.UnFocusMode = inputText.UnFocusMode
	p.ActiveMode = inputText.ActiveMode
	p.UnActiveMode = inputText.UnActiveMode
	return inputText
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
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_FOCUS_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeInputText) UnFocusMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	p.Node.uiRender()
}

func (p *NodeInputText) ActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeInputText) UnActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	p.Node.uiRender()
}
