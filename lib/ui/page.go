package ui

import (
	"container/list"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	IdToNodeMap map[string]*Node

	bufferers []termui.Bufferer

	script         *Script
	parseAgentMap  []*ParseAgent
	renderAgentMap []*RenderAgent

	doc               *html.Node
	parsingNodesStack *list.List
	FirstChildNode    *Node
	WorkingNodes      *list.List
	FocusNode         *list.Element
	ActiveNode        *Node

	renderingX int
	renderingY int
}

func newPage() *Page {
	ret := new(Page)

	ret.IdToNodeMap = make(map[string]*Node, 0)

	ret.bufferers = make([]termui.Bufferer, 0)

	ret.parsingNodesStack = list.New()
	ret.WorkingNodes = list.New()

	ret.prepareScript()
	ret.prepareParse()
	ret.prepareRender()

	return ret
}

func (p *Page) Serve() {
	defer termui.Close()

	termui.Render(p.bufferers...)

	p.registerHandles()

	termui.Loop()
}
