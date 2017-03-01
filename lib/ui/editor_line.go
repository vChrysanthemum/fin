package ui

import (
	"fmt"
	"strconv"

	"github.com/gizak/termui"
)

type EditorLine struct {
	ContentStartX, ContentStartY int

	Index      int
	EditorView *EditorView
	Data       []byte
	Cells      []termui.Cell
}

func (p *Editor) NewLine(editorView *EditorView) *EditorLine {
	return &EditorLine{
		EditorView:    editorView,
		ContentStartX: p.Block.InnerArea.Min.X,
		ContentStartY: p.Block.InnerArea.Min.Y,
	}
}

func (p *EditorView) InputModeAppendNewLine(inputModeCursor *EditorViewCursor) *EditorLine {
	ret := p.Editor.NewLine(p)

	if 0 == len(p.Lines) {
		ret.Index = 0
		p.Lines = append(p.Lines, ret)
		inputModeCursor.LineIndex = ret.Index

	} else if inputModeCursor.LineIndex == len(p.Lines)-1 {
		ret.Index = len(p.Lines)
		p.Lines = append(p.Lines, ret)
		inputModeCursor.LineIndex = ret.Index

	} else {
		for _, line := range p.Lines[inputModeCursor.LineIndex+1:] {
			line.Index++
		}
		ret.Index = inputModeCursor.LineIndex + 1

		n := len(p.Lines) + 1
		if cap(p.Lines) < n {
			_lines := make([]*EditorLine, len(p.Lines), n)
			copy(_lines, p.Lines)
			p.Lines = _lines
		}
		p.Lines = p.Lines[:n]

		copy(p.Lines[inputModeCursor.LineIndex+2:], p.Lines[inputModeCursor.LineIndex+1:])
		copy(p.Lines[inputModeCursor.LineIndex+1:], []*EditorLine{ret})
		inputModeCursor.LineIndex = ret.Index
	}

	if inputModeCursor.LineIndex > 0 {
		line := p.Lines[inputModeCursor.LineIndex-1]
		if inputModeCursor.CellOffX < len(line.Cells) {
			inputModeCursor.Line().Data =
				make([]byte, len(line.Data[line.Cells[inputModeCursor.CellOffX].BytesOff:]))

			copy(
				inputModeCursor.Line().Data,
				line.Data[line.Cells[inputModeCursor.CellOffX].BytesOff:])

			line.Data = line.Data[:line.Cells[inputModeCursor.CellOffX].BytesOff]
			inputModeCursor.CellOffX = 0
		}
	}

	if true == p.isDisplayEditorLineNumber {
		ret.ContentStartX = p.Block.InnerArea.Min.X +
			len(ret.getEditorLinePrefix(len(p.Lines), len(p.Lines)))
	}

	return ret
}

// InputModeReduceLine 缩减指定行
// 该操作将指定行数据追加到上一行中，然后删除指定行
func (p *EditorView) InputModeReduceLine(lineIndex int) {
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

// RemoveLines 删除 EditorView.Lines 中的几行
func (p *EditorView) RemoveLines(lineIndex, linesNum int) {
	if linesNum <= 0 || lineIndex >= len(p.Lines) ||
		(1 == len(p.Lines) && 0 == len(p.Lines[0].Data)) {
		return
	}

	if lineIndex+linesNum > len(p.Lines) {
		linesNum = len(p.Lines) - lineIndex
	}

	for _, line := range p.Lines[lineIndex+linesNum:] {
		line.Index -= linesNum
	}

	if 0 == lineIndex && 1 == linesNum {
		p.Lines = []*EditorLine{p.Lines[0]}
		p.Lines[0].Data = []byte{}
	} else if linesNum == len(p.Lines) {
		p.Lines = []*EditorLine{p.Editor.NewLine(p)}
	} else {
		p.Lines = append(p.Lines[:lineIndex], p.Lines[lineIndex+linesNum:]...)
	}
}

// InsertLines 在 EditorView.Lines 中插入几行数据
func (p *EditorView) InsertLines(lineIndex int, lines []EditorLine) {
	for _, line := range p.Lines[lineIndex:] {
		line.Index += len(lines)
	}

	paramlines := make([]*EditorLine, len(lines))
	for k, line := range lines {
		paramlines[k] = line.Copy()
		paramlines[k].Index = lineIndex + k
	}

	tmpLines := make([]*EditorLine, len(p.Lines[lineIndex:]))
	copy(tmpLines, p.Lines[lineIndex:])
	p.Lines = append(p.Lines[:lineIndex], paramlines...)
	p.Lines = append(p.Lines, tmpLines...)
}

// AppendLineData 在 EditorView.Lines 末尾插入 一行数据
func (p *EditorView) AppendLineData(data []byte) {
	line := p.Editor.NewLine(p)
	line.Index = len(p.Lines)
	line.Data = data
	p.Lines = append(p.Lines, line)
}

// InsertPointerLines 在 EditorView.Lines 中插入几行数据
func (p *EditorView) InsertPointerLines(lineIndex int, lines []*EditorLine) {
	for _, line := range p.Lines[lineIndex:] {
		line.Index += len(lines)
	}

	for k, line := range lines {
		line.Index = lineIndex + k
	}

	_lines := make([]*EditorLine, len(p.Lines[lineIndex:]))
	copy(_lines, p.Lines[lineIndex:])
	p.Lines = append(p.Lines[:lineIndex], lines...)
	p.Lines = append(p.Lines, _lines...)
}

// getEditorLinePrefix 获取 line 内容前缀
//
// param:
//		lineIndex				int		 目标 line 的相应 EditorView.Lines 中的index
//		lastEditorLineIndex		int		 输出 line 的前缀需要与整个页面的 line 前缀对齐
//									 	 lastEditorLineIndex 为 页面中最后一条 line 的 index
func (p *EditorLine) getEditorLinePrefix(lineIndex, lastEditorLineIndex int) string {
	if true == p.EditorView.isDisplayEditorLineNumber {
		lineNumberWidth := len(strconv.Itoa(lastEditorLineIndex + 1))
		if lineNumberWidth < 3 {
			lineNumberWidth = 3
		}

		return fmt.Sprintf(fmt.Sprintf("%%%ds ", lineNumberWidth), strconv.Itoa(lineIndex+1))
	}

	return ""
}

func (p *EditorLine) Copy() *EditorLine {
	ret := &EditorLine{
		ContentStartX: p.ContentStartX,
		ContentStartY: p.ContentStartY,
		Index:         p.Index,
		EditorView:    p.EditorView,
	}
	ret.Data = make([]byte, len(p.Data))
	copy(ret.Data, p.Data)
	ret.Cells = make([]termui.Cell, len(p.Cells))
	copy(ret.Cells, p.Cells)
	return ret
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
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.EditorView.TextFgColor, p.EditorView.TextBgColor)
	}
}

