package ui

import (
	"container/list"
	"image"
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type NodeKeyPress func(e termui.Event)

type NodeDataGetValuer interface {
	NodeDataGetValue() string
}

type NodeDataSetValueer interface {
	NodeDataSetValue(content string)
}

type NodeDataOnRemover interface {
	NodeDataOnRemove()
}

type NodeDataFocusModer interface {
	NodeDataFocusMode()
}

type NodeDataUnFocusModer interface {
	NodeDataUnFocusMode()
}

type NodeDataActiveModer interface {
	NodeDataActiveMode()
}

type NodeDataUnActiveModer interface {
	NodeDataUnActiveMode()
}

type NodeDataParseAttributer interface {
	NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool)
}

type Node struct {
	Id   string
	page *Page

	ChildrenCount int

	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	FocusTopNode, FocusBottomNode, FocusThisNode *list.Element
	FocusLeftNode, FocusRightNode                *list.Element

	// 是否要渲染子节点
	// 子节点将根据其父节点
	// Node.isShouldTermuiRenderChild 来决定是否 将 node.uiBuffer append 进 page.Bufferers
	// 例: TableTrTd 将用到该字段
	isShouldTermuiRenderChild bool
	isShouldCalculateHeight   bool
	isShouldCalculateWidth    bool

	HtmlAttribute map[string]html.Attribute

	// FocusMode
	isCalledFocusMode    bool
	tmpFocusModeBorder   bool
	tmpFocusModeBorderFg termui.Attribute

	// ActiveMode
	isCalledActiveMode    bool
	tmpActiveModeBorder   bool
	tmpActiveModeBorderFg termui.Attribute

	// 是否可以有交互的 Node
	isWorkNode bool

	// 进入 ActiveMode 所需要触发的 NodeJob
	LuaActiveModeHandlers map[string]NodeJob

	ColorFg string
	ColorBg string

	Display  *bool
	uiBuffer interface{}
	UIBlock  *termui.Block

	HtmlData string
	// NodeData 譬如 NodeTerminal
	Data interface{}

	KeyPress              NodeKeyPress
	KeyPressHandlers      map[string]NodeJob
	KeyPressEnterHandlers map[string]NodeJob

	CursorLocation image.Point
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
	p.UIBlock = nil
	p.Display = new(bool)
	*p.Display = true
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
	ret.LuaActiveModeHandlers = make(map[string]NodeJob, 0)
	ret.KeyPressHandlers = make(map[string]NodeJob, 0)
	ret.KeyPressEnterHandlers = make(map[string]NodeJob, 0)

	ret.isShouldCalculateHeight = true
	ret.isShouldCalculateWidth = true

	ret.HtmlAttribute = make(map[string]html.Attribute)

	ret.CursorLocation = image.Point{-1, -1}
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
	uiutils.UIRender(p.uiBuffer.(termui.Bufferer))
}

func (p *Node) SetCursor(x, y int) {
	p.CursorLocation.X = p.UIBlock.InnerArea.Min.X + x
	p.CursorLocation.Y = p.UIBlock.InnerArea.Min.Y + y
	uiutils.UISetCursor(p.CursorLocation.X, p.CursorLocation.Y)
	p.uiRender()
}

func (p *Node) ResumeCursor() {
	uiutils.UISetCursor(p.CursorLocation.X, p.CursorLocation.Y)
	p.uiRender()
}

func (p *Node) HideCursor() {
	uiutils.UISetCursor(-1, -1)
	p.uiRender()
}
