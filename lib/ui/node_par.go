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

	p.Data = nodePar

	uiBuffer := termui.NewPar(nodePar.Text)
	p.uiBuffer = uiBuffer
	p.uiBlock = &uiBuffer.Block

	p.uiBlock.Border = false
	p.uiBlock.Width = termui.TermWidth()
	p.uiBlock.Height = -1

	uiBuffer.TextFgColor = COLOR_DEFAULT_TEXT_COLOR_FG

	return nodePar
}

func (p *NodePar) SetText(content string) {
	p.Text = content
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = p.Text
	uiRender(uiBuffer)
}
