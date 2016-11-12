package ui

import "in/ui/canvas"

type NodeCanvas struct {
	*Node
	*canvas.Canvas
}

func (p *Node) InitNodeCanvas() *NodeCanvas {
	nodeCanvas := new(NodeCanvas)
	nodeCanvas.Node = p
	nodeCanvas.Canvas = canvas.NewCanvas()

	p.Data = nodeCanvas

	p.uiBuffer = nodeCanvas.Canvas
	p.uiBlock = &nodeCanvas.Canvas.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.uiBlock.Height = 10
	p.uiBlock.Border = true

	return nodeCanvas
}

func (p *NodeCanvas) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.uiBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.uiBlock.BorderFg
		p.Node.uiBlock.Border = true
		p.Node.uiBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeCanvas) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.uiBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.uiBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeCanvas) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.uiBlock.BorderFg
		p.Node.uiBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
		p.Node.ResumeCursor()
		p.Node.uiRender()
	}
}

func (p *NodeCanvas) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.uiBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Node.HideCursor()
		p.Node.uiRender()
	}
}
