package canvas

import "github.com/gizak/termui"

type Canvas struct {
	Block
	Image       [][]termui.Cell
	ItemFgColor termui.Attribute
	ItemBgColor termui.Attribute
	WrapLength  int // words wrap limit. Note it may not work properly with multi-width char
}

// NewCanvas returns a new *Canvas with given text as its content.
func NewCanvas() *Canvas {
	return &Canvas{
		Block:       *NewBlock(),
		ItemFgColor: termui.ThemeAttr("par.text.fg"),
		ItemBgColor: termui.ThemeAttr("par.text.bg"),
		WrapLength:  0,
	}
}

func (p *Canvas) Set(x, y int, cell *termui.Cell) {
	if nil == p.Image {
		p.Block.Align()

		sumY := p.Block.innerArea.Max.Y - p.Block.innerArea.Min.Y
		sumX := p.Block.innerArea.Max.X - p.Block.innerArea.Min.X
		p.Image = make([][]termui.Cell, sumY)
		for i := 0; i < sumY; i++ {
			p.Image[i] = make([]termui.Cell, sumX)
		}
	}

	p.Image[y][x] = *cell
}

// Buffer implements Bufferer interface.
func (p *Canvas) Buffer() termui.Buffer {
	buf := p.Block.Buffer()

	if nil == p.Image {
		return buf
	}

	trimItems := p.Image
	if len(trimItems) > p.innerArea.Dy() {
		trimItems = trimItems[:p.innerArea.Dy()]
	}
	for i, v := range trimItems {
		j := 0
		for _, vv := range v {
			w := vv.Width()
			if 0 == w {
				w = 1
			}
			buf.Set(p.innerArea.Min.X+j, p.innerArea.Min.Y+i, vv)
			j += w
		}
	}

	return buf
}
