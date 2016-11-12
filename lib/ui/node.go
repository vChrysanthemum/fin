package ui

import (
	"sync"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type NodeKeyPress func(e termui.Event)
type NodeFocusMode func()
type NodeUnFocusMode func()
type NodeActiveMode func()
type NodeUnActiveMode func()
type NodeGetValue func() string
type NodeSetText func(content string) (isNeedRerenderPage bool)
type NodeOnRemove func()

type NodeDataParseAttributer interface {
	NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool)
}

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
	isShouldCalculateHeight   bool
	isShouldCalculateWidth    bool

	isCalledFocusMode     bool
	tmpFocusModeBorder    bool
	tmpFocusModeBorderFg  termui.Attribute
	isCalledActiveMode    bool
	tmpActiveModeBorder   bool
	tmpActiveModeBorderFg termui.Attribute

	afterRenderHandle func()

	ColorFg string
	ColorBg string

	uiBuffer interface{}
	uiBlock  *termui.Block

	HtmlData string
	Data     interface{}

	KeyPress     NodeKeyPress
	FocusMode    NodeFocusMode
	UnFocusMode  NodeUnFocusMode
	ActiveMode   NodeActiveMode
	UnActiveMode NodeUnActiveMode

	SetText  NodeSetText
	GetValue NodeGetValue
	OnRemove NodeOnRemove

	KeyPressEnterHandlers map[string]NodeJob
	JobHanderLocker       sync.RWMutex
}

type NodeJobHandler func(node *Node, args ...interface{})
type NodeJob struct {
	*Node
	Handler NodeJobHandler
	Args    []interface{}
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
	ret.KeyPressEnterHandlers = make(map[string]NodeJob, 0)

	ret.isShouldCalculateHeight = true
	ret.isShouldCalculateWidth = true
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

func (p *Node) uiRender() {
	if nil == p.uiBuffer {
		return
	}
	termui.Render(p.uiBuffer.(termui.Bufferer))
	if nil != p.afterRenderHandle {
		p.afterRenderHandle()
	}
}
