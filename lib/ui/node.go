package ui

import (
	"container/list"
	"fin/ui/utils"
	"image"

	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
	"golang.org/x/net/html"
)

type NodeKeyPress func(keyStr string) (isExecNormalKeyPressWork bool)

type NodeDataGetValuer interface {
	// return:
	//		string	返回结果
	//		bool	结果是否有效
	NodeDataGetValue() (string, bool)
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
	NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool)
}

type Node struct {
	ID   string
	Tab  *NodeTabpaneTab
	page *Page

	ChildrenCount int

	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	FocusTopNode, FocusBottomNode, FocusThisNode *list.Element
	FocusLeftNode, FocusRightNode                *list.Element

	// 是否要渲染子节点
	// 子节点将根据其父节点
	// Node.isShouldTermuiRenderChild 来决定是否 将 node.UIBuffer append 进 page.Bufferers
	// 例: TableTrTd 将用到该字段
	isShouldTermuiRenderChild bool
	isShouldCalculateHeight   bool
	isShouldCalculateWidth    bool
	// 是否在 parse 阶段设置了位置
	isSettedPositionY bool
	isSettedPositionX bool
	// absolute relative default:relative
	Position string

	HTMLAttribute map[string]html.Attribute

	// FocusMode
	isCalledFocusMode    bool
	tmpFocusModeBorder   bool
	tmpFocusModeBorderFg termui.Attribute

	// ActiveMode
	isCalledActiveMode    bool
	tmpActiveModeBorder   bool
	tmpActiveModeBorderFg termui.Attribute
	tmpActiveModeBorderBg termui.Attribute

	// 是否可以有交互的 Node
	isWorkNode bool

	// 进入 ActiveMode 所需要触发的 NodeJob
	LuaActiveModeHandlers map[string]NodeJob

	ColorFg string
	ColorBg string

	Display  *bool
	UIBuffer interface{}
	UIBlock  *termui.Block

	HTMLData string
	// NodeData 譬如 NodeTerminal
	Data interface{}

	KeyPress              NodeKeyPress
	KeyPressHandlers      map[string]NodeJob
	KeyPressEnterHandlers map[string]NodeJob

	Cursor image.Point
}

type NodeJobHandler func(node *Node, args ...interface{})
type NodeJob struct {
	*Node
	Handler NodeJobHandler
	Args    []interface{}
}

type NodeBody struct{}

func (p *Node) addChild(child *Node) {
	if nil == p {
		return
	}

	child.Parent = p
	child.Parent.ChildrenCount++

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

func (p *Node) UIRender() {
	if false == p.CheckIfDisplay() {
		return
	}
	if nil == p.UIBuffer {
		return
	}
	utils.UIRender(p.UIBuffer.(termui.Bufferer))
}

func (p *Node) extractChildsMapIDNodes(ret *map[string]*Node) {
	var childNode *Node
	for childNode = p.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		childNode.extractChildsMapIDNodes(ret)
	}

	if "" != p.ID {
		(*ret)[p.ID] = p
	}
}

func (p *Node) ExtractChildsMapIDNodes() map[string]*Node {
	ret := make(map[string]*Node)
	p.extractChildsMapIDNodes(&ret)
	return ret
}

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
	ret.Display = new(bool)
	*ret.Display = true
	ret.page = p
	ret.HTMLData = htmlNode.Data
	ret.LuaActiveModeHandlers = make(map[string]NodeJob, 0)
	ret.KeyPressHandlers = make(map[string]NodeJob, 0)
	ret.KeyPressEnterHandlers = make(map[string]NodeJob, 0)

	ret.isShouldCalculateHeight = true
	ret.isShouldCalculateWidth = true
	ret.Position = "relative"

	ret.HTMLAttribute = make(map[string]html.Attribute)

	ret.Cursor = image.Point{-1, -1}
	return ret
}

// SetRelativeCursor 设定 光标
// relativeX relativeY 为相对 node 位置
func (p *Node) SetRelativeCursor(relativeX, relativeY int) (int, int) {
	maxWidth := p.UIBlock.InnerArea.Dx()
	maxHeight := p.UIBlock.InnerArea.Dy()

	if relativeX < 0 {
		relativeX = 0
	} else if relativeX > maxWidth-1 {
		relativeX = maxWidth - 1
	}
	if relativeY < 0 {
		relativeY = 0
	} else if relativeY > maxHeight-1 {
		relativeY = maxHeight - 1
	}

	p.Cursor.X = p.UIBlock.InnerArea.Min.X + p.UIBlock.X + relativeX
	p.Cursor.Y = p.UIBlock.InnerArea.Min.Y + p.UIBlock.Y + relativeY

	utils.UISetCursor(p.Cursor.X, p.Cursor.Y)
	p.UIRender()
	return relativeX, relativeY
}

func (p *Node) ResumeCursor() {
	utils.UISetCursor(p.Cursor.X, p.Cursor.Y)
	p.UIRender()
}

func (p *Node) HideCursor() {
	utils.UISetCursor(-1, -1)
	p.UIRender()
}

func (p *Node) CheckIfDisplay() bool {
	if false == *p.Display {
		return false
	}

	if nil != p.Tab && p.Tab.Index != p.Tab.NodeTabpane.Node.UIBuffer.(*extra.Tabpane).GetActiveIndex() {
		return false
	}

	return true
}
