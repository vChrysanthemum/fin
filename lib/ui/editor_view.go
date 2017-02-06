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

	// NormalMode
	NormalModeCommands     []EditorNormalModeCommand
	NormalModeCommandStack string

	// EditMode
	isDisplayEditorLineNumber bool
	Lines                     []*EditorLine
	EditModeCursor            *EditorViewCursor
	EditModeBufAreaHeight     int
	ActionGroup               *EditorActionGroup

	isShouldRefreshEditModeBuf    bool
	isShouldRefreshCommandModeBuf bool

	IsModifiable bool
}

func NewEditorView(editor *Editor) *EditorView {
	ret := &EditorView{
		Editor:       editor,
		Block:        &editor.Block,
		Lines:        []*EditorLine{},
		TextFgColor:  termui.ThemeAttr("par.text.fg"),
		TextBgColor:  termui.ThemeAttr("par.text.bg"),
		IsModifiable: true,
	}

	ret.Mode = EditorModeNone

	ret.PrepareNormalMode()
	ret.PrepareEditMode()

	ret.EditModeCursor = NewEditorViewCursor(ret)

	ret.ActionGroup = NewEditorActionGroup(ret)

	ret.isDisplayEditorLineNumber = true

	return ret
}

func (p *EditorView) RefreshBuf() {
	if true == p.isShouldRefreshCommandModeBuf {
		p.Editor.RefreshCommandModeBuf(p.Editor.CommandModeCursor)
	}

	if true == p.isShouldRefreshEditModeBuf {
		p.RefreshEditModeBuf(p.EditModeCursor)
	}

	if true == p.isShouldRefreshCommandModeBuf || true == p.isShouldRefreshEditModeBuf {
		for point, c := range p.Editor.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}

	editModeCursor := p.EditModeCursor
	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.LineIndex = editModeCursor.DisplayLinesBottomIndex
	}

	return
}
