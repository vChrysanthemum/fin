package ui

import (
	"fmt"
	"strconv"

	"github.com/gizak/termui"
)

type EditorLine struct {
	ContentStartX, ContentStartY int

	Index  int
	Editor *Editor
	Data   []byte
	Cells  []termui.Cell
}

func (p *Editor) NewLine() *EditorLine {
	return &EditorLine{
		Editor:        p,
		ContentStartX: p.Block.InnerArea.Min.X,
		ContentStartY: p.Block.InnerArea.Min.Y,
		Data:          make([]byte, 0),
	}
}

func (p *Editor) EditModeAppendNewLine(editModeCursor *EditorCursor) *EditorLine {
	ret := p.NewLine()

	if 0 == len(p.Lines) {
		ret.Index = 0
		p.Lines = append(p.Lines, ret)
		editModeCursor.LineIndex = ret.Index

	} else if editModeCursor.LineIndex == len(p.Lines)-1 {
		ret.Index = len(p.Lines)
		p.Lines = append(p.Lines, ret)
		editModeCursor.LineIndex = ret.Index

	} else {
		for _, line := range p.Lines[editModeCursor.LineIndex+1:] {
			line.Index++
		}
		ret.Index = editModeCursor.LineIndex + 1

		n := len(p.Lines) + 1
		if cap(p.Lines) < n {
			_lines := make([]*EditorLine, len(p.Lines), n)
			copy(_lines, p.Lines)
			p.Lines = _lines
		}
		p.Lines = p.Lines[:n]

		copy(p.Lines[editModeCursor.LineIndex+2:], p.Lines[editModeCursor.LineIndex+1:])
		copy(p.Lines[editModeCursor.LineIndex+1:], []*EditorLine{ret})
		editModeCursor.LineIndex = ret.Index
	}

	if editModeCursor.LineIndex > 0 {
		line := p.Lines[editModeCursor.LineIndex-1]
		if editModeCursor.CellOffX < len(line.Cells) {
			editModeCursor.Line().Data =
				make([]byte, len(line.Data[line.Cells[editModeCursor.CellOffX].BytesOff:]))

			copy(
				editModeCursor.Line().Data,
				line.Data[line.Cells[editModeCursor.CellOffX].BytesOff:])

			line.Data = line.Data[:line.Cells[editModeCursor.CellOffX].BytesOff]
			editModeCursor.CellOffX = 0
		}
	}

	if true == p.isDisplayEditorLineNumber {
		ret.ContentStartX = p.Block.InnerArea.Min.X +
			len(ret.getEditorLinePrefix(len(p.Lines), len(p.Lines)))
	}

	return ret
}

// EditModeReduceLine 缩减指定行
// 该操作将指定行数据追加到上一行中，然后删除指定行
func (p *Editor) EditModeReduceLine(lineIndex int) {
	var line *EditorLine

	if lineIndex <= 0 || lineIndex >= len(p.Lines) {
		return
	}

	for _, line = range p.Lines[lineIndex:] {
		line.Index--
	}

	line = p.Lines[lineIndex]
	prevLine := p.Lines[lineIndex-1]

	p.Lines = append(p.Lines[:lineIndex], p.Lines[lineIndex+1:]...)
	prevLine.Data = append(prevLine.Data, line.Data...)

	return
}

// 获取 line 内容前缀
//
// param:
//		lineIndex				int		 目标 line 的相应 Editor.Lines 中的index
//		lastEditorLineIndex		int		 输出 line 的前缀需要与整个页面的 line 前缀对齐
//									 	 lastEditorLineIndex 为 页面中最后一条 line 的 index
func (p *EditorLine) getEditorLinePrefix(lineIndex, lastEditorLineIndex int) string {
	if true == p.Editor.isDisplayEditorLineNumber {
		lineNumberWidth := len(strconv.Itoa(lastEditorLineIndex + 1))
		if lineNumberWidth < 3 {
			lineNumberWidth = 3
		}

		return fmt.Sprintf(fmt.Sprintf("%%%ds ", lineNumberWidth), strconv.Itoa(lineIndex+1))
	}

	return ""
}

func (p *EditorLine) Write(cursor *EditorCursor, keyStr string) {
	_off := 0

	if cursor.CellOffX >= len(p.Cells) {
		_off = len(p.Data)
		p.Data = append(p.Data, []byte(keyStr)...)
		p.Cells = append(p.Cells, termui.Cell{[]rune(keyStr)[0], p.Editor.TextFgColor, p.Editor.TextBgColor, 0, 0, _off})

	} else if 0 == cursor.CellOffX {
		_off = 0
		p.Data = append([]byte(keyStr), p.Data...)
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.Editor.TextFgColor, p.Editor.TextBgColor)

	} else {
		newData := make([]byte, len(p.Data)+len(keyStr))
		_off = p.Cells[cursor.CellOffX].BytesOff
		copy(newData, p.Data[:_off])
		copy(newData[_off:], []byte(keyStr))
		copy(newData[_off+len(keyStr):], p.Data[_off:])
		p.Data = newData
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.Editor.TextFgColor, p.Editor.TextBgColor)
	}

	cursor.CellOffX++
}

func (p *EditorLine) CommandModeBackspace(commandModeCursor *EditorCursor) {
	if commandModeCursor.CellOffX > len(p.Cells) {
		commandModeCursor.CellOffX = len(p.Cells)
	}

	if commandModeCursor.CellOffX <= 1 {
		return
	}

	if commandModeCursor.CellOffX == len(p.Cells) {
		p.Data = p.Data[:p.Cells[commandModeCursor.CellOffX-1].BytesOff]
		commandModeCursor.CellOffX--

	} else {
		p.Data = append(p.Data[:p.Cells[commandModeCursor.CellOffX-1].BytesOff],
			p.Data[p.Cells[commandModeCursor.CellOffX].BytesOff:]...)
		commandModeCursor.CellOffX--
	}
}

func (p *EditorLine) EditModeBackspace(editModeCursor *EditorCursor) {
	if editModeCursor.CellOffX > len(p.Cells) {
		editModeCursor.CellOffX = len(p.Cells)
	}

	if len(p.Editor.Lines) <= 1 && 0 == editModeCursor.CellOffX {
		return
	}

	if 0 == editModeCursor.CellOffX {
		if editModeCursor.LineIndex > 0 {
			p.Editor.EditModeReduceLine(editModeCursor.LineIndex)
			editModeCursor.LineIndex--
			editModeCursor.CellOffX = len(editModeCursor.Line().Cells)
		}

	} else if editModeCursor.CellOffX == len(p.Cells) {
		p.Data = p.Data[:p.Cells[editModeCursor.CellOffX-1].BytesOff]
		editModeCursor.CellOffX--

	} else {
		p.Data = append(p.Data[:p.Cells[editModeCursor.CellOffX-1].BytesOff],
			p.Data[p.Cells[editModeCursor.CellOffX].BytesOff:]...)
		editModeCursor.CellOffX--
	}
}

func (p *EditorLine) CutAway(offStart, offEnd int) {
	if offEnd > offStart {
		if offEnd >= len(p.Data) {
			p.Data = p.Data[:offStart]
		} else if 0 == offStart {
			p.Data = p.Data[offEnd:]
		} else {
			p.Data = append(p.Data[:offStart], p.Data[offEnd:]...)
		}
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.Editor.TextFgColor, p.Editor.TextBgColor)
	}
}

func (p *EditorLine) CleanData(editModeCursor *EditorCursor) {
	editModeCursor.CellOffX = 0
	p.Data = []byte{}
	p.Cells = []termui.Cell{}
}
