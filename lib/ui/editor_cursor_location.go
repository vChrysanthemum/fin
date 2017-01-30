package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorCursorLocation struct {
	Editor *Editor

	OffXCellIndex         int
	OffXCellIndexVertical int
}

func NewEditorCursorLocation(editor *Editor) *EditorCursorLocation {
	ret := &EditorCursorLocation{
		OffXCellIndex:         0,
		OffXCellIndexVertical: 0,
		Editor:                editor,
	}
	return ret
}

func (p *Editor) MoveCursorNRuneTop(cursorLocation *EditorCursorLocation, n int) {
	if n <= 0 {
		return
	}

	index := p.CurrentLineIndex - n
	if index < 0 {
		index = 0
	}

	p.CurrentLineIndex = index

	if index < p.DisplayLinesTopIndex {
		p.DisplayLinesTopIndex = index
		p.isEditModeBufDirty = true
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(p.CurrentLine().Cells) {
		cursorLocation.UISetCursor(p.CurrentLine().ContentStartX, p.CurrentLine().ContentStartY)

	} else {
		if cursorLocation.OffXCellIndex >= len(p.CurrentLine().Cells) {
			cursorLocation.OffXCellIndex = len(p.CurrentLine().Cells) - 1
		}

		cell := p.CurrentLine().Cells[cursorLocation.OffXCellIndex]
		cursorLocation.UISetCursor(cell.X, cell.Y)
	}
}

func (p *Editor) MoveCursorNRuneBottom(cursorLocation *EditorCursorLocation, n int) {
	if n <= 0 {
		return
	}

	index := p.CurrentLineIndex + n
	if index >= len(p.Lines) {
		index = len(p.Lines) - 1
	}

	p.CurrentLineIndex = index

	if index > p.DisplayLinesBottomIndex {
		p.DisplayLinesTopIndex += (index - p.DisplayLinesBottomIndex)
		p.isEditModeBufDirty = true
		p.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(p.CurrentLine().Cells) {
		cursorLocation.UISetCursor(p.CurrentLine().ContentStartX, p.CurrentLine().ContentStartY)

	} else {
		if cursorLocation.OffXCellIndex >= len(p.CurrentLine().Cells) {
			cursorLocation.OffXCellIndex = len(p.CurrentLine().Cells) - 1
		}

		cell := p.CurrentLine().Cells[cursorLocation.OffXCellIndex]
		cursorLocation.UISetCursor(cell.X, cell.Y)
	}
}

func (p *Editor) MoveCursorNRuneLeft(cursorLocation *EditorCursorLocation, line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursorLocation.OffXCellIndex = 0
		return
	}

	cursorLocation.OffXCellIndex -= n
	if cursorLocation.OffXCellIndex < 0 {
		cursorLocation.OffXCellIndex = 0
	}

	cell := line.Cells[cursorLocation.OffXCellIndex]
	cursorLocation.UISetCursor(cell.X, cell.Y)
}

func (p *Editor) MoveCursorNRuneRight(cursorLocation *EditorCursorLocation, line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		cursorLocation.OffXCellIndex = 0
		return
	}

	cursorLocation.OffXCellIndex += n
	if cursorLocation.OffXCellIndex >= len(line.Cells) {
		switch p.Mode {
		case EDITOR_NORMAL_MODE:
			cursorLocation.OffXCellIndex = len(line.Cells) - 1
			cell := line.Cells[cursorLocation.OffXCellIndex]
			cursorLocation.UISetCursor(cell.X, cell.Y)

		case EDITOR_EDIT_MODE:
			cursorLocation.OffXCellIndex = len(line.Cells)
			cell := line.Cells[cursorLocation.OffXCellIndex-1]
			cursorLocation.UISetCursor(cell.X+cell.Width(), cell.Y)

		case EDITOR_COMMAND_MODE:
			cursorLocation.OffXCellIndex = len(line.Cells)
			cell := line.Cells[cursorLocation.OffXCellIndex-1]
			cursorLocation.UISetCursor(cell.X+cell.Width(), cell.Y)
		}

	} else {
		cell := line.Cells[cursorLocation.OffXCellIndex]
		cursorLocation.UISetCursor(cell.X, cell.Y)
	}
}

func (p *EditorCursorLocation) RefreshCursorByEditorLine(line *EditorLine) {
	if nil == line {
		p.UISetCursor(p.Editor.Block.InnerArea.Min.X, p.Editor.Block.InnerArea.Min.Y)
		return
	}

	if 0 == len(line.Cells) {
		p.UISetCursor(line.ContentStartX, line.ContentStartY)
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
			switch p.Editor.Mode {
			case EDITOR_EDIT_MODE:
				p.Editor.DisplayLinesTopIndex += 1
				p.Editor.isShouldRefreshEditModeBuf = true
			}
		}

	} else {
		cell = line.Cells[p.OffXCellIndex]
		x, y = cell.X, cell.Y
	}

	p.UISetCursor(x, y)
}

func (p *EditorCursorLocation) UISetCursor(x, y int) {
	uiutils.UISetCursor(x, y)
}
