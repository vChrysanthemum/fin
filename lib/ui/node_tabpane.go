package ui

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

type NodeTabpane struct {
	*Node
	Tabs []*extra.Tab
}

func (p *Node) InitNodeTabpane() {
	nodeTabpane := new(NodeTabpane)
	nodeTabpane.Node = p
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
	Index int
}

func (p *Node) InitNodeTabpaneTab() {
	nodeTabpaneTab := new(NodeTabpaneTab)
	nodeTabpaneTab.Node = p
	p.Data = nodeTabpaneTab

	uiBuffer := extra.NewTab("")
	p.uiBuffer = uiBuffer
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true

	return
}

func (p *NodeTabpane) KeyPress(e termui.Event) {
	defer func() {
		if len(p.Node.KeyPressHandlers) > 0 {
			for _, v := range p.Node.KeyPressHandlers {
				v.Args = append(v.Args, e)
				v.Handler(p.Node, v.Args...)
			}
		}
	}()

	keyStr := e.Data.(termui.EvtKbd).KeyStr
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
			uiClear(p.Node.UIBlock.Y+1, -1)
			p.Node.page.Render()
			p.Node.page.uiRender()
		}
		return
	}

	if true == IsVimKeyPressRight(keyStr) {
		if true == uiBuffer.SetActiveRight() {
			uiClear(p.Node.UIBlock.Y+1, -1)
			p.Node.page.Render()
			p.Node.page.uiRender()
		}
		return
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
			p.Node.tmpActiveModeBorderFg = p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = COLOR_ACTIVE_MODE_BORDERFG
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
			p.Node.uiBuffer.(*extra.Tabpane).TabpaneBg = p.Node.tmpActiveModeBorderFg
			p.Node.uiRender()
		}
	}
}
