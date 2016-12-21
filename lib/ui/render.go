package ui

import (
	"container/list"

	uiutils "fin/ui/utils"

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
	if true == *node.Display && nil != renderAgent {
		renderAgent.render(node)
	}

	return nil
}

// 将 page 上的内容渲染到屏幕上
// 更新 FocusNode
// 更新 ActiveNode
// 更新 FocusNode / ActiveNode / WorkingNodes内元素之间方向关系
func (p *Page) uiRender() error {
	GCurrentRenderPage = p
	if 0 == len(p.Bufferers) {
		return nil
	}

	// 更新 FocusNode
	if nil != p.FocusNode {
		if nodeDataUnFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer); true == ok {
			nodeDataUnFocusModer.NodeDataUnFocusMode()
		}
	}

	// 更新 ActiveNode
	if nil != p.ActiveNode {
		if nodeDataUnActiveModer, ok := p.ActiveNode.Data.(NodeDataUnActiveModer); true == ok {
			nodeDataUnActiveModer.NodeDataUnActiveMode()
		}
	}

	// 更新 FocusNode / ActiveNode / WorkingNodes内元素之间方向关系
	if nil != p.ActiveNodeAfterReRender {
		p.SetActiveNode(p.ActiveNodeAfterReRender)
		p.FocusNode = nil
	} else if nil != p.FocusNode {
		p.SetActiveNode(p.FocusNode.Value.(*Node))
	}

	var (
		e, e2       *list.Element
		node, node2 *Node
	)
	if p.WorkingNodes.Len() > 0 {
		for e = p.WorkingNodes.Front(); e != nil; e = e.Next() {
			node = e.Value.(*Node)
			node.FocusThisNode = e
			node.FocusTopNode = nil
			node.FocusBottomNode = nil
			node.FocusLeftNode = nil
			node.FocusRightNode = nil
		}

		for e = p.WorkingNodes.Front(); e != nil; e = e.Next() {
			node = e.Value.(*Node)

			if false == *node.Display {
				continue
			}

			// 更新 WorkingNodes内元素之间上下方向关系
			for e2 = p.WorkingNodes.Front(); e2 != nil; e2 = e2.Next() {
				node2 = e2.Value.(*Node)

				if false == *node2.Display ||
					node == node2 ||
					nil != node.FocusBottomNode ||
					nil != node2.FocusTopNode {
					continue
				}

				if ((node.UIBlock.InnerArea.Min.X <= node2.UIBlock.InnerArea.Min.X &&
					node2.UIBlock.InnerArea.Min.X <= node.UIBlock.InnerArea.Max.X) ||

					(node.UIBlock.InnerArea.Min.X <= node2.UIBlock.InnerArea.Max.X &&
						node2.UIBlock.InnerArea.Max.X <= node.UIBlock.InnerArea.Max.X) ||

					(node2.UIBlock.InnerArea.Min.X <= node.UIBlock.InnerArea.Min.X &&
						node2.UIBlock.InnerArea.Max.X >= node.UIBlock.InnerArea.Max.X)) &&

					(node2.UIBlock.Y > node.UIBlock.Y) {

					node.FocusBottomNode = e2
					node2.FocusTopNode = e
					break
				}
			}

			// 更新 WorkingNodes内元素之间左右方向关系
			for e2 = p.WorkingNodes.Front(); e2 != nil; e2 = e2.Next() {
				node2 = e2.Value.(*Node)

				if false == *node2.Display ||
					node == node2 ||
					nil != node.FocusLeftNode ||
					nil != node2.FocusRightNode {
					continue
				}

				if ((node.UIBlock.InnerArea.Min.Y <= node2.UIBlock.InnerArea.Min.Y &&
					node2.UIBlock.InnerArea.Min.Y <= node.UIBlock.InnerArea.Max.Y) ||

					(node.UIBlock.InnerArea.Min.Y <= node2.UIBlock.InnerArea.Max.Y &&
						node2.UIBlock.InnerArea.Max.Y <= node.UIBlock.InnerArea.Max.Y) ||

					(node2.UIBlock.InnerArea.Min.Y <= node.UIBlock.InnerArea.Min.Y &&
						node2.UIBlock.InnerArea.Max.Y >= node.UIBlock.InnerArea.Max.Y)) &&

					(node2.UIBlock.X > node.UIBlock.X) {

					node.FocusRightNode = e2
					node2.FocusLeftNode = e
					break
				}
			}

		}
	}

	uiutils.UIRender(p.Bufferers...)

	return nil
}

func (p *Page) BufferersAppend(node *Node, buffer termui.Bufferer) {
	if nil != node && true == node.Parent.isShouldTermuiRenderChild {
	} else {
		p.Bufferers = append(p.Bufferers, buffer)
	}
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

// 重新渲染 page 并刷新内容到屏幕
func (p *Page) ReRender() {
	uiClear(0, -1)
	p.Render()
	p.uiRender()
}

// 清空 page 中所有元素，但不清空屏幕
func (p *Page) Clear() {
	p.Bufferers = make([]termui.Bufferer, 0)
	p.FocusNode = nil
	p.WorkingNodes = list.New()
	p.ActiveNode = nil
}
