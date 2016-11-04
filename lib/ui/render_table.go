package ui

import (
	"github.com/gizak/termui"
)

func (p *Page) renderBodyTable(node *Node) (isFallthrough bool) {
	isFallthrough = false

	var (
		nodeTr *Node
		nodeTd *NodeTableTrTd
		uiCols []*termui.Row
		uiRows []*termui.Row
		ok     bool
		_node  *Node
		_cols  int

		needCalculateColsNodeTdList []*NodeTableTrTd
	)

	uiRows = make([]*termui.Row, 0)
	for nodeTr = node.FirstChild; nodeTr != nil; nodeTr = nodeTr.NextSibling {
		uiCols = make([]*termui.Row, 0)
		_cols = 0
		needCalculateColsNodeTdList = make([]*NodeTableTrTd, 0)

		// 计算 nodeTd.Cols
		for _node = nodeTr.FirstChild; _node != nil; _node = _node.NextSibling {
			nodeTd, ok = _node.Data.(*NodeTableTrTd)
			if false == ok {
				continue
			}

			if 0 == nodeTd.Cols {
				needCalculateColsNodeTdList = append(needCalculateColsNodeTdList, nodeTd)
			} else {
				_cols += nodeTd.Cols
			}
		}

		_cols = (TABLE_TR_COLS - _cols) / len(needCalculateColsNodeTdList)
		if _cols <= 0 {
			_cols = 1
		}

		for _, nodeTd = range needCalculateColsNodeTdList {
			nodeTd.Cols = _cols
		}

		for _node = nodeTr.FirstChild; _node != nil; _node = _node.NextSibling {
			nodeTd, ok = _node.Data.(*NodeTableTrTd)
			if false == ok {
				continue
			}

			if nil != _node.FirstChild {
				uiCols = append(uiCols, termui.NewCol(
					nodeTd.Cols,
					nodeTd.Offset,
					_node.FirstChild.uiBuffer.(termui.GridBufferer)))
			}
		}

		uiRows = append(uiRows, termui.NewRow(uiCols...))
	}

	termui.Body.AddRows(uiRows...)

	return
}
