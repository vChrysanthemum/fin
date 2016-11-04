package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type NodeKeyPress func(e termui.Event)

type Node struct {
	page *Page

	ChildrenCount int

	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	// 是否要渲染子节点
	// 子节点将根据其父节点
	// Node.isShouldTermuiRenderChild 来决定是否 将 node.uiBuffer append 进 page.bufferers
	// 例: TableTrTd 将用到该字段
	isShouldTermuiRenderChild bool

	uiBuffer interface{}

	// TODO 重构代码
	// 这里用了绕了个弯
	// 这里利用 Height Width 为-1时，则由 render阶段来计算
	Width, Height int

	ColorFg     string
	BorderLabel string
	Border      bool
	HtmlData    string
	Data        interface{}
	KeyPress    NodeKeyPress
}

type NodeBody struct{}

func (p *Node) InitNodeBody() *NodeBody {
	nodeBody := new(NodeBody)
	p.Data = nodeBody
	return nodeBody
}

type NodeDiv struct{}

func (p *Node) InitNodeDiv() *NodeDiv {
	nodeDiv := new(NodeDiv)
	p.Data = nodeDiv
	return nodeDiv
}

type NodePar struct {
	Text string
}

func (p *Node) InitNodePar() *NodePar {
	nodePar := new(NodePar)
	p.Data = nodePar
	return nodePar
}

type NodeTable struct {
	NodeTrList []NodeTableTr
}

func (p *Node) InitNodeTable() *NodeTable {
	nodeTable := new(NodeTable)
	p.Data = nodeTable
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

func (p *Node) addChild(child *Node) {
	if nil == p {
		return
	}

	child.Parent = p
	child.Parent.ChildrenCount += 1
	child.FirstChild = nil
	child.LastChild = nil
	child.PrevSibling = nil
	child.NextSibling = nil

	if nil == p.FirstChild {
		p.FirstChild = child
	}

	if nil != p.LastChild {
		p.LastChild.NextSibling = child
		child.PrevSibling = p.LastChild
	}

	p.LastChild = child
}

func (p *Page) newNode(htmlNode *html.Node) *Node {
	ret := new(Node)
	ret.page = p
	ret.HtmlData = htmlNode.Data
	ret.Width = 1
	ret.Height = 1
	return ret
}
