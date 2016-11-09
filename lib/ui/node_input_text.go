package ui

import (
	"inn/ui/editor"

	"github.com/gizak/termui"
)

type NodeInputText struct {
	*Node
	*editor.Editor
	WaitKeyPressEnterChans []chan bool
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
	inputText.WaitKeyPressEnterChans = make([]chan bool, 0)
	p.Border = true
	p.BorderFg = COLOR_DEFAULT_BORDERFG
	p.Data = inputText
	p.KeyPress = inputText.KeyPress
	p.FocusMode = inputText.FocusMode
	p.UnFocusMode = inputText.UnFocusMode
	p.GetValue = inputText.GetValue
	p.OnKeyPressEnter = inputText.OnKeyPressEnter
	return inputText
}

func (p *NodeInputText) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
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

	p.Editor.Write(keyStr)
	termui.Render(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeInputText) FocusMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_FOCUSMODE_BORDERFG
	termui.Render(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeInputText) UnFocusMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	termui.Render(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeInputText) GetValue() string {
	return string(p.Editor.Lines[0].Data)
}

func (p *NodeInputText) OnKeyPressEnter() {
	c := make(chan bool, 0)
	p.WaitKeyPressEnterChans = append(p.WaitKeyPressEnterChans, c)
	<-c
}
