package editor

import (
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

	return ret
}

func (p *Editor) RemoveLine(line *Line) {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	p.CurrentLine = line.Prev

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
}

func (p *Editor) ClearLines() {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	p.FirstLine = nil
	p.LastLine = nil
	p.CurrentLine = nil
	p.Lines = []*Line{}
	p.CursorLocation.ResetLocation()
}

func (p *Line) Write(ch string) {
	off := p.Editor.CursorLocation.OffXCellIndex

	if off >= len(p.Data) {
		p.Data = append(p.Data, []byte(ch)...)

	} else if 0 == off {
		p.Data = append([]byte(ch), p.Data...)

	} else {
		newData := make([]byte, len(p.Data)+len(ch))
		_off, i := 0, 0
		for ; i < off; i += 1 {
			_off += utf8.RuneLen(p.Cells[i].Ch)
		}
		copy(newData, p.Data[:_off])
		copy(newData[_off:], []byte(ch))
		copy(newData[_off+len(ch):], p.Data[_off:])
		p.Data = newData
	}

	fg, bg := p.Editor.TextFgColor, p.Editor.TextBgColor
	cells := termui.DefaultTxBuilder.Build(ch, fg, bg)
	p.Editor.CursorLocation.OffXCellIndex += len(cells)
}

func (p *Line) Backspace() {
	return
	if 0 == len(p.Data) {
		return
	}
	_, rlen := utf8.DecodeLastRune(p.Data)
	p.Data = p.Data[:len(p.Data)-rlen]
}
