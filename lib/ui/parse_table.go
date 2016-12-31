package ui

import (
	"strconv"

	"github.com/gizak/termui"

	"golang.org/x/net/html"
)

func (p *Page) parseBodyTable(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	ret.InitNodeTable()

	return
}

func (p *Page) parseBodyTableTr(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	ret.InitNodeTableTr()

	return
}

func (p *Page) parseBodyTableTrTd(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = p.newNode(htmlNode)
	parentNode.addChild(ret)
	isFallthrough = true

	ret.InitNodeTableTrTd()
	nodeTableTrTd := ret.Data.(*NodeTableTrTd)

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

	return
}

func (p *NodeTable) NodeDataParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	uiBuffer := p.Node.uiBuffer.(*termui.Grid)

	for _, v := range attr {
		switch v.Key {
		case "top":
			tmp, err := strconv.Atoi(v.Val)
			if nil == err {
				uiBuffer.Y = tmp
				p.Node.isSettedPositionY = true
			}
		case "left":
			tmp, err := strconv.Atoi(v.Val)
			if nil == err {
				uiBuffer.X = tmp
				p.Node.isSettedPositionX = true
			}

		}
	}

	return

}
