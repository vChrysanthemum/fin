package ui

import (
	"in/ui/editor"

	"github.com/gizak/termui"
)

type NodeEditor struct {
	*Node
	*editor.Editor
}

func (p *Node) InitNodeEditor() {
	nodeEditor := new(NodeEditor)
	nodeEditor.Node = p
	nodeEditor.Editor = editor.NewEditor()

	p.Data = nodeEditor
	p.KeyPress = nodeEditor.KeyPress

	p.uiBuffer = nodeEditor.Editor
	p.UIBlock = &nodeEditor.Editor.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Height = 10
	p.UIBlock.Border = true

	return
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

func (p *NodeEditor) NodeDataAfterRenderHandle() {
	p.Editor.AfterRenderHandle()
}

func (p *NodeEditor) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeEditor) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeEditor) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	}
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeEditor) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Editor.UnActiveMode()
		p.Node.uiRender()
	}
}
