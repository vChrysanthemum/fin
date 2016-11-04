package ui

import "github.com/gizak/termui"

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
	}
}

func (p *Page) bufferersAppend(node *Node, buffer termui.Bufferer) {
	if nil != node && true == node.Parent.isShouldTermuiRenderChild {
	} else {
		p.bufferers = append(p.bufferers, buffer)
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

func (p *Page) render(node *Node) {
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
}

func (p *Page) Render() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	p.render(p.FirstChildNode)
}
