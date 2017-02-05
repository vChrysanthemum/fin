package ui

import (
	"fin/ui/utils"
	"image"

	"github.com/gizak/termui"
)

type TerminalCursor struct {
	Line        *TerminalLine
	Location    image.Point
	ParentBlock *termui.Block
}

func NewTerminalCursor(parentBlock *termui.Block) *TerminalCursor {
	ret := &TerminalCursor{
		Location:    image.Point{X: -1, Y: -1},
		ParentBlock: parentBlock,
	}
	return ret
}

func (p *TerminalCursor) ResetLocation() {
	p.Location.X = p.ParentBlock.InnerArea.Min.X
	p.Location.Y = p.ParentBlock.InnerArea.Min.Y
	utils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *TerminalCursor) SetCursor(x, y int) {
	p.Location.X = x
	p.Location.Y = y
	utils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *TerminalCursor) ResumeCursor() {
	utils.UISetCursor(p.Location.X, p.Location.Y)
}
