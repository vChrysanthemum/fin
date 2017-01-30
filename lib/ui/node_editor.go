package ui

type NodeEditor struct {
	*Node
	*Editor
}

func (p *Node) InitNodeEditor() {
	nodeEditor := new(NodeEditor)
	nodeEditor.Node = p
	nodeEditor.Editor = NewEditor()

	p.Data = nodeEditor
	p.KeyPress = nodeEditor.KeyPress

	p.UIBuffer = nodeEditor.Editor
	p.UIBlock = &nodeEditor.Editor.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Height = 10
	p.UIBlock.Border = true

	p.isWorkNode = true

	return
}

func (p *NodeEditor) KeyPress(keyStr string) (isExecNormalKeyPressWork bool) {
	isExecNormalKeyPressWork = false
	if true == p.Editor.Write(keyStr) {
		p.Node.QuitActiveMode()
		return
	}
	return
}

func (p *NodeEditor) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeEditor) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeEditor) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	}
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeEditor) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Editor.UnActiveMode()
		p.Node.uiRender()
	}
}
