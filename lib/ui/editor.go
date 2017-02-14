package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	termui.Block
	Buf *termui.Buffer

	// LastLineMode
	LastLineModeBufAreaY      int
	LastLineModeBufAreaHeight int
	LastLineModeBuf           *EditorLine
	LastLineModeCursor        *EditorCommandCursor

	KeyEvents                       chan string
	KeyEventsResultIsQuitActiveMode chan bool

	Views []*EditorView
	*EditorView

	TmpLinesBufs *EditorTmpLinesBuf
}

func NewEditor() *Editor {
	ret := &Editor{
		Block:        *termui.NewBlock(),
		TmpLinesBufs: NewEditorTmpLinesBuf(),
	}

	ret.PrepareLastLineMode()

	ret.LastLineModeCursor = NewEditorCommandCursor(ret)

	ret.KeyEvents = make(chan string, 200)
	ret.KeyEventsResultIsQuitActiveMode = make(chan bool)
	ret.RegisterKeyEventHandlers()

	ret.Views = append(ret.Views, NewEditorView(ret))
	ret.EditorView = ret.Views[0]

	return ret
}

func (p *Editor) Close() {
	close(p.KeyEvents)
	close(p.KeyEventsResultIsQuitActiveMode)
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.InputModeAppendNewLine(p.InputModeCursor)
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.Buf.IfNotRenderByTermUI = true
		p.LastLineModeBufAreaY = p.Block.InnerArea.Max.Y - 1
		p.LastLineModeBufAreaHeight = 1
		p.InputModeBufAreaHeight = p.Block.InnerArea.Dy() - p.LastLineModeBufAreaHeight
		p.isShouldRefreshInputModeBuf = true
		p.isShouldRefreshLastLineModeBuf = true
	} else {
		p.isShouldRefreshInputModeBuf = true
		p.isShouldRefreshLastLineModeBuf = true
	}

	if true == p.Block.Border {
		p.Block.DrawBorder(*p.Buf)
		p.Block.DrawBorderLabel(*p.Buf)
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	p.RefreshBuf()
	p.RefreshCursorByEditorLine()
	p.RefreshBuf()

	return *p.Buf
}

func (p *Editor) RefreshCursorByEditorLine() {
	switch p.Mode {
	case EditorInputMode:
		p.InputModeCursor.RefreshCursorByEditorLine(p.InputModeCursor.Line())
	case EditorCommandMode:
		p.InputModeCursor.RefreshCursorByEditorLine(p.InputModeCursor.Line())
	case EditorLastLineMode:
		p.LastLineModeCursor.RefreshCursorByEditorLine(p.LastLineModeBuf)
	}
}

func (p *Editor) ActiveMode() {
	p.InputModeEnter(p.InputModeCursor)
}

func (p *Editor) UnActiveMode() {
	p.Mode = EditorModeNone
	utils.UISetCursor(-1, -1)
}
