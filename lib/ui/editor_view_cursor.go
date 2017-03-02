package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorViewCursor struct {
	EditorView *EditorView

	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int

	cellOffXVertical int
	LineIndex        int
	*EditorCursor
}

func NewEditorViewCursor(editorView *EditorView) *EditorViewCursor {
	ret := &EditorViewCursor{
		EditorView:       editorView,
		cellOffXVertical: 0,
		EditorCursor:     NewEditorCursor(),
	}
	return ret
}

func (p *EditorView) InputModeMoveCursorUp(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.cellOffXVertical {
		cursor.cellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - runenum
	if index < 0 {
		index = 0
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index < cursor.DisplayLinesTopIndex {
		cursor.DisplayLinesTopIndex = index
		p.isShouldRefreshInputModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.UIWidth, cell.Y)
		} else {
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)
		}
	}

	if cursor.cellOffXVertical > cursor.CellOffX {
		if cursor.cellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells)
			}
		} else {
			cursor.CellOffX = cursor.cellOffXVertical
		}
	}
}

func (p *EditorView) CommandModeMoveCursorUp(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.cellOffXVertical {
		cursor.cellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - runenum
	if index < 0 {
		index = 0
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index < cursor.DisplayLinesTopIndex {
		cursor.DisplayLinesTopIndex = index
		p.isShouldRefreshInputModeBuf = true
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

	if cursor.cellOffXVertical > cursor.CellOffX {
		if cursor.cellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells) - 1
			}
		} else {
			cursor.CellOffX = cursor.cellOffXVertical
		}
	}
}

func (p *EditorView) InputModeMoveCursorDown(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.cellOffXVertical {
		cursor.cellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + runenum
	if index >= len(p.Lines) {
		index = len(p.Lines) - 1
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index > cursor.DisplayLinesBottomIndex {
		cursor.DisplayLinesTopIndex += (index - cursor.DisplayLinesBottomIndex)
		p.isShouldRefreshInputModeBuf = true
	}

	if 0 == len(line.Cells) {
		cursor.UISetCursor(line.ContentStartX, line.ContentStartY)

	} else {
		if cursor.CellOffX >= len(line.Cells) {
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.UIWidth, cell.Y)
		} else {
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)
		}
	}

	if cursor.cellOffXVertical > cursor.CellOffX {
		if cursor.cellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells)
			}
		} else {
			cursor.CellOffX = cursor.cellOffXVertical
		}
	}
}

func (p *EditorView) CommandModeMoveCursorDown(cursor *EditorViewCursor, runenum int) {
	if runenum <= 0 {
		return
	}

	if cursor.CellOffX > cursor.cellOffXVertical {
		cursor.cellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + runenum
	if index >= len(p.Lines) {
		index = len(p.Lines) - 1
	}

	cursor.LineIndex = index
	line := cursor.Line()

	if index > cursor.DisplayLinesBottomIndex {
		cursor.DisplayLinesTopIndex += (index - cursor.DisplayLinesBottomIndex)
		p.isShouldRefreshInputModeBuf = true
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

	if cursor.cellOffXVertical > cursor.CellOffX {
		if cursor.cellOffXVertical >= len(line.Cells) {
			if 0 == len(line.Cells) {
				cursor.CellOffX = 0
			} else {
				cursor.CellOffX = len(line.Cells) - 1
			}
		} else {
			cursor.CellOffX = cursor.cellOffXVertical
		}
	}
}

func (p *EditorView) MoveCursorLeftmost(cursor *EditorViewCursor, line *EditorLine) {
	cursor.CellOffX = 0

	if len(line.Cells) == 0 {
		return
	}

	cell := line.Cells[cursor.CellOffX]
	cursor.UISetCursor(cell.X, cell.Y)
	cursor.cellOffXVertical = cursor.CellOffX
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
	cursor.cellOffXVertical = cursor.CellOffX
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
		case EditorCommandMode:
			cursor.CellOffX = len(line.Cells) - 1
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)

		case EditorInputMode:
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.UIWidth, cell.Y)
		}

	} else {
		cell := line.Cells[cursor.CellOffX]
		cursor.UISetCursor(cell.X, cell.Y)
	}

	cursor.cellOffXVertical = cursor.CellOffX
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
		if cell.X+cell.UIWidth >= p.EditorView.Block.InnerArea.Max.X {
			x, y = line.ContentStartX, cell.Y+1
		} else {
			x, y = cell.X+cell.UIWidth, cell.Y
		}

		if y >= p.EditorView.Block.InnerArea.Max.Y {
			y = p.EditorView.Block.InnerArea.Max.Y - 1
			switch p.EditorView.Mode {
			case EditorInputMode:
				p.DisplayLinesTopIndex++
				p.EditorView.isShouldRefreshInputModeBuf = true
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
