package editor

import (
	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

type CursorLocation struct {
	IsDisplay   bool
	X, Y        int
	ParentBlock *termui.Block
}

func NewCursorLocation(parentBlock *termui.Block) *CursorLocation {
	ret := &CursorLocation{
		IsDisplay:   true,
		X:           -1,
		Y:           -1,
		ParentBlock: parentBlock,
	}
	return ret
}

func (p *CursorLocation) InitLocationIfNeeded() {
	if p.X < 0 {
		p.X = p.ParentBlock.InnerArea.Min.X
	}
	if p.Y < 0 {
		p.Y = p.ParentBlock.InnerArea.Min.Y
	}
}

func (p *CursorLocation) SetCursor(x, y int) {
	p.X = x
	p.Y = y
	termbox.SetCursor(p.X, p.Y)
}

func (p *CursorLocation) ResetCursor() {
	termbox.SetCursor(p.X, p.Y)
}
