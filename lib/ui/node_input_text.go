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

	p.Data = nodeInputText
	p.KeyPress = nodeInputText.KeyPress

	p.uiBuffer = nodeInputText.Editor
	p.UIBlock = &nodeInputText.Editor.Block

	p.isShouldCalculateWidth = false
	p.isShouldCalculateHeight = false
	p.UIBlock.Width = 6
	p.UIBlock.Height = 3
	p.UIBlock.Border = true

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

func (p *NodeInputText) NodeDataGetValue() string {
	if len(p.Editor.Lines) > 0 {
		return string(p.Editor.Lines[0].Data)
	} else {
		return ""
	}
}

func (p *NodeInputText) NodeDataAfterRenderHandle() {
	p.Editor.AfterRenderHandle()
}

func (p *NodeInputText) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeInputText) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeInputText) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	}
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeInputText) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Editor.UnActiveMode()
		p.Node.uiRender()
	}
}
