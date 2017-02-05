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
		cursor.RefreshCursorByEditorLine(line)
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
		cursor.RefreshCursorByEditorLine(line)
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
		cursor.RefreshCursorByEditorLine(line)
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
		cursor.RefreshCursorByEditorLine(line)
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
		case EDITOR_NORMAL_MODE:
			cursor.CellOffX = len(line.Cells) - 1
			cell := line.Cells[cursor.CellOffX]
			cursor.UISetCursor(cell.X, cell.Y)

		case EDITOR_EDIT_MODE:
			cursor.CellOffX = len(line.Cells)
			cell := line.Cells[cursor.CellOffX-1]
			cursor.UISetCursor(cell.X+cell.Width(), cell.Y)

		case EDITOR_COMMAND_MODE:
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
	case EDITOR_EDIT_MODE:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EDITOR_NORMAL_MODE:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EDITOR_COMMAND_MODE:
		p.CommandModeCursor.RefreshCursorByEditorLine(p.CommandModeBuf)
	}
}

func (p *EditorCursor) RefreshCursorByEditorLine(line *EditorLine) {
	if nil == line {
		p.UISetCursor(p.Editor.Block.InnerArea.Min.X, p.Editor.Block.InnerArea.Min.Y)
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
			case EDITOR_EDIT_MODE:
				p.DisplayLinesTopIndex += 1
				p.Editor.isShouldRefreshEditModeBuf = true
			}
		}

	} else {
		cell = line.Cells[p.CellOffX]
		x, y = cell.X, cell.Y
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
