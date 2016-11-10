// Copyright 2016 Zack Guo <gizak@icloud.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"strconv"
	"strings"
)

// Gauge is a progress bar like widget.
// A simple example:
/*
  g := termui.NewGauge()
  g.Percent = 40
  g.Width = 50
  g.Height = 3
  g.BorderLabel = "Slim Gauge"
  g.BarColor = termui.ColorRed
  g.PercentColor = termui.ColorBlue
*/

const ColorUndef Attribute = Attribute(^uint16(0))

type Gauge struct {
	Block
	Percent                 int
	BarColor                Attribute
	PercentColor            Attribute
	PercentColorHighlighted Attribute
	Label                   string
	LabelAlign              Align
}

// NewGauge return a new gauge with current theme.
func NewGauge() *Gauge {
	g := &Gauge{
		Block:                   *NewBlock(),
		PercentColor:            ThemeAttr("gauge.percent.fg"),
		BarColor:                ThemeAttr("gauge.bar.bg"),
		Label:                   "{{percent}}%",
		LabelAlign:              AlignCenter,
		PercentColorHighlighted: ColorUndef,
	}

	g.Width = 12
	g.Height = 5
	return g
}

// Buffer implements Bufferer interface.
func (g *Gauge) Buffer() Buffer {
	buf := g.Block.Buffer()

	// plot bar
	w := g.Percent * g.InnerArea.Dx() / 100
	for i := 0; i < g.InnerArea.Dy(); i++ {
		for j := 0; j < w; j++ {
			c := Cell{}
			c.Ch = ' '
			c.Bg = g.BarColor
			if c.Bg == ColorDefault {
				c.Bg |= AttrReverse
			}
			buf.Set(g.InnerArea.Min.X+j, g.InnerArea.Min.Y+i, c)
		}
	}

	// plot percentage
	s := strings.Replace(g.Label, "{{percent}}", strconv.Itoa(g.Percent), -1)
	pry := g.InnerArea.Min.Y + g.InnerArea.Dy()/2
	rs := str2runes(s)
	var pos int
	switch g.LabelAlign {
	case AlignLeft:
		pos = 0

	case AlignCenter:
		pos = (g.InnerArea.Dx() - strWidth(s)) / 2

	case AlignRight:
		pos = g.InnerArea.Dx() - strWidth(s) - 1
	}
	pos += g.InnerArea.Min.X

	for i, v := range rs {
		c := Cell{
			Ch: v,
			Fg: g.PercentColor,
		}

		if w+g.InnerArea.Min.X > pos+i {
			c.Bg = g.BarColor
			if c.Bg == ColorDefault {
				c.Bg |= AttrReverse
			}

			if g.PercentColorHighlighted != ColorUndef {
				c.Fg = g.PercentColorHighlighted
			}
		} else {
			c.Bg = g.Block.Bg
		}

		buf.Set(1+pos+i, pry, c)
	}
	return buf
}
