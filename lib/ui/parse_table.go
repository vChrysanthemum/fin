package ui

import (
	"strconv"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyTable(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTable := ret.InitNodeTable()

	ret.Data = nodeTable

	return
}

func (p *Page) parseBodyTableTr(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTableTr := ret.InitNodeTableTr()

	ret.Data = nodeTableTr

	return
}

func (p *Page) parseBodyTableTrTd(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	nodeTableTrTd := ret.InitNodeTableTrTd()

	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "offset":
			nodeTableTrTd.Offset, _ = strconv.Atoi(attr.Val)
			if nodeTableTrTd.Offset < 0 {
				nodeTableTrTd.Offset = 0
			}
		case "cols":
			nodeTableTrTd.Cols, _ = strconv.Atoi(attr.Val)
			if nodeTableTrTd.Cols <= 0 {
				// nodeTableTrTd.Cols == 0 时，其Cols将在render阶段计算
				nodeTableTrTd.Cols = 0
			}
		}
	}

	ret.isShouldTermuiRenderChild = true
	ret.Data = nodeTableTrTd

	return
}
