package ui

import (
	"container/list"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type Page struct {
	// 当该 page 为 Modal 类型时，则存在该值
	// MainPage 用于表示 Modal 所属 Page
	MainPage *Page

	Title string

	IDToNodeMap map[string]*Node

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

	HookersAfterFirstUIRender []Hooker
}

func newPage() *Page {
	ret := new(Page)

	ret.IDToNodeMap = make(map[string]*Node, 0)

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
	p.UIRender()
	go p.Script.Run()
	for k, _ := range p.HookersAfterFirstUIRender {
		p.HookersAfterFirstUIRender[k].Do(p.HookersAfterFirstUIRender[k].Arg)
	}
	MainLoop()
}
