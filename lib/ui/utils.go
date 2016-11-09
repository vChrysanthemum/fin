package ui

import (
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

func ColorToTermuiAttribute(color string) termui.Attribute {
	switch color {
	case "black":
		return termui.ColorBlack
	case "red":
		return termui.ColorRed
	case "green":
		return termui.ColorGreen
	case "yello":
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

	return termui.ColorDefault
}

func uirender(bs ...termui.Bufferer) {
	termui.Render(bs...)
}
