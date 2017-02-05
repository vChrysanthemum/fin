package ui

type NodeInputText struct {
	*Node
	*Terminal
}

func (p *Node) InitNodeInputText() {
	nodeInputText := new(NodeInputText)
	nodeInputText.Node = p
	nodeInputText.Terminal = NewTerminal()
	nodeInputText.Terminal.Cursor.Line = nodeInputText.Terminal.InitNewLine()
	nodeInputText.Terminal.Block.Border = true
	p.Data = nodeInputText
	p.KeyPress = nodeInputText.KeyPress

	p.UIBuffer = nodeInputText.Terminal
	p.UIBlock = &nodeInputText.Terminal.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = false
	p.isShouldCalculateHeight = false
	p.UIBlock.Width = 6
	p.UIBlock.Height = 3
	p.UIBlock.Border = true

	p.isWorkNode = true

	return
}

func (p *NodeInputText) KeyPress(keyStr string) (isExecNormalKeyPressWork bool) {
	isExecNormalKeyPressWork = true
	defer func() {
		if len(p.Node.KeyPressHandlers) > 0 {
			for _, v := range p.Node.KeyPressHandlers {
				v.Args = append(v.Args, keyStr)
				v.Handler(p.Node, v.Args...)
			}
		}
	}()

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

	if "C-8" == keyStr {
		if len(p.Terminal.Cursor.Line.Data) == 0 {
			return
		}
	}

	p.Terminal.Write(keyStr)
	p.Node.uiRender()
	return
}

func (p *NodeInputText) NodeDataGetValue() (string, bool) {
	if len(p.Terminal.Lines) == 0 {
		return "", false
	} else {
		return string(p.Terminal.Lines[0].Data), true
	}
}

func (p *NodeInputText) NodeDataSetValue(content string) {
	uiBuffer := p.Node.UIBuffer.(*Terminal)
	if len(uiBuffer.Lines) > 0 {
		uiBuffer.Lines[0].Data = []byte(content)
	}
	p.Node.uiRender()
	return
}

func (p *NodeInputText) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeInputText) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeInputText) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	}
	p.Terminal.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeInputText) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Terminal.UnActiveMode()
		p.Node.uiRender()
	}
}
