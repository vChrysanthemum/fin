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

func (p *Editor) NewEditorLine() *EditorLine {
	return &EditorLine{
		Editor:        p,
		ContentStartX: p.Block.InnerArea.Min.X,
		ContentStartY: p.Block.InnerArea.Min.Y,
		Data:          make([]byte, 0),
	}
}

func (p *Editor) EditorEditModeAppendNewEditorLine() *EditorLine {
	ret := p.NewEditorLine()

	if 0 == len(p.EditorLines) {
		ret.Index = 0
		p.EditorLines = append(p.EditorLines, ret)
		p.CurrentEditorLineIndex = ret.Index

	} else if p.CurrentEditorLineIndex == len(p.EditorLines)-1 {
		ret.Index = len(p.EditorLines)
		p.EditorLines = append(p.EditorLines, ret)
		p.CurrentEditorLineIndex = ret.Index

	} else {
		for _, line := range p.EditorLines[p.CurrentEditorLineIndex+1:] {
			line.Index += 1
		}
		ret.Index = p.CurrentEditorLineIndex + 1

		n := len(p.EditorLines) + 1
		if cap(p.EditorLines) < n {
			_lines := make([]*EditorLine, len(p.EditorLines), n)
			copy(_lines, p.EditorLines)
			p.EditorLines = _lines
		}
		p.EditorLines = p.EditorLines[:n]

		copy(p.EditorLines[p.CurrentEditorLineIndex+2:], p.EditorLines[p.CurrentEditorLineIndex+1:])
		copy(p.EditorLines[p.CurrentEditorLineIndex+1:], []*EditorLine{ret})
		p.CurrentEditorLineIndex = ret.Index
	}

	if p.CurrentEditorLineIndex > 0 {
		line := p.EditorLines[p.CurrentEditorLineIndex-1]
		if p.EditorEditModeOffXCellIndex < len(line.Cells) {
			p.CurrentEditorLine().Data = make([]byte, len(line.Data[line.Cells[p.EditorEditModeOffXCellIndex].BytesOff:]))
			copy(p.CurrentEditorLine().Data, line.Data[line.Cells[p.EditorEditModeOffXCellIndex].BytesOff:])
			line.Data = line.Data[:line.Cells[p.EditorEditModeOffXCellIndex].BytesOff]
			p.EditorEditModeOffXCellIndex = 0
		}
	}

	if true == p.isDisplayEditorLineNumber {
		ret.ContentStartX = p.Block.InnerArea.Min.X +
			len(ret.getEditorLinePrefix(len(p.EditorLines), len(p.EditorLines)))
	}

	return ret
}

func (p *Editor) EditorEditModeRemoveCurrentEditorLine() {
	var line *EditorLine

	if 0 == len(p.EditorLines) {
		return
	}

	for _, line = range p.EditorLines[p.CurrentEditorLineIndex:] {
		line.Index--
	}

	line = p.CurrentEditorLine()

	p.EditorLines = append(p.EditorLines[:p.CurrentEditorLineIndex], p.EditorLines[p.CurrentEditorLineIndex+1:]...)
	if p.CurrentEditorLineIndex > 0 {
		p.CurrentEditorLineIndex--
	}

	p.EditorEditModeOffXCellIndex = len(p.CurrentEditorLine().Cells)
	p.CurrentEditorLine().Data = append(p.CurrentEditorLine().Data, line.Data...)

	return
}

// 获取 line 内容前缀
//
// param:
//		lineIndex			int		 目标 line 的相应 Editor.EditorLines 中的index
//		lastEditorLineIndex		int		 输出 line 的前缀需要与整个页面的 line 前缀对齐
//									 lastEditorLineIndex 为 页面中最后一条 line 的 index
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

func (p *EditorLine) Write(keyStr string) {
	off := p.Editor.EditorCursorLocation.getOffXCellIndex()
	_off := 0

	if *off >= len(p.Cells) {
		_off = len(p.Data)
		p.Data = append(p.Data, []byte(keyStr)...)
		p.Cells = append(p.Cells, termui.Cell{[]rune(keyStr)[0], p.Editor.TextFgColor, p.Editor.TextBgColor, 0, 0, _off})

	} else if 0 == *off {
		_off = 0
		p.Data = append([]byte(keyStr), p.Data...)
		p.Cells = append(p.Cells, termui.Cell{[]rune(keyStr)[0], p.Editor.TextFgColor, p.Editor.TextBgColor, 0, 0, _off})

	} else {
		newData := make([]byte, len(p.Data)+len(keyStr))
		_off = p.Cells[*off].BytesOff
		copy(newData, p.Data[:_off])
		copy(newData[_off:], []byte(keyStr))
		copy(newData[_off+len(keyStr):], p.Data[_off:])
		p.Data = newData
		p.Cells = DefaultRawTextBuilder.Build(string(p.Data), p.Editor.TextFgColor, p.Editor.TextBgColor)
	}

	*off++
}

func (p *EditorLine) Backspace() {
	off := p.Editor.EditorCursorLocation.getOffXCellIndex()

	if *off > len(p.Cells) {
		*off = len(p.Cells)
	}

	if *off == 0 && 1 == len(p.Editor.EditorLines) {
		return
	}

	if 0 == *off {
		if EDITOR_EDIT_MODE == p.Editor.Mode {
			p.Editor.EditorEditModeRemoveCurrentEditorLine()
		}

	} else if *off == len(p.Cells) {
		p.Data = p.Data[:p.Cells[*off-1].BytesOff]
		*off -= 1

	} else {
		p.Data = append(p.Data[:p.Cells[*off-1].BytesOff], p.Data[p.Cells[*off].BytesOff:]...)
		*off -= 1
	}
}

func (p *EditorLine) CleanData() {
	off := p.Editor.EditorCursorLocation.getOffXCellIndex()
	*off = 0
	p.Data = []byte{}
	p.Cells = []termui.Cell{}
}
