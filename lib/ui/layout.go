package ui

import "container/list"

// LayoutExecFunc 计算 Node 的布局函数
// isFallthrough 计算完该 Node 布局，是否继续计算 ChildNodes 布局
type LayoutExecFunc func(node *Node) (isFallthrough bool)

type LayoutAgent struct {
	path   []string
	layout LayoutExecFunc
}

func (p *Page) prepareLayout() {
	p.layoutAgentMap = []*LayoutAgent{
		&LayoutAgent{[]string{"body", "table"}, p.layoutBodyTable},
		&LayoutAgent{[]string{"body", "tabpane"}, p.layoutBodyTabpane},
		&LayoutAgent{[]string{"body", "tabpane", "tab"}, p.layoutBodyTabpaneTab},
	}
}

func (p *Page) checkIfHTMLNodeMatchLayoutAgentPath(node *Node, layoutAgent *LayoutAgent, index int) bool {
	if index < 0 {
		return true
	}

	if nil == node {
		return false
	}

	if node.HTMLData == layoutAgent.path[index] {
		index--
	}
	return p.checkIfHTMLNodeMatchLayoutAgentPath(node.Parent, layoutAgent, index)
}

func (p *Page) fetchLayoutAgentByNode(node *Node) (ret *LayoutAgent) {
	var layoutAgent *LayoutAgent

	ret = nil
	for _, layoutAgent = range p.layoutAgentMap {
		if layoutAgent.path[len(layoutAgent.path)-1] != node.HTMLData {
			continue
		}

		if true == p.checkIfHTMLNodeMatchLayoutAgentPath(node, layoutAgent, len(layoutAgent.path)-1) {
			ret = layoutAgent
			break
		}
	}

	return ret
}

// 计算布局
//
// param:
// 		node 							*Node 	需要计算布局的 Node
func (p *Page) layout(node *Node) error {
	var (
		layoutAgent    *LayoutAgent
		child          *Node
		isFallthrough  bool
		layoutExecFunc LayoutExecFunc
	)

	layoutAgent = p.fetchLayoutAgentByNode(node)
	if true == node.CheckIfDisplay() {
		if nil != layoutAgent {
			layoutExecFunc = layoutAgent.layout
		} else {
			layoutExecFunc = p.normalLayoutNodeBlock
		}

		if true == node.isWorkNode {
			p.pushWorkingNode(node)
		}

		isFallthrough = layoutExecFunc(node)

		if false == isFallthrough {
			return nil
		}
	}

	for child = node.FirstChild; child != nil; child = child.NextSibling {
		p.layout(child)
	}

	return nil
}

// Layout 计算 page 中所有元素的布局
func (p *Page) Layout() error {
	p.layoutingX = 0
	p.layoutingY = 0

	p.WorkingNodes = list.New()
	err := p.layout(p.FirstChildNode)
	if nil != err {
		return err
	}

	return nil
}
