package editor

import (
	"fmt"
	"strconv"

	"github.com/gizak/termui"
)

type Line struct {
	ContentStartX, ContentStartY int

	Index  int
	Editor *Editor
	Data   []byte
	Cells  []termui.Cell
}

func (p *Editor) AppendNewLine() *Line {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	ret := &Line{
		Editor:        p,
		ContentStartX: p.Block.InnerArea.Min.X,
		ContentStartY: p.Block.InnerArea.Min.Y,
		Data:          make([]byte, 0),
	}

	if 0 == len(p.Lines) {
		ret.Index = 0
		p.Lines = append(p.Lines, ret)
		p.CurrentLineIndex = ret.Index

	} else if p.CurrentLineIndex == len(p.Lines)-1 {
		ret.Index = len(p.Lines)
		p.Lines = append(p.Lines, ret)
		p.CurrentLineIndex = ret.Index

	} else {
		for _, line := range p.Lines[p.CurrentLineIndex+1:] {
			line.Index += 1
		}
		ret.Index = p.CurrentLineIndex + 1

		n := len(p.Lines) + 1
		if cap(p.Lines) < n {
			_lines := make([]*Line, len(p.Lines), n)
			copy(_lines, p.Lines)
			p.Lines = _lines
		}
		p.Lines = p.Lines[:n]

		copy(p.Lines[p.CurrentLineIndex+2:], p.Lines[p.CurrentLineIndex+1:])
		copy(p.Lines[p.CurrentLineIndex+1:], []*Line{ret})
		p.CurrentLineIndex = ret.Index
	}

	if p.CurrentLineIndex > 0 {
		line := p.Lines[p.CurrentLineIndex-1]
		if p.OffXCellIndex < len(line.Cells) {
			p.CurrentLine().Data = line.Data[line.Cells[p.OffXCellIndex].BytesOff:]
			line.Data = line.Data[:line.Cells[p.OffXCellIndex].BytesOff]
			p.OffXCellIndex = 0
		}
	}

	if true == p.isDisplayLineNumber {
		ret.ContentStartX = p.Block.InnerArea.Min.X +
			len(ret.getLinePrefix(len(p.Lines), len(p.Lines)))
	}

	return ret
}

func (p *Editor) RemoveCurrentLine() {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	var line *Line

	if 0 == len(p.Lines) {
		return
	}

	for _, line = range p.Lines[p.CurrentLineIndex:] {
		line.Index--
	}

	line = p.CurrentLine()

	p.Lines = append(p.Lines[:p.CurrentLineIndex], p.Lines[p.CurrentLineIndex+1:]...)
	if p.CurrentLineIndex > 0 {
		p.CurrentLineIndex--
	}

	p.OffXCellIndex = len(p.CurrentLine().Cells)
	p.CurrentLine().Data = append(p.CurrentLine().Data, line.Data...)

	return
}

// 获取 line 内容前缀
//
// param:
//		lineIndex			int		 目标 line 的相应 Editor.Lines 中的index
//		lastLineIndex		int		 输出 line 的前缀需要与整个页面的 line 前缀对齐
//									 lastLineIndex 为 页面中最后一条 line 的 index
func (p *Line) getLinePrefix(lineIndex, lastLineIndex int) string {
	if true == p.Editor.isDisplayLineNumber {
		lineNumberWidth := len(strconv.Itoa(lastLineIndex + 1))
		if lineNumberWidth < 3 {
			lineNumberWidth = 3
		}

		return fmt.Sprintf(fmt.Sprintf("%%%ds ", lineNumberWidth), strconv.Itoa(lineIndex+1))
	}

	return ""
}

func (p *Line) Write(keyStr string) {
	off := p.Editor.CursorLocation.OffXCellIndex

	if off >= len(p.Cells) {
		p.Data = append(p.Data, []byte(keyStr)...)

	} else if 0 == off {
		p.Data = append([]byte(keyStr), p.Data...)

	} else {
		newData := make([]byte, len(p.Data)+len(keyStr))
		_off := p.Cells[off].BytesOff
		copy(newData, p.Data[:_off])
		copy(newData[_off:], []byte(keyStr))
		copy(newData[_off+len(keyStr):], p.Data[_off:])
		p.Data = newData
	}

	p.Editor.CursorLocation.OffXCellIndex++
}

func (p *Line) Backspace() {
	if p.Editor.CursorLocation.OffXCellIndex > len(p.Cells) {
		p.Editor.CursorLocation.OffXCellIndex = len(p.Cells)
	}
	off := p.Editor.CursorLocation.OffXCellIndex

	if off == 0 && 1 == len(p.Editor.Lines) {
		return
	}

	if 0 == off {
		p.Editor.RemoveCurrentLine()

	} else if off == len(p.Cells) {
		p.Data = p.Data[:p.Cells[off-1].BytesOff]
		p.Editor.CursorLocation.OffXCellIndex -= 1

	} else {
		p.Data = append(p.Data[:p.Cells[off].BytesOff], p.Data[p.Cells[off+1].BytesOff:]...)
		p.Editor.CursorLocation.OffXCellIndex -= 1
	}
}
