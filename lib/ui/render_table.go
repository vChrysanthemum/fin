package ui

import (
	"github.com/gizak/termui"
)

func (p *Page) _renderBodyTableOneRow(nodeTr *Node) []*termui.Row {
	var (
		nodeTd         *Node
		nodeTdData     *NodeTableTrTd
		nodeTdChild    *Node
		nodeTdChildren []termui.GridBufferer
		uiCols         []*termui.Row
		ok             bool
		_cols          int

		needCalculateColNodeTdList []*NodeTableTrTd
	)

	// 计算 nodeTd.Cols
	_cols = 0
	needCalculateColNodeTdList = make([]*NodeTableTrTd, 0)
	for nodeTd = nodeTr.FirstChild; nodeTd != nil; nodeTd = nodeTd.NextSibling {
		nodeTdData, ok = nodeTd.Data.(*NodeTableTrTd)
		if false == ok {
			continue
		}

		if 0 == nodeTdData.Cols {
			needCalculateColNodeTdList = append(needCalculateColNodeTdList, nodeTdData)
		} else {
			_cols += nodeTdData.Cols
		}
	}

	_cols = (TABLE_TR_COLS - _cols) / len(needCalculateColNodeTdList)
	if _cols <= 0 {
		_cols = 1
	}

	for _, nodeTdData = range needCalculateColNodeTdList {
		nodeTdData.Cols = _cols
	}

	// 渲染 BodyTableTrTd
	uiCols = make([]*termui.Row, 0)
	for nodeTd = nodeTr.FirstChild; nodeTd != nil; nodeTd = nodeTd.NextSibling {
		nodeTdData, ok = nodeTd.Data.(*NodeTableTrTd)
		if false == ok {
			continue
		}

		nodeTdChildren = make([]termui.GridBufferer, 0)
		for nodeTdChild = nodeTd.FirstChild; nodeTdChild != nil; nodeTdChild = nodeTdChild.NextSibling {

			if _, ok = nodeTdChild.Data.(*NodeTableTr); true == ok {
				nodeTdChildren = append(nodeTdChildren, termui.NewRow(p._renderBodyTableOneRow(nodeTdChild)...))
			} else {
				nodeTdChildren = append(nodeTdChildren, nodeTdChild.uiBuffer.(termui.GridBufferer))
			}
		}
		if len(nodeTdChildren) > 0 {
			uiCols = append(uiCols, termui.NewCol(nodeTdData.Cols, nodeTdData.Offset, nodeTdChildren...))
		}
	}

	return uiCols
}

func (p *Page) renderBodyTable(node *Node) (isFallthrough bool) {
	isFallthrough = false

	var (
		nodeTr *Node
		uiCols []*termui.Row
		uiRows []*termui.Row
	)

	nodeTableData := node.Data.(*NodeTable)

	uiRows = make([]*termui.Row, 0)
	for nodeTr = node.FirstChild; nodeTr != nil; nodeTr = nodeTr.NextSibling {
		uiCols = p._renderBodyTableOneRow(nodeTr)

		uiRows = append(uiRows, termui.NewRow(uiCols...))
	}

	nodeTableData.Body.BgColor = termui.ThemeAttr("bg")
	nodeTableData.Body.Width = termui.TermWidth()
	nodeTableData.Body.AddRows(uiRows...)
	nodeTableData.Body.Align()

	p.bufferersAppend(node, nodeTableData.Body)

	return
}
