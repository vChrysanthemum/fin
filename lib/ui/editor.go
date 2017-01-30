package ui

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	Mode EditorMode

	Buf *termui.Buffer

	isDisplayEditorLineNumber bool

	// EditorNormalMode
	EditorNormalModeCommands     []EditorNormalModeCommand
	EditorNormalModeCommandStack string

	EditorLines []*EditorLine

	EditorCommandModeBuf *EditorLine

	CurrentEditorLineIndex int

	termui.Block
	EditorEditModeBufAreaHeight    int
	EditorCommandModeBufAreaY      int
	EditorCommandModeBufAreaHeight int

	TextFgColor                   termui.Attribute
	TextBgColor                   termui.Attribute
	DisplayEditorLinesTopIndex    int
	DisplayEditorLinesBottomIndex int
	*EditorCursorLocation

	isEditorEditModeBufDirty            bool
	isEditorCommandModeBufDirty         bool
	isShouldRefreshEditorEditModeBuf    bool
	isShouldRefreshEditorCommandModeBuf bool
	KeyEvents                           chan string
	KeyEventsResultIsQuitActiveMode     chan bool
}

func NewEditor() *Editor {
	ret := &Editor{
		EditorLines:                []*EditorLine{},
		Block:                      *termui.NewBlock(),
		TextFgColor:                termui.ThemeAttr("par.text.fg"),
		TextBgColor:                termui.ThemeAttr("par.text.bg"),
		DisplayEditorLinesTopIndex: 0,
	}
	ret.Mode = EDITOR_MODE_NONE
	ret.PrepareEditorNormalMode()
	ret.PrepareEditorEditMode()
	ret.PrepareEditorCommandMode()
	ret.EditorCursorLocation = NewEditorCursorLocation(ret)
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

func (p *Editor) CurrentEditorLine() *EditorLine {
	if p.CurrentEditorLineIndex >= len(p.EditorLines) {
		return nil
	}
	return p.EditorLines[p.CurrentEditorLineIndex]
}

func (p *Editor) UpdateCurrentEditorLineData(line string) {
	p.CurrentEditorLine().Data = []byte(line)
}

func (p *Editor) RefreshBuf() {
	if true == p.isShouldRefreshEditorCommandModeBuf {
		p.RefreshEditorCommandModeBuf()
	}

	if true == p.isShouldRefreshEditorEditModeBuf {
		p.RefreshEditorEditModeBuf()
	}

	if true == p.isShouldRefreshEditorCommandModeBuf || true == p.isShouldRefreshEditorEditModeBuf {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	return
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.EditorLines) {
			p.EditorEditModeAppendNewEditorLine()
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.Buf.IfNotRenderByTermUI = true
		p.EditorCommandModeBufAreaY = p.Block.InnerArea.Max.Y - 1
		p.EditorCommandModeBufAreaHeight = 1
		p.EditorEditModeBufAreaHeight = p.Block.InnerArea.Dy() - p.EditorCommandModeBufAreaHeight
		p.isShouldRefreshEditorEditModeBuf = true
		p.isShouldRefreshEditorCommandModeBuf = true
	} else {
		p.isShouldRefreshEditorEditModeBuf = true
		p.isShouldRefreshEditorCommandModeBuf = true
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
	p.isEditorEditModeBufDirty = true
	p.EditorEditModeEnter()
	p.EditorCursorLocation.IsDisplay = true
	p.EditorCursorLocation.RefreshCursorByEditorLine(p.CurrentEditorLine())
}

func (p *Editor) UnActiveMode() {
	p.isEditorEditModeBufDirty = true
	p.Mode = EDITOR_MODE_NONE
	p.EditorCursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
