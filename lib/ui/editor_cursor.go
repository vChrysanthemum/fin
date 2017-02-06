package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorCursor struct {
	Editor *Editor

	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int

	LineIndex        int
	CellOffX         int
	CellOffXVertical int
}

func NewEditorCursor(editor *Editor) *EditorCursor {
	ret := &EditorCursor{
		CellOffX:         0,
		CellOffXVertical: 0,
		Editor:           editor,
	}
	return ret
}

func (p *Editor) EditModeMoveCursorNRuneUp(cursor *EditorCursor, n int) {
	if n <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - n
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

func (p *Editor) NormalModeMoveCursorNRuneUp(cursor *EditorCursor, n int) {
	if n <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex - n
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

func (p *Editor) EditModeMoveCursorNRuneDown(cursor *EditorCursor, n int) {
	if n <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + n
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

func (p *Editor) NormalModeMoveCursorNRuneDown(cursor *EditorCursor, n int) {
	if n <= 0 {
		return
	}

	if cursor.CellOffX > cursor.CellOffXVertical {
		cursor.CellOffXVertical = cursor.CellOffX
	}

	index := cursor.LineIndex + n
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

func (p *Editor) MoveCursorNRuneLeft(cursor *EditorCursor, line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursor.CellOffX = 0
		return
	}

	cursor.CellOffX -= n
	if cursor.CellOffX < 0 {
		cursor.CellOffX = 0
	}

	cell := line.Cells[cursor.CellOffX]
	cursor.UISetCursor(cell.X, cell.Y)

	cursor.CellOffXVertical = cursor.CellOffX
}

func (p *Editor) MoveCursorNRuneRight(cursor *EditorCursor, line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursor.CellOffX = 0
		return
	}

	cursor.CellOffX += n
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

		case EditorCommandMode:
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

func (p *Editor) RefreshCursorByEditorLine() {
	switch p.Mode {
	case EditorEditMode:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EditorNormalMode:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EditorCommandMode:
		p.CommandModeCursor.RefreshCursorByEditorLine(p.CommandModeBuf)
	}
}

func (p *EditorCursor) RefreshCursorByEditorLine(line *EditorLine) {
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
			switch p.Editor.Mode {
			case EditorEditMode:
				p.DisplayLinesTopIndex++
				p.Editor.isShouldRefreshEditModeBuf = true
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

func (p *EditorCursor) UISetCursor(x, y int) {
	utils.UISetCursor(x, y)
}

func (p *EditorCursor) Line() *EditorLine {
	if p.LineIndex >= len(p.Editor.Lines) {
		return nil
	}
	return p.Editor.Lines[p.LineIndex]
}
