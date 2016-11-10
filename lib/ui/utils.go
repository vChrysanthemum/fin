package ui

import (
	"image"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

func FormatStringWithWidth(src string, width int) string {
	tail := width - rw.StringWidth(src)
	if tail > 0 {
		return src + string(make([]byte, tail))
	}
	return src
}

func ColorToTermuiAttribute(color string, defaultColor termui.Attribute) termui.Attribute {
	switch color {
	case "black":
		return termui.ColorBlack
	case "red":
		return termui.ColorRed
	case "green":
		return termui.ColorGreen
	case "yellow":
		return termui.ColorYellow
	case "blue":
		return termui.ColorBlue
	case "magenta":
		return termui.ColorMagenta
	case "cyan":
		return termui.ColorCyan
	case "white":
		return termui.ColorWhite
	}

	return defaultColor
}

type ClearScreenBuffer struct {
	buf termui.Buffer
}

func NewClearScreenBuffer() *ClearScreenBuffer {
	buf := termui.NewBuffer()
	min := image.Point{0, 0}
	max := image.Point{termui.TermWidth() - 1, termui.TermHeight() - 1}
	buf.SetArea(image.Rectangle{min, max})
	buf.Fill(' ', termui.ColorDefault, termui.ColorDefault)
	return &ClearScreenBuffer{
		buf: buf,
	}
}

func (p *ClearScreenBuffer) Buffer() termui.Buffer {
	return p.buf
}

func (p *ClearScreenBuffer) RefreshArea() {
	min := image.Point{0, 0}
	max := image.Point{termui.TermWidth() - 1, termui.TermHeight() - 1}
	p.buf.SetArea(image.Rectangle{min, max})
}

func uirender(bs ...termui.Bufferer) {
	termui.Render(bs...)
}

func uiclear() {
	termui.Render(GClearScreenBuffer)
}

func maxint(data ...int) int {
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}
