package ui

import (
	"container/list"
	"sync"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type Page struct {
	// 当该 page 为 Modal 类型时，则存在该值
	// MainPage 用于表示 Modal 所属 Page
	MainPage *Page

	Title string

	IdToNodeMap map[string]*Node

	Bufferers []termui.Bufferer

	parseAgentMap  []*ParseAgent
	renderAgentMap []*RenderAgent
	layoutAgentMap []*LayoutAgent
	Script         *Script
	Modals         map[string]*NodeModal
	CurrentModal   *NodeModal

	doc                     *html.Node
	parsingNodesStack       *list.List
	FirstChildNode          *Node
	FocusNode               *list.Element
	ActiveNode              *Node
	ActiveNodeAfterReRender *Node
	// 能接受用户访问的 Nodes
	WorkingNodes *list.List

	layoutingX int
	layoutingY int

	recoverVal interface{}

	KeyPressHandleLocker sync.RWMutex
}

func newPage() *Page {
	ret := new(Page)

	ret.IdToNodeMap = make(map[string]*Node, 0)

	ret.parsingNodesStack = list.New()
	ret.WorkingNodes = list.New()

	ret.prepareScript()
	ret.prepareModals()
	ret.prepareParse()
	ret.prepareRender()
	ret.prepareLayout()

	return ret
}

func (p *Page) Serve() {
	p.uiRender()

	go func() {
		defer func() {
			if r := recover(); nil != r {
				termui.StopLoop()
				p.recoverVal = r
			}
		}()
		p.Script.Run()
	}()
}
