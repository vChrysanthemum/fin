package editor

import (
	"image"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

type CursorLocation struct {
	IsDisplay   bool
	Location    image.Point
	ParentBlock *termui.Block
}

func NewCursorLocation(parentBlock *termui.Block) *CursorLocation {
	ret := &CursorLocation{
		IsDisplay:   true,
		Location:    image.Point{X: -1, Y: -1},
		ParentBlock: parentBlock,
	}
	return ret
}

func (p *CursorLocation) ResetLocation() {
	p.Location.X = p.ParentBlock.InnerArea.Min.X
	p.Location.Y = p.ParentBlock.InnerArea.Min.Y
}

func (p *CursorLocation) InitLocationIfNeeded() {
	if p.Location.X < 0 {
		p.Location.X = p.ParentBlock.InnerArea.Min.X
	}
	if p.Location.Y < 0 {
		p.Location.Y = p.ParentBlock.InnerArea.Min.Y
	}
}

func (p *CursorLocation) SetCursor(x, y int) {
	p.Location.X = x
	p.Location.Y = y
	termbox.SetCursor(p.Location.X, p.Location.Y)
}

func (p *CursorLocation) ResetCursor() {
	termbox.SetCursor(p.Location.X, p.Location.Y)
}
