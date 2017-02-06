package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorCommandCursor struct {
	Editor *Editor
	*EditorCursor
}

func NewEditorCommandCursor(editor *Editor) *EditorCommandCursor {
	ret := &EditorCommandCursor{
		Editor:       editor,
		EditorCursor: NewEditorCursor(),
	}
	return ret
}

func (p *Editor) MoveCommandCursorLeft(cursor *EditorCommandCursor, line *EditorLine, runenum int) {
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
}

func (p *Editor) MoveCommandCursorRight(cursor *EditorCommandCursor, line *EditorLine, runenum int) {
	if runenum <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursor.CellOffX = 0
		return
	}

	cursor.CellOffX += runenum
	if cursor.CellOffX >= len(line.Cells) {
		cursor.CellOffX = len(line.Cells)
		cell := line.Cells[cursor.CellOffX-1]
		cursor.UISetCursor(cell.X+cell.Width(), cell.Y)

	} else {
		cell := line.Cells[cursor.CellOffX]
		cursor.UISetCursor(cell.X, cell.Y)
	}
}

func (p *EditorCommandCursor) RefreshCursorByEditorLine(line *EditorLine) {
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
		if cell.X+width >= p.Editor.Block.InnerArea.Max.X {
			x, y = line.ContentStartX, cell.Y+1
		} else {
			x, y = cell.X+width, cell.Y
		}

		if y >= p.Editor.Block.InnerArea.Max.Y {
			y = p.Editor.Block.InnerArea.Max.Y - 1
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

func (p *EditorCommandCursor) UISetCursor(x, y int) {
	utils.UISetCursor(x, y)
}
