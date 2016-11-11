package editor

import (
	"math"
	"unicode/utf8"

	rw "github.com/mattn/go-runewidth"
)

type Line struct {
	Data []byte
	Next *Line
	Prev *Line
}

func (p *Editor) InitNewLine() *Line {
	ret := &Line{
		Data: make([]byte, 0),
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

	p.DisplayLinesRange[1] += 1
	maxHeight := p.InnerArea.Dy()
	maxWidth := p.InnerArea.Dx()
	height := 1
	index := 0
	for i := p.DisplayLinesRange[1] - 2; i >= 0; i-- {
		height += int(math.Ceil(float64(rw.StringWidth(string(p.Lines[i].Data))) / float64(maxWidth)))
		if height > maxHeight {
			break
		}
		index = i
	}
	p.DisplayLinesRange[0] = index

	// p.CursorPosition.X = p.InnerArea.Min.X

	return ret
}

func (p *Editor) RemoveLine(line *Line) {
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

func (p *Line) Write(ch string) {
	p.Data = append(p.Data, []byte(ch)...)
}

func (p *Line) Backspace() {
	if 0 == len(p.Data) {
		return
	}
	_, rlen := utf8.DecodeLastRune(p.Data)
	p.Data = p.Data[:len(p.Data)-rlen]
}
