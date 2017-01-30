package ui

import (
	"unicode/utf8"

	"github.com/gizak/termui"
)

var DefaultRawTextBuilder = NewRawTextBuilder()

type RawTextBuilder struct{}

func (p *RawTextBuilder) Build(s string, fg, bg termui.Attribute) []termui.Cell {
	rs := str2runes(s)
	cs := make([]termui.Cell, len(rs))
	_off := 0
	for i := range cs {
		if i > 0 {
			_off += utf8.RuneLen(cs[i-1].Ch)
			cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: _off}
		} else {
			cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: 0}
		}
	}

	return cs
}

func NewRawTextBuilder() RawTextBuilder {
	return RawTextBuilder{}
}
