// Copyright 2016 Zack Guo <gizak@icloud.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

// Par displays a paragraph.
/*
  par := termui.NewPar("Simple Text")
  par.Height = 3
  par.Width = 17
  par.BorderLabel = "Label"
*/
type Par struct {
	Block
	Text        string
	TextFgColor Attribute
	TextBgColor Attribute
	WrapLength  int // words wrap limit. Note it may not work properly with multi-width char
}

// NewPar returns a new *Par with given text as its content.
func NewPar(s string) *Par {
	return &Par{
		Block:       *NewBlock(),
		Text:        s,
		TextFgColor: ThemeAttr("par.text.fg"),
		TextBgColor: ThemeAttr("par.text.bg"),
		WrapLength:  0,
	}
}

// Buffer implements Bufferer interface.
func (p *Par) Buffer() Buffer {
	buf := p.Block.Buffer()

	fg, bg := p.TextFgColor, p.TextBgColor
	cs := DefaultTxBuilder.Build(p.Text, fg, bg)

	// wrap if WrapLength set
	if p.WrapLength < 0 {
		cs = WrapTx(cs, p.Width-2)
	} else if p.WrapLength > 0 {
		cs = WrapTx(cs, p.WrapLength)
	}

	y, x, n := 0, 0, 0
	for y < p.InnerArea.Dy() && n < len(cs) {
		w := cs[n].Width()
		if cs[n].Ch == '\n' || x+w > p.InnerArea.Dx() {
			y++
			x = 0 // set x = 0
			if cs[n].Ch == '\n' {
				n++
			}

			if y >= p.InnerArea.Dy() {
				buf.Set(p.InnerArea.Min.X+p.InnerArea.Dx()-1,
					p.InnerArea.Min.Y+p.InnerArea.Dy()-1,
					Cell{Ch: '…', Fg: p.TextFgColor, Bg: p.TextBgColor})
				break
			}
			continue
		}

		buf.Set(p.InnerArea.Min.X+x, p.InnerArea.Min.Y+y, cs[n])

		n++
		x += w
	}

	return buf
}
