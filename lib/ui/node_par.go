package ui

import (
	. "in/ui/utils"

	"github.com/gizak/termui"
)

type NodePar struct {
	*Node
}

func (p *Node) InitNodePar() *NodePar {
	nodePar := new(NodePar)
	nodePar.Node = p
	p.Data = nodePar
	p.SetText = nodePar.SetText

	p.Data = nodePar

	uiBuffer := termui.NewPar("")
	p.uiBuffer = uiBuffer
	p.uiBlock = &uiBuffer.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = true
	p.uiBlock.Border = false

	uiBuffer.TextFgColor = COLOR_DEFAULT_TEXT_COLOR_FG

	return nodePar
}

func (p *NodePar) SetText(content string) (isNeedRerenderPage bool) {
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = content

	height := CalculateTextHeight(content, uiBuffer.Width)

	if height != uiBuffer.Height {
		isNeedRerenderPage = true
	}
	return
}
