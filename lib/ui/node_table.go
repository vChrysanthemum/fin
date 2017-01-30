package ui

import "github.com/gizak/termui"

type NodeTable struct {
	Node *Node
	Body *termui.Grid
}

func (p *Node) InitNodeTable() {
	nodeTable := new(NodeTable)
	nodeTable.Node = p
	nodeTable.Body = termui.NewGrid()
	nodeTable.Body.X = 0
	nodeTable.Body.Y = 0
	p.Data = nodeTable

	p.UIBuffer = nodeTable.Body
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
