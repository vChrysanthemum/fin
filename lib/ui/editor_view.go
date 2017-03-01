package ui

import (
	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

type EditorView struct {
	Editor *Editor
	*termui.Block
	TextFgColor termui.Attribute
	TextBgColor termui.Attribute

	Mode EditorMode

	// CommandMode
	CommandModeCommands     []EditorCommandModeCommand
	CommandModeCommandStack string

	// InputMode
	inputModeBufAreaHeight    int
	isDisplayEditorLineNumber bool
	InputModeCursor           *EditorViewCursor
	ActionGroup               *EditorActionGroup
	Lines                     []*EditorLine

	isShouldRefreshInputModeBuf    bool
	isShouldRefreshLastLineModeBuf bool

	IsModifiable bool

	FilePath string
}

func (p *Editor) NewEditorView() *EditorView {
	ret := &EditorView{
		Editor:       p,
		Block:        &p.Block,
		Lines:        []*EditorLine{},
		TextFgColor:  termui.ThemeAttr("par.text.fg"),
		TextBgColor:  termui.ThemeAttr("par.text.bg"),
		IsModifiable: true,
	}

	ret.Mode = EditorModeNone

	ret.PrepareCommandMode()
	ret.PrepareInputMode()

	ret.InputModeCursor = NewEditorViewCursor(ret)

	ret.ActionGroup = NewEditorActionGroup(ret)

	ret.isDisplayEditorLineNumber = true

	return ret
}

func (p *EditorView) InputModeBufAreaHeight() int {
	return p.Editor.Block.InnerArea.Dy() - p.Editor.LastLineModeBufAreaHeight
}

func (p *EditorView) RefreshBuf() {
	if true == p.isShouldRefreshLastLineModeBuf {
		p.Editor.RefreshLastLineModeBuf(p.Editor.LastLineModeCursor)
	}

	if true == p.isShouldRefreshInputModeBuf {
		p.RefreshInputModeBuf(p.InputModeCursor)
	}

	if true == p.isShouldRefreshLastLineModeBuf || true == p.isShouldRefreshInputModeBuf {
		for point, c := range p.Editor.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	inputModeCursor := p.InputModeCursor
	if inputModeCursor.LineIndex > inputModeCursor.DisplayLinesBottomIndex {
		inputModeCursor.LineIndex = inputModeCursor.DisplayLinesBottomIndex
	}

	return
}
