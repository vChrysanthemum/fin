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

	return nodeCanvas
}
