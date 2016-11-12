package ui

import (
	"in/ui/editor"

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
	p.KeyPress = nodeEditor.KeyPress
	p.FocusMode = nodeEditor.FocusMode
	p.UnFocusMode = nodeEditor.UnFocusMode
	p.ActiveMode = nodeEditor.ActiveMode
	p.UnActiveMode = nodeEditor.UnActiveMode

	p.uiBuffer = nodeEditor.Editor
	p.uiBlock = &nodeEditor.Editor.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.uiBlock.Height = 10
	p.uiBlock.Border = false

	return nodeEditor
}

func (p *NodeEditor) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
		return
	}

	p.Editor.Write(keyStr)
	p.Node.uiRender()
}

func (p *NodeEditor) afterRenderHandle() {
	p.Editor.AfterRenderHandle()
}

func (p *NodeEditor) FocusMode() {
	p.Node.isCalledFocusMode = true
	p.Node.tmpFocusModeBorder = p.Node.uiBlock.Border
	p.Node.tmpFocusModeBorderFg = p.Node.uiBlock.BorderFg
	p.Node.uiBlock.Border = true
	p.Node.uiBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
	p.Node.uiRender()
}

func (p *NodeEditor) UnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.uiBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.uiBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeEditor) ActiveMode() {
	p.Node.isCalledActiveMode = true
	p.Node.tmpActiveModeBorderFg = p.Node.uiBlock.BorderFg
	p.Node.uiBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeEditor) UnActiveMode() {
	if true == p.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.uiBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Editor.UnActiveMode()
		p.Node.uiRender()
	}
}
