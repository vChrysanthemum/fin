package editor

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

type CursorLocation struct {
	IsDisplay     bool
	OffXCellIndex int
	Editor        *Editor
}

func NewCursorLocation(editor *Editor) *CursorLocation {
	ret := &CursorLocation{
		IsDisplay:     false,
		OffXCellIndex: 0,
		Editor:        editor,
	}
	return ret
}

func (p *CursorLocation) ResetLocation() {
	uiutils.UISetCursor(p.Editor.Block.InnerArea.Min.X, p.Editor.Block.InnerArea.Min.Y)
}

func (p *CursorLocation) ResumeCursor() {
	//uiutils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *CursorLocation) MoveCursorNRuneTop(n int) {
	if n <= 0 {
		return
	}

	if nil == p.Editor.CurrentLine.Prev || len(p.Editor.Lines) == 0 {
		return
	}

	var index, i int

	if n >= len(p.Editor.Lines) {
		index = 0
		goto MOVE_END
	}

	for i = 0; i <= n; i++ {
		if p.Editor.Lines[i] == p.Editor.CurrentLine {
			index = 0
			goto MOVE_END
		}
	}

	for i = n + 1; i < len(p.Editor.Lines); i++ {
		if p.Editor.Lines[i] == p.Editor.CurrentLine {
			index = i - n
			goto MOVE_END
		}
	}

MOVE_END:
	p.Editor.CurrentLine = p.Editor.Lines[index]
	if index < p.Editor.DisplayLinesTopIndex {
		p.Editor.DisplayLinesTopIndex = index
		p.Editor.RefreshContent()
	}

	if 0 == len(p.Editor.CurrentLine.Cells) {
		uiutils.UISetCursor(p.Editor.CurrentLine.ContentStartX, p.Editor.CurrentLine.ContentStartY)
		uiutils.UIRender(p.Editor)

	} else {
		if p.OffXCellIndex >= len(p.Editor.CurrentLine.Cells) {
			p.OffXCellIndex = len(p.Editor.CurrentLine.Cells) - 1
		}

		cell := p.Editor.CurrentLine.Cells[p.OffXCellIndex]
		uiutils.UISetCursor(cell.X, cell.Y)
		uiutils.UIRender(p.Editor)
	}
}

func (p *CursorLocation) MoveCursorNRuneBottom(n int) {
	if n <= 0 {
		return
	}

	if nil == p.Editor.CurrentLine.Next || len(p.Editor.Lines) == 0 {
		return
	}

	var i, index, last int

	if n >= len(p.Editor.Lines) {
		index = len(p.Editor.Lines) - 1
		goto MOVE_END
	}

	last = len(p.Editor.Lines) - n
	if last <= 0 {
		index = len(p.Editor.Lines) - 1
		goto MOVE_END
	}

	for i = 0; i < len(p.Editor.Lines); i++ {
		if i >= last {
			index = len(p.Editor.Lines) - 1
			goto MOVE_END
		}

		if p.Editor.Lines[i] == p.Editor.CurrentLine {
			index = i + n
			goto MOVE_END
		}
	}

MOVE_END:
	p.Editor.CurrentLine = p.Editor.Lines[index]
	if index > p.Editor.DisplayLinesBottomIndex {
		p.Editor.DisplayLinesTopIndex += (index - p.Editor.DisplayLinesBottomIndex)
		p.Editor.RefreshContent()
	}

	if 0 == len(p.Editor.CurrentLine.Cells) {
		uiutils.UISetCursor(p.Editor.CurrentLine.ContentStartX, p.Editor.CurrentLine.ContentStartY)
		uiutils.UIRender(p.Editor)

	} else {
		if p.OffXCellIndex >= len(p.Editor.CurrentLine.Cells) {
			p.OffXCellIndex = len(p.Editor.CurrentLine.Cells) - 1
		}

		cell := p.Editor.CurrentLine.Cells[p.OffXCellIndex]
		uiutils.UISetCursor(cell.X, cell.Y)
		uiutils.UIRender(p.Editor)
	}
}

func (p *CursorLocation) MoveCursorNRuneLeft(n int) {
	if n <= 0 {
		return
	}

	if len(p.Editor.CurrentLine.Cells) == 0 {
		p.OffXCellIndex = 0
		return
	}

	p.OffXCellIndex -= n
	if p.OffXCellIndex < 0 {
		p.OffXCellIndex = 0
	}

	cell := p.Editor.CurrentLine.Cells[p.OffXCellIndex]
	uiutils.UISetCursor(cell.X, cell.Y)
	uiutils.UIRender(p.Editor)
}

func (p *CursorLocation) MoveCursorNRuneRight(n int) {
	if n <= 0 {
		return
	}

	if len(p.Editor.CurrentLine.Cells) == 0 {
		p.OffXCellIndex = 0
		return
	}

	p.OffXCellIndex += n
	if p.OffXCellIndex >= len(p.Editor.CurrentLine.Cells) {
		p.OffXCellIndex = len(p.Editor.CurrentLine.Cells) - 1
	}

	cell := p.Editor.CurrentLine.Cells[p.OffXCellIndex]
	uiutils.UISetCursor(cell.X, cell.Y)
	uiutils.UIRender(p.Editor)
}

func (p *CursorLocation) RefreshCursorByLine(line *Line) {
	if nil == line {
		uiutils.UISetCursor(p.Editor.Block.InnerArea.Min.X, p.Editor.Block.InnerArea.Min.Y)
		return
	}

	if 0 == len(line.Cells) {
		uiutils.UISetCursor(line.ContentStartX, line.ContentStartY)
		return
	}

	if p.OffXCellIndex >= len(line.Cells) {
		p.OffXCellIndex = len(line.Cells)
	}

	var x, y int
	var cell termui.Cell
	if p.OffXCellIndex == len(line.Cells) {
		cell = line.Cells[p.OffXCellIndex-1]
		width := cell.Width()
		if cell.X+width >= p.Editor.Block.InnerArea.Max.X {
			x, y = line.ContentStartX, cell.Y+1
		} else {
			x, y = cell.X+width, cell.Y
		}

		if y >= p.Editor.Block.InnerArea.Max.Y {
			y = p.Editor.Block.InnerArea.Max.Y - 1
			p.Editor.DisplayLinesTopIndex += 1
			p.Editor.RefreshContent()
		}

	} else {
		cell = line.Cells[p.OffXCellIndex]
		x, y = cell.X, cell.Y
	}

	uiutils.UISetCursor(x, y)
}
