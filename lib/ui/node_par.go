package ui

import "github.com/gizak/termui"

type NodePar struct {
	*Node
	Text string
}

func (p *Node) InitNodePar() *NodePar {
	nodePar := new(NodePar)
	nodePar.Node = p
	p.Data = nodePar
	p.SetText = nodePar.SetText
	return nodePar
}

func (p *NodePar) SetText(content string) {
	p.Text = content
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = p.Text
	uirender(uiBuffer)
}
