package ui

import (
	"container/list"

	"github.com/gizak/termui"
)

type RenderExecFunc func(node *Node) (isFallthrough bool)

type RenderAgent struct {
	path   []string
	render RenderExecFunc
}

func (p *Page) prepareRender() {
	p.renderAgentMap = []*RenderAgent{
		&RenderAgent{[]string{"body", "div"}, p.renderBodyDiv},
		&RenderAgent{[]string{"body", "select"}, p.renderBodySelect},
		&RenderAgent{[]string{"body", "editor"}, p.renderBodyEditor},
		&RenderAgent{[]string{"body", "par"}, p.renderBodyPar},
		&RenderAgent{[]string{"body", "table"}, p.renderBodyTable},
		&RenderAgent{[]string{"body", "inputtext"}, p.renderBodyInputText},
		&RenderAgent{[]string{"body", "canvas"}, p.renderBodyCanvas},
		&RenderAgent{[]string{"body", "terminal"}, p.renderBodyTerminal},
		&RenderAgent{[]string{"body", "gauge"}, p.renderBodyGauge},
	}
}

func (p *Page) BufferersAppend(node *Node, buffer termui.Bufferer) {
	if nil != node && true == node.Parent.isShouldTermuiRenderChild {
	} else {
		p.Bufferers = append(p.Bufferers, buffer)
	}
}

func (p *Page) checkIfHtmlNodeMatchRenderAgentPath(node *Node, renderAgent *RenderAgent, index int) bool {
	if index < 0 {
		return true
	}

	if nil == node {
		return false
	}

	if node.HtmlData == renderAgent.path[index] {
		index--
	}
	return p.checkIfHtmlNodeMatchRenderAgentPath(node.Parent, renderAgent, index)
}

func (p *Page) fetchRenderAgentByNode(node *Node) (ret *RenderAgent) {
	var renderAgent *RenderAgent

	ret = nil
	for _, renderAgent = range p.renderAgentMap {
		if renderAgent.path[len(renderAgent.path)-1] != node.HtmlData {
			continue
		}

		if true == p.checkIfHtmlNodeMatchRenderAgentPath(node, renderAgent, len(renderAgent.path)-1) {
			ret = renderAgent
			break
		}
	}

	return ret
}

func (p *Page) render(node *Node) error {
	var (
		renderAgent *RenderAgent
		child       *Node
	)

	for child = node.FirstChild; child != nil; child = child.NextSibling {
		p.render(child)
	}

	renderAgent = p.fetchRenderAgentByNode(node)
	if nil != renderAgent {
		renderAgent.render(node)
	}

	return nil
}

func (p *Page) normalRenderNodeBlock(node *Node) {
	if nil == node.UIBlock {
		return
	}

	if true == node.isShouldCalculateWidth {
		node.UIBlock.Width = termui.TermWidth()
	}

	node.UIBlock.X = p.renderingX
	node.UIBlock.Y = p.renderingY
}

func (p *Page) Clear() {
	p.Bufferers = make([]termui.Bufferer, 0)
	p.FocusNode = nil
	p.WorkingNodes = list.New()
	p.ActiveNode = nil

	p.renderingX = 0
	p.renderingY = 0
}

func (p *Page) Render() error {
	p.Clear()

	err := p.render(p.FirstChildNode)
	if nil != err {
		return err
	}

	return nil
}
