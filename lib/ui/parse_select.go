package ui

import (
	rw "github.com/mattn/go-runewidth"
	"golang.org/x/net/html"
)

func (p *Page) parseBodySelect(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeSelect := ret.InitNodeSelect()

	ret.Data = nodeSelect

	return
}

func (p *Page) parseBodySelectOption(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = nil
	isFallthrough = false

	var (
		nodeSelect *NodeSelect
		nodeOption NodeSelectOption
		ok         bool
		attr       html.Attribute
		width      int
	)

	nodeSelect, ok = parentNode.Data.(*NodeSelect)
	if false == ok {
		return
	}

	if nil == htmlNode.FirstChild {
		return
	}

	nodeOption.Data = htmlNode.FirstChild.Data

	for _, attr = range htmlNode.Attr {
		if "value" == attr.Key {
			nodeOption.Value = attr.Val
		}
	}

	nodeSelect.Children = append(nodeSelect.Children, nodeOption)

	width = rw.StringWidth(nodeOption.Data)
	if width > nodeSelect.ChildrenMaxStringWidth {
		nodeSelect.ChildrenMaxStringWidth = width
	}

	return
}
