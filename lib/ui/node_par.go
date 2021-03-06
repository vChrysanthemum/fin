package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type NodePar struct {
	*Node
}

func (p *Node) InitNodePar() {
	nodePar := new(NodePar)
	nodePar.Node = p

	p.Data = nodePar

	uiBuffer := termui.NewPar("")
	p.UIBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = true
	p.UIBlock.Border = false

	uiBuffer.TextFgColor = ColorDefaultTextColorFg

	return
}

func (p *NodePar) NodeDataSetValue(content string) {
	uiBuffer := p.Node.UIBuffer.(*termui.Par)
	uiBuffer.Text = content

	height := utils.CalculateTextHeight(content, uiBuffer.Width)

	if height > uiBuffer.InnerArea.Dy() {
		p.Node.page.ReRender()
	} else {
		p.Node.UIRender()
	}
	return
}

func (p *NodePar) NodeDataGetValue() (string, bool) {
	return p.Node.UIBuffer.(*termui.Par).Text, true
}
