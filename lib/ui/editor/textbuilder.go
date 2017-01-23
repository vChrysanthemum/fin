package editor

import "github.com/gizak/termui"

var DefaultRawTextBuilder = NewRawTextBuilder()

type RawTextBuilder struct{}

func (p *RawTextBuilder) Build(s string, fg, bg termui.Attribute) []termui.Cell {
	rs := str2runes(s)
	cs := make([]termui.Cell, len(rs))
	for i := range cs {
		cs[i] = termui.Cell{Ch: rs[i], Fg: fg, Bg: bg}
	}

	return cs
}

func NewRawTextBuilder() RawTextBuilder {
	return RawTextBuilder{}
}
