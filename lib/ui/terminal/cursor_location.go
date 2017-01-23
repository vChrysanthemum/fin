package terminal

import (
	"image"
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
)

type CursorLocation struct {
	IsDisplay   bool
	Location    image.Point
	ParentBlock *termui.Block
}

func NewCursorLocation(parentBlock *termui.Block) *CursorLocation {
	ret := &CursorLocation{
		IsDisplay:   false,
		Location:    image.Point{X: -1, Y: -1},
		ParentBlock: parentBlock,
	}
	return ret
}

func (p *CursorLocation) ResetLocation() {
	p.Location.X = p.ParentBlock.InnerArea.Min.X
	p.Location.Y = p.ParentBlock.InnerArea.Min.Y
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *CursorLocation) SetCursor(x, y int) {
	p.Location.X = x
	p.Location.Y = y
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}

func (p *CursorLocation) ResumeCursor() {
	uiutils.UISetCursor(p.Location.X, p.Location.Y)
}
