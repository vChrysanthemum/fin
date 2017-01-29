package editor

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	Mode EditorMode

	Buf *termui.Buffer

	isDisplayLineNumber bool

	// NormalMode
	NormalModeCommands     []NormalModeCommand
	NormalModeCommandStack string

	Lines []*Line

	CommandModeBuf *Line

	CurrentLineIndex int

	termui.Block
	EditModeBufAreaHeight    int
	CommandModeBufAreaY      int
	CommandModeBufAreaHeight int

	TextFgColor             termui.Attribute
	TextBgColor             termui.Attribute
	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int
	*CursorLocation

	isEditModeBufDirty              bool
	isCommandModeBufDirty           bool
	isShouldRefreshEditModeBuf      bool
	isShouldRefreshCommandModeBuf   bool
	KeyEvents                       chan string
	KeyEventsResultIsQuitActiveMode chan bool
}

func NewEditor() *Editor {
	ret := &Editor{
		Lines:                []*Line{},
		Block:                *termui.NewBlock(),
		TextFgColor:          termui.ThemeAttr("par.text.fg"),
		TextBgColor:          termui.ThemeAttr("par.text.bg"),
		DisplayLinesTopIndex: 0,
	}
	ret.Mode = EDITOR_MODE_NONE
	ret.PrepareNormalMode()
	ret.PrepareEditMode()
	ret.PrepareCommandMode()
	ret.CursorLocation = NewCursorLocation(ret)
	ret.isDisplayLineNumber = true
	ret.KeyEvents = make(chan string, 200)
	ret.KeyEventsResultIsQuitActiveMode = make(chan bool)
	ret.RegisterKeyEventHandlers()
	return ret
}

func (p *Editor) Close() {
	close(p.KeyEvents)
	close(p.KeyEventsResultIsQuitActiveMode)
}

func (p *Editor) CurrentLine() *Line {
	if p.CurrentLineIndex >= len(p.Lines) {
		return nil
	}
	return p.Lines[p.CurrentLineIndex]
}

func (p *Editor) UpdateCurrentLineData(line string) {
	p.CurrentLine().Data = []byte(line)
}

func (p *Editor) RefreshBuf() {
	if true == p.isShouldRefreshCommandModeBuf {
		p.RefreshCommandModeBuf()
	}

	if true == p.isShouldRefreshEditModeBuf {
		p.RefreshEditModeBuf()
	}

	if true == p.isShouldRefreshCommandModeBuf || true == p.isShouldRefreshEditModeBuf {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	return
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.EditModeAppendNewLine()
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
	p.UIRender()

	return *p.Buf
}

func (p *Editor) ActiveMode() {
	p.isEditModeBufDirty = true
	p.EditModeEnter()
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
}

func (p *Editor) UnActiveMode() {
	p.isEditModeBufDirty = true
	p.Mode = EDITOR_MODE_NONE
	p.CursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
