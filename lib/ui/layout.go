package ui

import "container/list"

// 计算 Node 的布局函数
// isFallthrough 计算完该 Node 布局，是否继续计算 ChildNodes 布局
// isChildNodesAvailWorkNode 该 Node 的 ChildNodes 是否可以成为 WorkNode，与用户进行交互
// isChildNodesAvailWorkNode 例: tabpane 将利用该参数
type LayoutExecFunc func(node *Node, isParentDeclareAvailWorkNode bool) (isFallthrough, isChildNodesAvailWorkNode bool)

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

func (p *Page) checkIfHtmlNodeMatchLayoutAgentPath(node *Node, layoutAgent *LayoutAgent, index int) bool {
	if index < 0 {
		return true
	}

	if nil == node {
		return false
	}

	if node.HtmlData == layoutAgent.path[index] {
		index--
	}
	return p.checkIfHtmlNodeMatchLayoutAgentPath(node.Parent, layoutAgent, index)
}

func (p *Page) fetchLayoutAgentByNode(node *Node) (ret *LayoutAgent) {
	var layoutAgent *LayoutAgent

	ret = nil
	for _, layoutAgent = range p.layoutAgentMap {
		if layoutAgent.path[len(layoutAgent.path)-1] != node.HtmlData {
			continue
		}

		if true == p.checkIfHtmlNodeMatchLayoutAgentPath(node, layoutAgent, len(layoutAgent.path)-1) {
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
// 		isParentDeclareAvailWorkNode	bool 	父节点是否已申明子孙节点可以为 WorkNode
func (p *Page) layout(node *Node, isParentDeclareAvailWorkNode bool) error {
	var (
		layoutAgent    *LayoutAgent
		child          *Node
		isFallthrough  bool
		layoutExecFunc LayoutExecFunc
	)

	layoutAgent = p.fetchLayoutAgentByNode(node)
	if true == *node.Display {
		if nil != layoutAgent {
			layoutExecFunc = layoutAgent.layout
		} else {
			layoutExecFunc = p.normalLayoutNodeBlock
		}

		if true == node.isWorkNode && true == isParentDeclareAvailWorkNode {
			p.pushWorkingNode(node)
		}

		if false == isParentDeclareAvailWorkNode {
			isFallthrough, _ = layoutExecFunc(node, isParentDeclareAvailWorkNode)
		} else {
			isFallthrough, isParentDeclareAvailWorkNode = layoutExecFunc(node, isParentDeclareAvailWorkNode)
		}

		if false == isFallthrough {
			return nil
		}
	}

	for child = node.FirstChild; child != nil; child = child.NextSibling {
		p.layout(child, isParentDeclareAvailWorkNode)
	}

	return nil
}

// 计算 page 中所有元素的布局
func (p *Page) Layout() error {
	p.layoutingX = 0
	p.layoutingY = 0

	p.WorkingNodes = list.New()
	err := p.layout(p.FirstChildNode, true)
	if nil != err {
		return err
	}

	return nil
}
