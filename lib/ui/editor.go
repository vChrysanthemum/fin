package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	termui.Block
	Buf         *termui.Buffer
	TextFgColor termui.Attribute
	TextBgColor termui.Attribute

	Mode EditorMode

	// NormalMode
	NormalModeCommands     []EditorNormalModeCommand
	NormalModeCommandStack string

	// CommandMode
	CommandModeBufAreaY      int
	CommandModeBufAreaHeight int
	CommandModeBuf           *EditorLine
	CommandModeCursor        *EditorCursor

	// EditMode
	isDisplayEditorLineNumber bool
	Lines                     []*EditorLine
	EditModeCursor            *EditorCursor
	EditModeBufAreaHeight     int
	ActionGroup               *EditorActionGroup

	isShouldRefreshEditModeBuf      bool
	isShouldRefreshCommandModeBuf   bool
	KeyEvents                       chan string
	KeyEventsResultIsQuitActiveMode chan bool
}

func NewEditor() *Editor {
	ret := &Editor{
		Lines:       []*EditorLine{},
		Block:       *termui.NewBlock(),
		TextFgColor: termui.ThemeAttr("par.text.fg"),
		TextBgColor: termui.ThemeAttr("par.text.bg"),
	}

	ret.Mode = EDITOR_MODE_NONE

	ret.PrepareNormalMode()
	ret.PrepareEditMode()
	ret.PrepareCommandMode()

	ret.EditModeCursor = NewEditorCursor(ret)
	ret.CommandModeCursor = NewEditorCursor(ret)

	ret.ActionGroup = NewEditorActionGroup(ret)

	ret.isDisplayEditorLineNumber = true

	ret.KeyEvents = make(chan string, 200)
	ret.KeyEventsResultIsQuitActiveMode = make(chan bool)
	ret.RegisterKeyEventHandlers()

	return ret
}

func (p *Editor) Close() {
	close(p.KeyEvents)
	close(p.KeyEventsResultIsQuitActiveMode)
}

func (p *Editor) RefreshBuf() {
	if true == p.isShouldRefreshCommandModeBuf {
		p.RefreshCommandModeBuf(p.CommandModeCursor)
	}

	if true == p.isShouldRefreshEditModeBuf {
		p.RefreshEditModeBuf(p.EditModeCursor)
	}

	if true == p.isShouldRefreshCommandModeBuf || true == p.isShouldRefreshEditModeBuf {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	editModeCursor := p.EditModeCursor
	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.LineIndex = editModeCursor.DisplayLinesBottomIndex
	}

	return
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.EditModeAppendNewLine(p.EditModeCursor)
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.Buf.IfNotRenderByTermUI = true
		p.CommandModeBufAreaY = p.Block.InnerArea.Max.Y - 1
		p.CommandModeBufAreaHeight = 1
		p.EditModeBufAreaHeight = p.Block.InnerArea.Dy() - p.CommandModeBufAreaHeight
		p.isShouldRefreshEditModeBuf = true
		p.isShouldRefreshCommandModeBuf = true
	} else {
		p.isShouldRefreshEditModeBuf = true
		p.isShouldRefreshCommandModeBuf = true
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

func (p *Editor) ActiveMode() {
	p.EditModeEnter(p.EditModeCursor)
}

func (p *Editor) UnActiveMode() {
	p.Mode = EDITOR_MODE_NONE
	utils.UISetCursor(-1, -1)
}
