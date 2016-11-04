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
	p.KeyPress = nodeEditor.KeyPress
	return nodeEditor
}

func (p *NodeEditor) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.page.ActiveNode = nil
		return
	}

	p.Editor.Write(keyStr)
	termui.Render(p.Node.uiBuffer.(termui.Bufferer))
}
