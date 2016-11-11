package editor

import (
	"strings"
	"sync"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

type Editor struct {
	FirstLine, LastLine, CurrentLine *Line

	LinesLocker sync.RWMutex
	Lines       []*Line

	termui.Block

	TextFgColor       termui.Attribute
	TextBgColor       termui.Attribute
	WrapLength        int // words wrap limit. Note it may not work properly with multi-width char
	DisplayLinesRange [2]int
	*CursorLocation
}

func NewEditor() *Editor {
	ret := &Editor{
		Lines:             []*Line{},
		Block:             *termui.NewBlock(),
		TextFgColor:       termui.ThemeAttr("par.text.fg"),
		TextBgColor:       termui.ThemeAttr("par.text.bg"),
		DisplayLinesRange: [2]int{0, 1},
	}
	ret.CursorLocation = NewCursorLocation(&ret.Block)
	return ret
}

func (p *Editor) Text() string {
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

func (p *Editor) WriteNewLine(line string) {
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

func (p *Editor) Write(keyStr string) {
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
			p.RemoveLine(p.CurrentLine)
		}
		return
	}

	p.CurrentLine.Write(keyStr)
}

func (p *Editor) Buffer() termui.Buffer {
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

	if true == p.CursorLocation.IsDisplay {
		if 0 == len(cs) {
			p.CursorLocation.ResetLocation()
		} else {
			finalX = p.InnerArea.Min.X + x
			finalY = p.InnerArea.Min.Y + y
			p.CursorLocation.SetCursor(finalX, finalY)
		}
	}

	return buf
}

func (p *Editor) ActiveMode() {
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.InitLocationIfNeeded()
	p.CursorLocation.ResetCursor()
}

func (p *Editor) UnActiveMode() {
	p.CursorLocation.IsDisplay = false
	termbox.SetCursor(-1, -1)
}
