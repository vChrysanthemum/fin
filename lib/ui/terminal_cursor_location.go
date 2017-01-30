package ui

import (
	uiutils "fin/ui/utils"
	"image"

	"github.com/gizak/termui"
)

type TerminalCursorLocation struct {
	Location    image.Point
	ParentBlock *termui.Block
}

func NewTerminalCursorLocation(parentBlock *termui.Block) *TerminalCursorLocation {
	ret := &TerminalCursorLocation{
		Location:    image.Point{X: -1, Y: -1},
		ParentBlock: parentBlock,
	}
	return ret
}

func (p *TerminalCursorLocation) ResetLocation() {
	p.Location.X = p.ParentBlock.InnerArea.Min.X
	p.Location.Y = p.ParentBlock.InnerArea.Min.Y
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *TerminalCursorLocation) SetCursor(x, y int) {
	p.Location.X = x
	p.Location.Y = y
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *TerminalCursorLocation) ResumeCursor() {
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}
