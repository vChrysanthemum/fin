package ui

import (
	"inn/ui/editor"

	"github.com/gizak/termui"
)

type NodeEditor struct {
	*Node
	*editor.Editor
}

func (p *Node) InitNodeEditor() *NodeEditor {
	nodeEditor := new(NodeEditor)
	nodeEditor.Node = p
	nodeEditor.Editor = editor.NewEditor()
	p.Data = nodeEditor
	p.Border = false
	p.KeyPress = nodeEditor.KeyPress
	p.FocusMode = nodeEditor.FocusMode
	p.UnFocusMode = nodeEditor.UnFocusMode
	p.ActiveMode = nodeEditor.ActiveMode
	p.UnActiveMode = nodeEditor.UnActiveMode
	return nodeEditor
}

func (p *NodeEditor) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
		return
	}

	p.Editor.Write(keyStr)
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeEditor) FocusMode() {
	p.Node.uiBuffer.(*editor.Editor).Border = true
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_FOCUS_MODE_BORDERFG
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeEditor) UnFocusMode() {
	p.Node.uiBuffer.(*editor.Editor).Border = p.Node.Border
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeEditor) ActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	p.Node.uiBuffer.(*editor.Editor).ActiveMode()
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeEditor) UnActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	p.Node.uiBuffer.(*editor.Editor).UnActiveMode()
	uirender(p.Node.uiBuffer.(termui.Bufferer))
}
