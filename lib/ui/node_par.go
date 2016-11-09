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

func (p *NodePar) RenderText() {
	if "" != p.Node.ColorFg {
		p.Text = "[" + p.Text + "]" + "(fg-" + p.Node.ColorFg + ")"
	}
}

func (p *NodePar) SetText(content string) {
	p.Text = content
	p.RenderText()
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = p.Text
	termui.Render(uiBuffer)
}
