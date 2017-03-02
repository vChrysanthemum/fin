package ui

import (
	"unicode/utf8"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

var DefaultRawTextBuilder = NewRawTextBuilder()

type RawTextBuilder struct {
	TabWidth int
}

func (p *RawTextBuilder) Build(bs []byte, fg, bg termui.Attribute) []termui.Cell {
	rs := str2runes(string(bs))
	cs := make([]termui.Cell, len(rs))
	_off := 0
	_uiOff := 0
	for i := range cs {
		if 0 == i {
			if '\t' == rs[i] {
				cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: 0, UIWidth: p.TabWidth}
			} else {
				cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: 0, UIWidth: rw.RuneWidth(rs[i])}
			}
			continue
		}

		_uiOff += cs[i-1].UIWidth
		_off += utf8.RuneLen(cs[i-1].Ch)
		if '\t' == rs[i] {
			cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: _off, UIWidth: p.TabWidth}
		} else {
			cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg, BytesOff: _off, UIWidth: rw.RuneWidth(rs[i])}
		}

	}

	return cs
}

func NewRawTextBuilder() RawTextBuilder {
	return RawTextBuilder{
		TabWidth: 4,
	}
}
