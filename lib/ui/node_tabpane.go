package ui

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

type NodeTabpane struct {
	*Node
	Tabs []extra.Tab
}

func (p *Node) InitNodeTabpane() {
	nodeTabpane := new(NodeTabpane)
	nodeTabpane.Node = p
	p.Data = nodeTabpane
	p.KeyPress = nodeTabpane.KeyPress

	uiBuffer := extra.NewTabpane()
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block

	uiBuffer.Width = 30

	return
}

type NodeTabpaneTab struct {
	Tab *extra.Tab
}

func (p *Node) InitNodeTabpaneTab() {
	nodeTabpaneTab := new(NodeTabpaneTab)
	nodeTabpaneTab.Tab = extra.NewTab("")
	p.Data = nodeTabpaneTab
	return
}

func (p *NodeTabpane) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
		return
	}

	uiBuffer := p.Node.uiBuffer.(*extra.Tabpane)

	if 0 == len(p.Tabs) {
		return
	}

	if "<left>" == keyStr {
		uiBuffer.SetActiveLeft()
		p.page.Rerender()
		return
	}

	if "<right>" == keyStr {
		uiBuffer.SetActiveRight()
		p.page.Rerender()
		return
	}
}

func (p *NodeTabpane) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeTabpane) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeTabpane) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeTabpane) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Node.uiRender()
	}
}
