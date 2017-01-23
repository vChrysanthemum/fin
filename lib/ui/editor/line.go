package editor

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/gizak/termui"
)

type Line struct {
	ContentStartX, ContentStartY int

	Editor *Editor
	Data   []byte
	Cells  []termui.Cell
	Next   *Line
	Prev   *Line
}

func (p *Editor) InitNewLine() *Line {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	ret := &Line{
		Editor:        p,
		ContentStartX: p.Block.InnerArea.Min.X,
		ContentStartY: p.Block.InnerArea.Min.Y,
		Data:          make([]byte, 0),
	}
	p.Lines = append(p.Lines, ret)

	if nil == p.FirstLine {
		p.FirstLine = ret
	}

	if nil != p.LastLine {
		p.LastLine.Next = ret
		ret.Prev = p.LastLine
	}

	p.LastLine = ret

	if true == p.isDisplayLineNumber {
		ret.ContentStartX = p.Block.InnerArea.Min.X +
			len(ret.getLinePrefix(len(p.Lines), len(p.Lines)))
	}

	return ret
}

func (p *Editor) RemoveLine(line *Line) {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	if nil != line.Prev {
		line.Prev.Next = line.Next
	}
	if nil != line.Next {
		line.Next.Prev = line.Prev
	}

	if p.FirstLine == line {
		p.FirstLine = p.FirstLine.Next
	}

	if p.LastLine == line {
		p.LastLine = p.LastLine.Prev
	}

	for k, v := range p.Lines {
		if line == v {
			p.Lines = append(p.Lines[:k], p.Lines[k+1:]...)
		}
	}

	p.CurrentLine = line.Prev
	p.CursorLocation.OffXCellIndex = len(p.CurrentLine.Cells)
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
		_off, i := 0, 0
		for ; i < off; i += 1 {
			_off += utf8.RuneLen(p.Cells[i].Ch)
		}
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
		p.Editor.RemoveLine(p)

	} else if off == len(p.Cells) {
		p.Data = p.Data[:len(p.Data)-utf8.RuneLen(p.Cells[off-1].Ch)]
		p.Editor.CursorLocation.OffXCellIndex -= 1

	} else {
		_off, i := 0, 0
		for ; i < off-1; i += 1 {
			_off += utf8.RuneLen(p.Cells[i].Ch)
		}
		p.Data = append(p.Data[:_off], p.Data[_off+utf8.RuneLen(p.Cells[off-1].Ch):]...)
		p.Editor.CursorLocation.OffXCellIndex -= 1
	}
}
