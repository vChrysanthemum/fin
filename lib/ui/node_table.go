package ui

import "github.com/gizak/termui"

type NodeTable struct {
	NodeTrList []NodeTableTr
	Body       *termui.Grid
}

func (p *Node) InitNodeTable() *NodeTable {
	nodeTable := new(NodeTable)
	nodeTable.Body = termui.NewGrid()
	nodeTable.Body.X = 0
	nodeTable.Body.Y = 0
	p.Data = nodeTable

	p.uiBuffer = nodeTable.Body
	p.UIBlock = nil

	return nodeTable
}

type NodeTableTr struct{}

func (p *Node) InitNodeTableTr() *NodeTableTr {
	nodeTableTr := new(NodeTableTr)
	p.Data = nodeTableTr
	return nodeTableTr
}

type NodeTableTrTd struct {
	Cols   int
	Offset int
}

func (p *Node) InitNodeTableTrTd() *NodeTableTrTd {
	nodeTableTrTd := new(NodeTableTrTd)
	p.Data = nodeTableTrTd
	return nodeTableTrTd
}
