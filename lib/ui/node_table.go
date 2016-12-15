package ui

import "github.com/gizak/termui"

type NodeTable struct {
	Body *termui.Grid
}

func (p *Node) InitNodeTable() {
	nodeTable := new(NodeTable)
	nodeTable.Body = termui.NewGrid()
	nodeTable.Body.X = 0
	nodeTable.Body.Y = 0
	p.Data = nodeTable

	p.uiBuffer = nodeTable.Body
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true

	return
}

type NodeTableTr struct{}

func (p *Node) InitNodeTableTr() {
	nodeTableTr := new(NodeTableTr)
	p.Data = nodeTableTr
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true
	return
}

type NodeTableTrTd struct {
	Cols   int
	Offset int
}

func (p *Node) InitNodeTableTrTd() {
	nodeTableTrTd := new(NodeTableTrTd)
	p.Data = nodeTableTrTd
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true
	return
}
