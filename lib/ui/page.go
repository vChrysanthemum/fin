package ui

import (
	"container/list"
	"log"

	"github.com/gizak/termui"
	"golang.org/x/net/html"
)

func init() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
}

type Page struct {
	Title       string
	IdToNodeMap map[string]*Node

	Bufferers []termui.Bufferer

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

	ret.Bufferers = make([]termui.Bufferer, 0)

	ret.parsingNodesStack = list.New()
	ret.WorkingNodes = list.New()

	ret.prepareScript()
	ret.prepareParse()
	ret.prepareRender()

	return ret
}

func (p *Page) dumpNodesHtmlData(node *Node) {
	log.Println(node.HtmlData)
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		p.dumpNodesHtmlData(childNode)
	}
}

func (p *Page) DumpNodesHtmlData() {
	p.dumpNodesHtmlData(p.FirstChildNode)
}

func (p *Page) Serve() {
	defer termui.Close()

	termui.Render(p.Bufferers...)

	p.registerHandles()
	go p.script.Run()

	termui.Loop()
}
