package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorViewCursor struct {
	EditorView *EditorView

	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int

	LineIndex        int
	CellOffXVertical int
	*EditorCursor
}

func NewEditorViewCursor(editorView *EditorView) *EditorViewCursor {
	ret := &EditorViewCursor{
		EditorView:       editorView,
		CellOffXVertical: 0,
		EditorCursor:     NewEditorCursor(),
	}
	return ret
}

func (p *EditorView) EditModeMoveCursorUp(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - runenum
	if index < 0 {
		index = 0
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index < cursor.DisplayLinesTopIndex {
		cursor.DisplayLinesTopIndex = index
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.Width(), cell.Y+cell.Width())
		} else {
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)
		}
	}

	if cursor.CellOffXVertical > cursor.CellOffX {
		if cursor.CellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells)
			}
		} else {
			cursor.CellOffX = cursor.CellOffXVertical
		}
	}
}

func (p *EditorView) NormalModeMoveCursorUp(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - runenum
	if index < 0 {
		index = 0
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index < cursor.DisplayLinesTopIndex {
		cursor.DisplayLinesTopIndex = index
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells) - 1
		}

		cell := line.Cells[cursor.CellOffX]
		cursor.UISetCursor(cell.X, cell.Y)
	}

	if cursor.CellOffXVertical > cursor.CellOffX {
		if cursor.CellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells) - 1
			}
		} else {
			cursor.CellOffX = cursor.CellOffXVertical
		}
	}
}

func (p *EditorView) EditModeMoveCursorDown(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + runenum
	if index >= len(p.Lines) {
		index = len(p.Lines) - 1
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index > cursor.DisplayLinesBottomIndex {
		cursor.DisplayLinesTopIndex += (index - cursor.DisplayLinesBottomIndex)
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.Width(), cell.Y+cell.Width())
		} else {
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)
		}
	}

	if cursor.CellOffXVertical > cursor.CellOffX {
		if cursor.CellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells)
			}
		} else {
			cursor.CellOffX = cursor.CellOffXVertical
		}
	}
}

func (p *EditorView) NormalModeMoveCursorDown(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + runenum
	if index >= len(p.Lines) {
		index = len(p.Lines) - 1
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index > cursor.DisplayLinesBottomIndex {
		cursor.DisplayLinesTopIndex += (index - cursor.DisplayLinesBottomIndex)
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells) - 1
		}

		cell := line.Cells[cursor.CellOffX]
		cursor.UISetCursor(cell.X, cell.Y)
	}

	if cursor.CellOffXVertical > cursor.CellOffX {
		if cursor.CellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells) - 1
			}
		} else {
			cursor.CellOffX = cursor.CellOffXVertical
		}
	}
}

func (p *EditorView) MoveCursorLeft(cursor *EditorViewCursor, line *EditorLine, runenum int) {
	if runenum <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursor.CellOffX = 0
		return
	}

	cursor.CellOffX -= runenum
	if cursor.CellOffX < 0 {
		cursor.CellOffX = 0
	}

	cell := line.Cells[cursor.CellOffX]
	cursor.UISetCursor(cell.X, cell.Y)

	cursor.CellOffXVertical = cursor.CellOffX
}

func (p *EditorView) MoveCursorRight(cursor *EditorViewCursor, line *EditorLine, runenum int) {
	if runenum <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursor.CellOffX = 0
		return
	}

	cursor.CellOffX += runenum
	if cursor.CellOffX >= len(line.Cells) {
		switch p.Mode {
		case EditorNormalMode:
			cursor.CellOffX = len(line.Cells) - 1
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)

		case EditorEditMode:
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.Width(), cell.Y)
		}

	} else {
		cell := line.Cells[cursor.CellOffX]
		cursor.UISetCursor(cell.X, cell.Y)
	}

	cursor.CellOffXVertical = cursor.CellOffX
}

func (p *EditorViewCursor) RefreshCursorByEditorLine(line *EditorLine) {
	if nil == line {
		return
	}

	if 0 == len(line.Cells) {
		p.CellOffX = 0
		p.UISetCursor(line.ContentStartX, line.ContentStartY)
		return
	}

	if p.CellOffX >= len(line.Cells) {
		p.CellOffX = len(line.Cells)
	}

	var x, y int
	var cell termui.Cell
	if p.CellOffX == len(line.Cells) {
		cell = line.Cells[p.CellOffX-1]
		width := cell.Width()
		if cell.X+width >= p.EditorView.Block.InnerArea.Max.X {
			x, y = line.ContentStartX, cell.Y+1
		} else {
			x, y = cell.X+width, cell.Y
		}

		if y >= p.EditorView.Block.InnerArea.Max.Y {
			y = p.EditorView.Block.InnerArea.Max.Y - 1
			switch p.EditorView.Mode {
			case EditorEditMode:
				p.DisplayLinesTopIndex++
				p.EditorView.isShouldRefreshEditModeBuf = true
			}
		}

	} else {
		cell = line.Cells[p.CellOffX]
		x, y = cell.X, cell.Y
	}

	if 0 == y && 0 == x {
		x, y = line.ContentStartX, line.ContentStartY
	}
	p.UISetCursor(x, y)
}

func (p *EditorViewCursor) Line() *EditorLine {
	if p.LineIndex >= len(p.EditorView.Lines) {
		return nil
	}
	return p.EditorView.Lines[p.LineIndex]
}

func (p *EditorViewCursor) UISetCursor(x, y int) {
	utils.UISetCursor(x, y)
}
