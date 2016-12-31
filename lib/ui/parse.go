package ui

import (
	"strings"

	"golang.org/x/net/html"
)

// 解析 Node 的函数
// isFallthrough 解析完该 Node ，是否继续解析 ChildNodes
// isFallthrough 例: parseBodyTable 用到该参数
type ParseExecFunc func(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool)

type ParseAgent struct {
	path  []string
	parse ParseExecFunc
}

func (p *Page) prepareParse() {
	p.parseAgentMap = []*ParseAgent{
		&ParseAgent{[]string{"html"}, nil},
		&ParseAgent{[]string{"html", "script"}, p.parseHtmlScript},
		&ParseAgent{[]string{"head"}, nil},
		&ParseAgent{[]string{"head", "title"}, p.parseHeadTitle},
		&ParseAgent{[]string{"body"}, p.parseBody},
		&ParseAgent{[]string{"body", "a"}, nil},
		&ParseAgent{[]string{"body", "div"}, p.parseBodyDiv},
		&ParseAgent{[]string{"body", "table"}, p.parseBodyTable},
		&ParseAgent{[]string{"body", "table", "tr"}, p.parseBodyTableTr},
		&ParseAgent{[]string{"body", "table", "tr", "td"}, p.parseBodyTableTrTd},
		&ParseAgent{[]string{"body", "select"}, p.parseBodySelect},
		&ParseAgent{[]string{"body", "select", "option"}, p.parseBodySelectOption},
		&ParseAgent{[]string{"body", "editor"}, p.parseBodyEditor},
		&ParseAgent{[]string{"body", "par"}, p.parseBodyPar},
		&ParseAgent{[]string{"body", "inputtext"}, p.parseBodyInputText},
		&ParseAgent{[]string{"body", "canvas"}, p.parseBodyCanvas},
		&ParseAgent{[]string{"body", "terminal"}, p.parseBodyTerminal},
		&ParseAgent{[]string{"body", "gauge"}, p.parseBodyGauge},
		&ParseAgent{[]string{"body", "tabpane"}, p.parseBodyTabpane},
		&ParseAgent{[]string{"body", "tabpane", "tab"}, p.parseBodyTabpaneTab},
		&ParseAgent{[]string{"body", "modal"}, p.parseBodyModal},
	}
}

func (p *Page) pushParsingNodesStack(node *Node) {
	if nil == p.FirstChildNode {
		p.FirstChildNode = node
	}

	p.parsingNodesStack.PushBack(node)
}

func (p *Page) popParsingNodesStack() *Node {
	ent := p.parsingNodesStack.Back()
	if nil == ent {
		return nil
	}

	p.parsingNodesStack.Remove(ent)
	return ent.Value.(*Node)
}

func (p *Page) checkIfHtmlNodeMatchParseAgentPath(htmlNode *html.Node, parseAgent *ParseAgent, index int) bool {
	if index < 0 {
		return true
	}

	if nil == htmlNode {
		return false
	}

	if htmlNode.Data == parseAgent.path[index] {
		index--
	}
	return p.checkIfHtmlNodeMatchParseAgentPath(htmlNode.Parent, parseAgent, index)
}

func (p *Page) fetchParseAgentByNode(htmlNode *html.Node) (ret *ParseAgent) {
	var parseAgent *ParseAgent

	ret = nil
	for _, parseAgent = range p.parseAgentMap {
		if parseAgent.path[len(parseAgent.path)-1] != htmlNode.Data {
			continue
		}

		if true == p.checkIfHtmlNodeMatchParseAgentPath(htmlNode, parseAgent, len(parseAgent.path)-1) {
			ret = parseAgent
			break
		}
	}

	return ret
}

func (p *Page) filter(htmlNode *html.Node) {
	var childHtmlNode *html.Node

	for childHtmlNode = htmlNode.FirstChild; childHtmlNode != nil; childHtmlNode = childHtmlNode.NextSibling {
		childHtmlNode.Data = strings.Trim(childHtmlNode.Data, " \r\n\t")
		p.filter(childHtmlNode)
	}
}

func (p *Page) parseHtmlNodeToNode(htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	var (
		parentNode *Node
		parseAgent *ParseAgent
	)

	parseAgent = p.fetchParseAgentByNode(htmlNode)

	if nil == parseAgent || nil == parseAgent.parse {
		return nil, true
	}

	ent := p.parsingNodesStack.Back()
	if nil != ent {
		parentNode = ent.Value.(*Node)
	}

	ret, isFallthrough = parseAgent.parse(parentNode, htmlNode)

	if nil != ret && nil != ret.Parent && nil != ret.Parent.Tab {
		ret.Tab = ret.Parent.Tab
	}

	return
}

func (p *Page) parseNodeAttribute(node *Node, attr []html.Attribute) {
	// 公用的解析
	for _, v := range attr {
		switch v.Key {
		case "id":
			p.IdToNodeMap[v.Val] = node
			node.Id = v.Val
		}
	}

	node.ParseAttribute(attr)

	if nil != node.Parent && "" == node.ColorFg {
		node.ColorFg = node.Parent.ColorFg
	}
}

func (p *Page) parse(htmlNode *html.Node) *Node {
	var (
		childHtmlNode *html.Node
		node          *Node
		isFallthrough bool
	)

	node, isFallthrough = p.parseHtmlNodeToNode(htmlNode)
	if nil != node {
		p.pushParsingNodesStack(node)

		p.parseNodeAttribute(node, htmlNode.Attr)
	}

	if true == isFallthrough {
		for childHtmlNode = htmlNode.FirstChild; childHtmlNode != nil; childHtmlNode = childHtmlNode.NextSibling {
			p.parse(childHtmlNode)
		}
	}

	if nil != node {
		p.popParsingNodesStack()
	}

	return node
}

func Parse(content string) (ret *Page, err error) {
	ret = newPage()

	ret.doc, err = html.Parse(strings.NewReader(content))
	ret.filter(ret.doc)
	ret.parse(ret.doc)

	return ret, err
}
