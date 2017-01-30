package ui

import (
	uiutils "fin/ui/utils"
	"strings"
	"sync"

	"github.com/gizak/termui"
)

type Terminal struct {
	FirstTerminalLine, LastTerminalLine, CurrentLine *TerminalLine

	LinesLocker sync.RWMutex
	Lines       []*TerminalLine

	termui.Block

	TextFgColor               termui.Attribute
	TextBgColor               termui.Attribute
	WrapLength                int // words wrap limit. Note it may not work properly with multi-width char
	DisplayLinesRange [2]int
	*TerminalCursorLocation
}

func NewTerminal() *Terminal {
	ret := &Terminal{
		Lines:             []*TerminalLine{},
		Block:                     *termui.NewBlock(),
		TextFgColor:               termui.ThemeAttr("par.text.fg"),
		TextBgColor:               termui.ThemeAttr("par.text.bg"),
		DisplayLinesRange: [2]int{0, 1},
	}
	ret.TerminalCursorLocation = NewTerminalCursorLocation(&ret.Block)
	return ret
}

func (p *Terminal) Text() string {
	var printLines []string
	for k, line := range p.Lines {
		if k < p.DisplayLinesRange[0] {
			continue
		}
		if k >= p.DisplayLinesRange[1] {
			continue
		}
		printLines = append(printLines, string(line.Data))
	}
	return strings.Join(printLines, "\n")
}

func (p *Terminal) UpdateCurrentLineData(line string) {
	p.CurrentLine.Data = []byte(line)
}

func (p *Terminal) WriteNewLine(line string) {
	if 0 == len(p.Lines) {
		p.CurrentLine = p.InitNewLine()
	}

	// 如果上一行不为空，则启用新一行
	// 反之则利用上一行
	if len(p.CurrentLine.Data) > 0 {
		p.CurrentLine = p.InitNewLine()
	}

	p.CurrentLine.Data = []byte(line)
}

func (p *Terminal) Write(keyStr string) {
	if 0 == len(p.Lines) {
		p.CurrentLine = p.InitNewLine()
	}

	if "<space>" == keyStr {
		keyStr = " "
	}

	if "<tab>" == keyStr {
		keyStr = "\t"
	}

	if "<enter>" == keyStr {
		p.CurrentLine = p.InitNewLine()
		return
	}

	if "C-8" == keyStr {
		if len(p.CurrentLine.Data) > 0 {
			p.CurrentLine.Backspace()
		} else {
			p.RemoveTerminalLine(p.CurrentLine)
		}
		return
	}

	p.CurrentLine.Write(keyStr)
}

func (p *Terminal) Buffer() termui.Buffer {
	buf := p.Block.Buffer()

	fg, bg := p.TextFgColor, p.TextBgColor
	cs := termui.DefaultTxBuilder.Build(p.Text(), fg, bg)

	// wrap if WrapLength set
	if p.WrapLength < 0 {
		cs = termui.WrapTx(cs, p.Width-2)
	} else if p.WrapLength > 0 {
		cs = termui.WrapTx(cs, p.WrapLength)
	}

	finalX, finalY := 0, 0
	y, x, n, w := 0, 0, 0, 0
	for y < p.InnerArea.Dy() && n < len(cs) {
		w = cs[n].Width()
		if cs[n].Ch == '\n' || x+w > p.InnerArea.Dx() {
			y++
			x = 0 // set x = 0
			if cs[n].Ch == '\n' {
				n++
			}

			if y >= p.InnerArea.Dy() {
				buf.Set(p.InnerArea.Min.X+p.InnerArea.Dx()-1,
					p.InnerArea.Min.Y+p.InnerArea.Dy()-1,
					termui.Cell{Ch: '…', Fg: p.TextFgColor, Bg: p.TextBgColor})
				break
			}
			continue
		}

		finalX = p.InnerArea.Min.X + x
		finalY = p.InnerArea.Min.Y + y
		buf.Set(finalX, finalY, cs[n])

		n++
		x += w
	}

	if true == p.TerminalCursorLocation.IsDisplay {
		if 0 == len(cs) {
			p.TerminalCursorLocation.ResetLocation()
		} else {
			finalX = p.InnerArea.Min.X + x
			finalY = p.InnerArea.Min.Y + y
			p.TerminalCursorLocation.SetCursor(finalX, finalY)
		}
	}

	return buf
}

func (p *Terminal) ActiveMode() {
	p.TerminalCursorLocation.IsDisplay = true
	p.TerminalCursorLocation.ResumeCursor()
}

func (p *Terminal) UnActiveMode() {
	p.TerminalCursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
