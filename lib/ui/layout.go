package ui

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

func (p *Page) layout(node *Node) error {
	var (
		layoutAgent   *LayoutAgent
		child         *Node
		isFallthrough bool
	)

	layoutAgent = p.fetchLayoutAgentByNode(node)
	if true == *node.Display {
		if nil != layoutAgent {
			isFallthrough = layoutAgent.layout(node)
			if false == isFallthrough {
				return nil
			}
		} else {
			p.normalLayoutNodeBlock(node)
		}
	}

	for child = node.FirstChild; child != nil; child = child.NextSibling {
		p.layout(child)
	}

	return nil
}

// 计算 page 中所有元素的布局
func (p *Page) Layout() error {
	p.layoutingX = 0
	p.layoutingY = 0

	err := p.layout(p.FirstChildNode)
	if nil != err {
		return err
	}

	return nil
}
