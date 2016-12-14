package ui

import "github.com/gizak/termui"

func (p *Page) _renderBodyTableOneRow(nodeTr *Node) []*termui.Row {
	var (
		nodeTd         *Node
		nodeDataTd     *NodeTableTrTd
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
		nodeDataTd, ok = nodeTd.Data.(*NodeTableTrTd)
		if false == ok {
			continue
		}

		if 0 == nodeDataTd.Cols {
			needCalculateColNodeTdList = append(needCalculateColNodeTdList, nodeDataTd)
		} else {
			_cols += nodeDataTd.Cols
		}
	}

	if len(needCalculateColNodeTdList) > 0 {
		_cols = (TABLE_TR_COLS - _cols) / len(needCalculateColNodeTdList)
		if _cols <= 0 {
			_cols = 1
		}

		for _, nodeDataTd = range needCalculateColNodeTdList {
			nodeDataTd.Cols = _cols
		}
	}

	// 渲染 BodyTableTrTd
	uiCols = make([]*termui.Row, 0)
	for nodeTd = nodeTr.FirstChild; nodeTd != nil; nodeTd = nodeTd.NextSibling {
		nodeDataTd, ok = nodeTd.Data.(*NodeTableTrTd)
		if false == ok {
			continue
		}

		nodeTdChildren = make([]termui.GridBufferer, 0)
		for nodeTdChild = nodeTd.FirstChild; nodeTdChild != nil; nodeTdChild = nodeTdChild.NextSibling {
			if true == nodeTdChild.isShouldHide {
				continue
			}

			if _, ok = nodeTdChild.Data.(*NodeTableTr); true == ok {
				nodeTdChildren = append(nodeTdChildren, termui.NewRow(p._renderBodyTableOneRow(nodeTdChild)...))
			} else {
				nodeTdChildren = append(nodeTdChildren, nodeTdChild.uiBuffer.(termui.GridBufferer))
			}
		}
		if len(nodeTdChildren) > 0 {
			uiCols = append(uiCols, termui.NewCol(nodeDataTd.Cols, nodeDataTd.Offset, nodeTdChildren...))
		}
	}

	return uiCols
}

func (p *Page) renderBodyTable(node *Node) {
	var (
		nodeTr *Node
		uiCols []*termui.Row
		uiRows []*termui.Row
	)

	nodeDataTable := node.Data.(*NodeTable)

	nodeDataTable.Body.Rows = []*termui.Row{}

	uiRows = make([]*termui.Row, 0)
	for nodeTr = node.FirstChild; nodeTr != nil; nodeTr = nodeTr.NextSibling {
		uiCols = p._renderBodyTableOneRow(nodeTr)

		uiRows = append(uiRows, termui.NewRow(uiCols...))
	}

	nodeDataTable.Body.BgColor = termui.ThemeAttr("bg")
	nodeDataTable.Body.Width = termui.TermWidth()
	nodeDataTable.Body.AddRows(uiRows...)

	p.BufferersAppend(node, nodeDataTable.Body)

	return
}
