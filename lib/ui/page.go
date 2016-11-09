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
	Title string

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

func (p *Page) RemoveNode(node *Node) {
	if p.ActiveNode == node {
		node.QuitActiveMode()
	}

	if p.FirstChildNode == node {
		p.FirstChildNode = node.NextSibling
	}

	if nil != p.FocusNode && p.FocusNode.Value.(*Node) == node {
		p.FocusNode = p.FocusNode.Next()
	}

	for k, v := range p.Bufferers {
		if v == node.uiBuffer {
			p.Bufferers = append(p.Bufferers[:k], p.Bufferers[k+1:]...)
			break
		}
	}

	for nodeElement := p.WorkingNodes.Front(); nodeElement != nil; nodeElement = nodeElement.Next() {
		if nodeElement.Value.(*Node) == node {
			p.WorkingNodes.Remove(nodeElement)
			break
		}
	}

	if nil != node.Parent {
		node.Parent.ChildrenCount -= 1
		if node.Parent.FirstChild == node {
			node.Parent.FirstChild = node.NextSibling
		}
		if node.Parent.LastChild == node {
			node.Parent.LastChild = node.PrevSibling
		}
	}

	if nil != node.NextSibling {
		node.NextSibling.PrevSibling = node.PrevSibling
	}

	if nil != node.PrevSibling {
		node.PrevSibling.NextSibling = node.NextSibling
	}
}

func (p *Page) Serve() {
	defer termui.Close()

	termui.Render(p.Bufferers...)

	p.registerHandles()
	go p.script.Run()

	termui.Loop()
}