func (p *EditorLine) LastLineModeBackspace(lastLineModeCursor *EditorCommandCursor) {
	if lastLineModeCursor.CellOffX > len(p.Cells) {
		lastLineModeCursor.CellOffX = len(p.Cells)
	}

	if lastLineModeCursor.CellOffX <= 1 {
		return
	}

	if lastLineModeCursor.CellOffX == len(p.Cells) {
		p.Data = p.Data[:p.Cells[lastLineModeCursor.CellOffX-1].BytesOff]
		lastLineModeCursor.CellOffX--

	} else {
		p.Data = append(p.Data[:p.Cells[lastLineModeCursor.CellOffX-1].BytesOff],
			p.Data[p.Cells[lastLineModeCursor.CellOffX].BytesOff:]...)
		lastLineModeCursor.CellOffX--
	}
}

func (p *EditorLine) InputModeBackspace(inputModeCursor *EditorViewCursor) {
	if inputModeCursor.CellOffX > len(p.Cells) {
		inputModeCursor.CellOffX = len(p.Cells)
	}

	if len(p.EditorView.Lines) <= 1 && 0 == inputModeCursor.CellOffX {
		return
	}

	if 0 == inputModeCursor.CellOffX {
		if inputModeCursor.LineIndex > 0 {
			p.EditorView.InputModeReduceLine(inputModeCursor.LineIndex)
			inputModeCursor.LineIndex--
			inputModeCursor.CellOffX = len(inputModeCursor.Line().Cells)
		}

	} else if inputModeCursor.CellOffX == len(p.Cells) {
		p.Data = p.Data[:p.Cells[inputModeCursor.CellOffX-1].BytesOff]
		inputModeCursor.CellOffX--

	} else {
		p.Data = append(p.Data[:p.Cells[inputModeCursor.CellOffX-1].BytesOff],
			p.Data[p.Cells[inputModeCursor.CellOffX].BytesOff:]...)
		inputModeCursor.CellOffX--
	}
}

func (p *EditorLine) Write(cursor *EditorCursor, keyStr string) {
	_off := 0

	if cursor.CellOffX >= len(p.Cells) {
		_off = len(p.Data)
		p.Data = append(p.Data, []byte(keyStr)...)
		p.Cells = append(p.Cells, termui.Cell{[]rune(keyStr)[0], p.EditorView.TextFgColor, p.EditorView.TextBgColor, 0, 0, _off})

	} else if 0 == cursor.CellOffX {
		_off = 0
		p.Data = append([]byte(keyStr), p.Data...)
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.EditorView.TextFgColor, p.EditorView.TextBgColor)

	} else {
		newData := make([]byte, len(p.Data)+len(keyStr))
		_off = p.Cells[cursor.CellOffX].BytesOff
		copy(newData, p.Data[:_off])
		copy(newData[_off:], []byte(keyStr))
		copy(newData[_off+len(keyStr):], p.Data[_off:])
		p.Data = newData
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.EditorView.TextFgColor, p.EditorView.TextBgColor)
	}

	cursor.CellOffX++
}

func (p *EditorLine) CleanData(inputModeCursor *EditorCursor) {
	inputModeCursor.CellOffX = 0
	p.Data = []byte{}
	p.Cells = []termui.Cell{}
}
