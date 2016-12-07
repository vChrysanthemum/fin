package ui

import (
	uiutils "in/ui/utils"

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
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = true
	p.UIBlock.Border = false

	uiBuffer.TextFgColor = COLOR_DEFAULT_TEXT_COLOR_FG

	return
}

func (p *NodePar) NodeDataSetText(content string) (isNeedRerenderPage bool) {
	uiBuffer := p.Node.uiBuffer.(*termui.Par)
	uiBuffer.Text = content

	height := uiutils.CalculateTextHeight(content, uiBuffer.Width)

	if height > uiBuffer.InnerArea.Dy() {
		isNeedRerenderPage = true
	}
	return
}

func (p *NodePar) NodeDataGetValue() string {
	return p.Node.uiBuffer.(*termui.Par).Text
}
