package canvas

import (
	"sync"

	"github.com/gizak/termui"
)

type Canvas struct {
	termui.Block
	Image      [][]termui.Cell
	WrapLength int // words wrap limit. Note it may not work properly with multi-width char
	Locker     sync.RWMutex
}

// NewCanvas returns a new *Canvas with given text as its content.
func NewCanvas() *Canvas {
	return &Canvas{
		Block:      *termui.NewBlock(),
		WrapLength: 0,
	}
}

func (p *Canvas) initImage() {
	p.Block.Align()
	dx := p.Block.InnerArea.Dx()
	dy := p.Block.InnerArea.Dy()
	p.Image = make([][]termui.Cell, dy)
	for i := 0; i < dy; i++ {
		p.Image[i] = make([]termui.Cell, dx)
	}
}

func (p *Canvas) Clean() {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	p.initImage()
}

func (p *Canvas) UnSet(x, y int) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	if nil == p.Image {
		p.initImage()
	}

	if x < 0 || y < 0 || x >= p.Block.InnerArea.Dx() || y >= p.Block.InnerArea.Dy() {
		return
	} else {
		p.Image[y][x] = termui.Cell{}
	}
}

func (p *Canvas) Set(x, y int, cell *termui.Cell) {
	p.Locker.Lock()
	defer p.Locker.Unlock()
	if nil == p.Image {
		p.initImage()
	}

	if x < 0 || y < 0 || x >= p.Block.InnerArea.Dx() || y >= p.Block.InnerArea.Dy() {
		return
	} else {
		cell.X, cell.Y = x, y
		p.Image[y][x] = *cell
	}
}

// Buffer implements Bufferer interface.
func (p *Canvas) Buffer() termui.Buffer {
	buf := p.Block.Buffer()

	if nil == p.Image {
		return buf
	}

	trimItems := p.Image
	if len(trimItems) > p.InnerArea.Dy() {
		trimItems = trimItems[:p.InnerArea.Dy()]
	}
	for i, v := range trimItems {
		j := 0
		for _, vv := range v {
			w := vv.Width()
			if 0 == w {
				w = 1
			}
			buf.Set(p.InnerArea.Min.X+j, p.InnerArea.Min.Y+i, vv)
			j += w
		}
	}

	return buf
}
