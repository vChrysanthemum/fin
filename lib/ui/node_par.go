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

func (p *NodePar) RenderText() string {
	var ret string
	if "" != p.Node.ColorFg {
		ret = "[" + p.Text + "]" + "(fg-" + p.Node.ColorFg + ")"
	} else {
		ret = p.Text
	}
	return ret
}

func (p *NodePar) SetText(content string) {
	p.Text = content
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = p.RenderText()
	termui.Render(uiBuffer)
}
