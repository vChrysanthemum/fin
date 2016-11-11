package ui

import (
	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type NodeKeyPress func(e termui.Event)
type NodeFocusMode func()
type NodeUnFocusMode func()
type NodeActiveMode func()
type NodeUnActiveMode func()
type NodeGetValue func() string
type NodeSetText func(content string)
type NodeOnKeyPressEnter func()
type NodeOnRemove func()

type Node struct {
	Id   string
	page *Page

	ChildrenCount int

	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	// 是否要渲染子节点
	// 子节点将根据其父节点
	// Node.isShouldTermuiRenderChild 来决定是否 将 node.uiBuffer append 进 page.Bufferers
	// 例: TableTrTd 将用到该字段
	isShouldTermuiRenderChild bool

	uiBuffer interface{}
	uiBlock  *termui.Block

	// TODO 重构代码
	// 这里用了绕了个弯
	// 这里利用 Height Width 为-1时，则由 render阶段来计算
	Width, Height int

	ColorFg       string
	ColorBg       string
	BorderLabelFg termui.Attribute
	BorderLabel   string
	Border        bool
	BorderFg      termui.Attribute
	HtmlData      string
	Data          interface{}

	KeyPress     NodeKeyPress
	FocusMode    NodeFocusMode
	UnFocusMode  NodeUnFocusMode
	ActiveMode   NodeActiveMode
	UnActiveMode NodeUnActiveMode

	SetText         NodeSetText
	GetValue        NodeGetValue
	OnKeyPressEnter NodeOnKeyPressEnter
	OnRemove        NodeOnRemove
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

func (p *Page) newNode(htmlNode *html.Node) *Node {
	ret := new(Node)
	ret.page = p
	ret.HtmlData = htmlNode.Data
	ret.Width = 1
	ret.Height = 1
	ret.BorderLabelFg = COLOR_DEFAULT_BORDER_LABEL_FG
	ret.BorderFg = COLOR_DEFAULT_BORDER_FG
	return ret
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

func (p *Node) render() {
	if nil == p.uiBuffer {
		return
	}
	uirender(p.uiBuffer.(termui.Bufferer))
}
