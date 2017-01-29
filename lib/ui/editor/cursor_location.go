package editor

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

type CursorLocation struct {
	IsDisplay bool
	Editor    *Editor

	EditModeOffXCellIndex              int
	offXCellIndexForVerticalMoveCursor int

	CommandModeOffXCellIndex int
}

func NewCursorLocation(editor *Editor) *CursorLocation {
	ret := &CursorLocation{
		IsDisplay:             false,
		EditModeOffXCellIndex: 0,
		Editor:                editor,
	}
	return ret
}

func (p *CursorLocation) getOffXCellIndex() *int {
	switch p.Editor.Mode {
	case EDITOR_NORMAL_MODE:
		return &p.EditModeOffXCellIndex
	case EDITOR_EDIT_MODE:
		return &p.EditModeOffXCellIndex
	case EDITOR_COMMAND_MODE:
		return &p.CommandModeOffXCellIndex
	}
	return nil
}

func (p *CursorLocation) getLineByMode() *Line {
	switch p.Editor.Mode {
	case EDITOR_NORMAL_MODE:
		return p.Editor.CurrentLine()
	case EDITOR_EDIT_MODE:
		return p.Editor.CurrentLine()
	case EDITOR_COMMAND_MODE:
		return p.Editor.CommandModeBuf
	}
	return nil
}

func (p *CursorLocation) MoveCursorNRuneTop(n int) {
	if n <= 0 {
		return
	}

	offXCellIndex := p.getOffXCellIndex()
	if nil == offXCellIndex {
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
		if *offXCellIndex >= len(p.Editor.CurrentLine().Cells) {
			*offXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
		}

		cell := p.Editor.CurrentLine().Cells[*offXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
	}
}

func (p *CursorLocation) MoveCursorNRuneBottom(n int) {
	if n <= 0 {
		return
	}

	offXCellIndex := p.getOffXCellIndex()
	if nil == offXCellIndex {
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
		if *offXCellIndex >= len(p.Editor.CurrentLine().Cells) {
			*offXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
		}

		cell := p.Editor.CurrentLine().Cells[*offXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
	}
}

func (p *CursorLocation) MoveCursorNRuneLeft(n int) {
	if n <= 0 {
		return
	}

	offXCellIndex := p.getOffXCellIndex()
	if nil == offXCellIndex {
		return
	}

	line := p.getLineByMode()
	if len(line.Cells) == 0 {
		*offXCellIndex = 0
		return
	}

	*offXCellIndex -= n
	if *offXCellIndex < 0 {
		*offXCellIndex = 0
	}

	cell := p.getLineByMode().Cells[*offXCellIndex]
	p.UISetCursor(cell.X, cell.Y)
}

func (p *CursorLocation) MoveCursorNRuneRight(n int) {
	if n <= 0 {
		return
	}

	offXCellIndex := p.getOffXCellIndex()
	if nil == offXCellIndex {
		return
	}

	line := p.getLineByMode()
	if len(line.Cells) == 0 {
		*offXCellIndex = 0
		return
	}

	*offXCellIndex += n
	if *offXCellIndex >= len(line.Cells) {
		switch p.Editor.Mode {
		case EDITOR_NORMAL_MODE:
			*offXCellIndex = len(line.Cells) - 1
			cell := line.Cells[*offXCellIndex]
			p.UISetCursor(cell.X, cell.Y)

		case EDITOR_EDIT_MODE:
			*offXCellIndex = len(line.Cells) - 1
			cell := line.Cells[*offXCellIndex]
			p.UISetCursor(cell.X, cell.Y)

		case EDITOR_COMMAND_MODE:
			*offXCellIndex = len(line.Cells)
			cell := line.Cells[*offXCellIndex-1]
			p.UISetCursor(cell.X+cell.Width(), cell.Y)
		}

	} else {
		cell := line.Cells[*offXCellIndex]
		p.UISetCursor(cell.X, cell.Y)
	}
}

func (p *CursorLocation) RefreshCursorByLine(line *Line) {
	offXCellIndex := p.getOffXCellIndex()
	if nil == offXCellIndex {
		return
	}

	if nil == line {
		p.UISetCursor(p.Editor.Block.InnerArea.Min.X, p.Editor.Block.InnerArea.Min.Y)
		return
	}

	if 0 == len(line.Cells) {
		p.UISetCursor(line.ContentStartX, line.ContentStartY)
		return
	}

	if *offXCellIndex >= len(line.Cells) {
		*offXCellIndex = len(line.Cells)
	}

	var x, y int
	var cell termui.Cell
	if *offXCellIndex == len(line.Cells) {
		cell = line.Cells[*offXCellIndex-1]
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
		cell = line.Cells[*offXCellIndex]
		x, y = cell.X, cell.Y
	}

	p.UISetCursor(x, y)
}

func (p *CursorLocation) UISetCursor(x, y int) {
	if false == p.IsDisplay {
		return
	}
	uiutils.UISetCursor(x, y)
}
