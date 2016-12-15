package ui

import (
	"container/list"

	"github.com/gizak/termui"
)

type RenderExecFunc func(node *Node)

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
		&RenderAgent{[]string{"body", "tabpane"}, p.renderBodyTabpane},
		&RenderAgent{[]string{"body", "modal"}, p.renderBodyModal},
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
	if true == *node.Display {
		if nil != renderAgent {
			renderAgent.render(node)
		}
	}

	return nil
}

// 渲染 page 中所有元素，但不输出到屏幕
func (p *Page) Render() error {
	p.Clear()

	err := p.render(p.FirstChildNode)
	if nil != err {
		return err
	}

	err = p.Layout()
	if nil != err {
		return err
	}

	return nil
}

// 清空 page 中所有元素，但不清空屏幕
func (p *Page) Clear() {
	p.Bufferers = make([]termui.Bufferer, 0)
	p.FocusNode = nil
	p.WorkingNodes = list.New()
	p.ActiveNode = nil
}

func (p *Page) BufferersAppend(node *Node, buffer termui.Bufferer) {
	if nil != node && true == node.Parent.isShouldTermuiRenderChild {
	} else {
		p.Bufferers = append(p.Bufferers, buffer)
	}
}
