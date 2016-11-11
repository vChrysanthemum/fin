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
	return nodeCanvas
}
