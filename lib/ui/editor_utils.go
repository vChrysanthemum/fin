package ui

import (
	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

func str2runes(s string) []rune {
	return []rune(s)
}

func toTmAttr(x termui.Attribute) termbox.Attribute {
	return termbox.Attribute(x)
}
