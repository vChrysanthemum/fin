package ui

import (
	"math"
	"unicode/utf8"

	rw "github.com/mattn/go-runewidth"
)

type TerminalLine struct {
	Data []byte
	Next *TerminalLine
	Prev *TerminalLine
}

func (p *Terminal) InitNewLine() *TerminalLine {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	ret := &TerminalLine{
		Data: make([]byte, 0),
	}
	p.Lines = append(p.Lines, ret)
	p.DisplayLinesRange[1] = len(p.Lines) - 1

	if nil == p.FirstTerminalLine {
		p.FirstTerminalLine = ret
	}

	if nil != p.LastTerminalLine {
		p.LastTerminalLine.Next = ret
		ret.Prev = p.LastTerminalLine
	}

	p.LastTerminalLine = ret

	p.DisplayLinesRange[1]++
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

func (p *Terminal) RemoveTerminalLine(line *TerminalLine) {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	p.Cursor.Line = line.Prev

	if nil != line.Prev {
		line.Prev.Next = line.Next
	}
	if nil != line.Next {
		line.Next.Prev = line.Prev
	}

	if p.FirstTerminalLine == line {
		p.FirstTerminalLine = p.FirstTerminalLine.Next
	}

	if p.LastTerminalLine == line {
		p.LastTerminalLine = p.LastTerminalLine.Prev
	}

	for k, v := range p.Lines {
		if line == v {
			p.Lines = append(p.Lines[:k], p.Lines[k+1:]...)
		}
	}
}

func (p *Terminal) ClearLines() {
	p.LinesLocker.Lock()
	defer p.LinesLocker.Unlock()

	p.FirstTerminalLine = nil
	p.LastTerminalLine = nil
	p.Lines = []*TerminalLine{}
	p.DisplayLinesRange = [2]int{0, 1}

	p.Cursor.ResetLocation()
	p.Cursor.Line = nil
}

func (p *TerminalLine) Write(ch string) {
	p.Data = append(p.Data, []byte(ch)...)
}

func (p *TerminalLine) Backspace() {
	if 0 == len(p.Data) {
		return
	}
	_, rlen := utf8.DecodeLastRune(p.Data)
	p.Data = p.Data[:len(p.Data)-rlen]
}
