package ui

import "github.com/gizak/termui/extra"

type NodeTabpane struct {
	*Node
	Tabs            []*extra.Tab
	TabsNameToIndex map[string]int
}

func (p *Node) InitNodeTabpane() {
	nodeTabpane := new(NodeTabpane)
	nodeTabpane.Node = p
	nodeTabpane.TabsNameToIndex = make(map[string]int)
	p.Data = nodeTabpane
	p.KeyPress = nodeTabpane.KeyPress

	uiBuffer := extra.NewTabpane()
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	uiBuffer.Border = false
	uiBuffer.TabpaneBg = COLOR_DEFAULT_TABPANE_FG
	uiBuffer.TabpaneBg = COLOR_DEFAULT_TABPANE_BG

	p.isWorkNode = true

	return
}

type NodeTabpaneTab struct {
	*Node
	Index       int
	NodeTabpane *NodeTabpane
	Name        string
}

func (p *Node) InitNodeTabpaneTab(parentNode *Node) {
	nodeTabpaneTab := new(NodeTabpaneTab)
	nodeTabpaneTab.Node = p
	nodeTabpaneTab.NodeTabpane, _ = parentNode.Data.(*NodeTabpane)
	p.Data = nodeTabpaneTab

	uiBuffer := extra.NewTab("")
	p.uiBuffer = uiBuffer
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true

	if nil != nodeTabpaneTab.NodeTabpane {
		nodeTabpaneTab.NodeTabpane.Tabs = append(nodeTabpaneTab.NodeTabpane.Tabs, uiBuffer)
	}

	return
}

func (p *NodeTabpane) KeyPress(keyStr string) (isExecNormalKeyPressWork bool) {
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

	uiBuffer := p.Node.uiBuffer.(*extra.Tabpane)

	if 0 == len(p.Tabs) {
		return
	}

	if true == IsVimKeyPressLeft(keyStr) {
		if true == uiBuffer.SetActiveLeft() {
			uiClear(p.Node.UIBlock.Height, -1)
			p.NodeDataUnActiveMode()
			p.Node.page.Render()
			p.Node.page.uiRender()
		}
		return
	}

	if true == IsVimKeyPressRight(keyStr) {
		if true == uiBuffer.SetActiveRight() {
			uiClear(p.Node.UIBlock.Height, -1)
			p.NodeDataUnActiveMode()
			p.Node.page.Render()
			p.Node.page.uiRender()
		}
		return
	}

	return
}

func (p *NodeTabpane) SetActiveTab(name string) {
	if index, ok := p.TabsNameToIndex[name]; true == ok {
		uiBuffer := p.uiBuffer.(*extra.Tabpane)
		if true == uiBuffer.SetActiveTab(index) {
			uiClear(p.Node.UIBlock.Height, -1)
			p.NodeDataUnActiveMode()
			p.Node.page.Render()
			p.Node.page.uiRender()
		}
	}
}

func (p *NodeTabpane) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		if true == p.Node.UIBlock.Border {
			p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
			p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
			p.Node.UIBlock.Border = true
			p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
			p.Node.uiRender()
		} else {
			p.Node.tmpFocusModeBorderFg = p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = COLOR_FOCUS_MODE_BORDERFG
			p.Node.uiRender()
		}
	}
}

func (p *NodeTabpane) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		if true == p.Node.UIBlock.Border {
			p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
			p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
			p.Node.uiRender()
		} else {
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = p.Node.tmpFocusModeBorderFg
			p.Node.uiRender()
		}
	}
}

func (p *NodeTabpane) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		if true == p.Node.UIBlock.Border {
			p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
			p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
			p.Node.uiRender()
		} else {
			p.Node.tmpActiveModeBorderBg = p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = COLOR_ACTIVE_MODE_BORDERBG
			p.Node.uiRender()
		}
	}
}

func (p *NodeTabpane) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		if true == p.Node.UIBlock.Border {
			p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
			p.Node.uiRender()
		} else {
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = p.Node.tmpActiveModeBorderBg
			p.Node.uiRender()
		}
	}
}
