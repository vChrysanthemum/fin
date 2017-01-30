package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

type EditorCursorLocation struct {
	IsDisplay bool
	Editor    *Editor

	OffXCellIndex         int
	OffXCellIndexVertical int
}

func NewEditorCursorLocation(editor *Editor) *EditorCursorLocation {
	ret := &EditorCursorLocation{
		IsDisplay:             false,
		OffXCellIndex:         0,
		OffXCellIndexVertical: 0,
		Editor:                editor,
	}
	return ret
}

func (p *EditorCursorLocation) MoveCursorNRuneTop(n int) {
	if n <= 0 {
		return
	}

	index := p.Editor.CurrentLineIndex - n
	if index < 0 {
		index = 0
	}

	p.Editor.CurrentLineIndex = index

	if index < p.Editor.DisplayLinesTopIndex {
		p.Editor.DisplayLinesTopIndex = index
		p.Editor.isEditModeBufDirty = true
		p.Editor.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(p.Editor.CurrentLine().Cells) {
		p.UISetCursor(p.Editor.CurrentLine().ContentStartX, p.Editor.CurrentLine().ContentStartY)

	} else {
		if p.OffXCellIndex >= len(p.Editor.CurrentLine().Cells) {
			p.OffXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
		}

		cell := p.Editor.CurrentLine().Cells[p.OffXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
	}
}

func (p *EditorCursorLocation) MoveCursorNRuneBottom(n int) {
	if n <= 0 {
		return
	}

	index := p.Editor.CurrentLineIndex + n
	if index >= len(p.Editor.Lines) {
		index = len(p.Editor.Lines) - 1
	}

	p.Editor.CurrentLineIndex = index

	if index > p.Editor.DisplayLinesBottomIndex {
		p.Editor.DisplayLinesTopIndex += (index - p.Editor.DisplayLinesBottomIndex)
		p.Editor.isEditModeBufDirty = true
		p.Editor.isShouldRefreshEditModeBuf = true
	}

	if 0 == len(p.Editor.CurrentLine().Cells) {
		p.UISetCursor(p.Editor.CurrentLine().ContentStartX, p.Editor.CurrentLine().ContentStartY)

	} else {
		if p.OffXCellIndex >= len(p.Editor.CurrentLine().Cells) {
			p.OffXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
		}

		cell := p.Editor.CurrentLine().Cells[p.OffXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
	}
}

func (p *EditorCursorLocation) MoveCursorNRuneLeft(line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		p.OffXCellIndex = 0
		return
	}

	p.OffXCellIndex -= n
	if p.OffXCellIndex < 0 {
		p.OffXCellIndex = 0
	}

	cell := line.Cells[p.OffXCellIndex]
	p.UISetCursor(cell.X, cell.Y)
}

func (p *EditorCursorLocation) MoveCursorNRuneRight(line *EditorLine, n int) {
	if n <= 0 {
		return
	}

	if len(line.Cells) == 0 {
		p.OffXCellIndex = 0
		return
	}

	p.OffXCellIndex += n
	if p.OffXCellIndex >= len(line.Cells) {
		switch p.Editor.Mode {
		case EDITOR_NORMAL_MODE:
			p.OffXCellIndex = len(line.Cells) - 1
			cell := line.Cells[p.OffXCellIndex]
			p.UISetCursor(cell.X, cell.Y)

		case EDITOR_EDIT_MODE:
			p.OffXCellIndex = len(line.Cells) - 1
			cell := line.Cells[p.OffXCellIndex]
			p.UISetCursor(cell.X, cell.Y)

		case EDITOR_COMMAND_MODE:
			p.OffXCellIndex = len(line.Cells)
			cell := line.Cells[p.OffXCellIndex-1]
			p.UISetCursor(cell.X+cell.Width(), cell.Y)
		}

	} else {
		cell := line.Cells[p.OffXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
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
	if false == p.IsDisplay {
		return
	}
	uiutils.UISetCursor(x, y)
}
