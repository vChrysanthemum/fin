package ui

import (
	"regexp"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

type CommandModeCommandHandler func(matchKey interface{}, inputModeCursor *EditorViewCursor)

type EditorCommandModeCommand struct {
	MatchKey interface{}
	Handler  CommandModeCommandHandler
}

func (p *EditorView) PrepareCommandMode() {
	p.CommandModeCommands = []EditorCommandModeCommand{
		{byte('i'), p.commandEnterInputModeBackward},
		{byte('a'), p.commandEnterInputModeForward},
		{byte(':'), p.commandEnterLastLineMode},
		{byte('u'), p.commandUndo},
		{"C-r", p.commandRedo},
		{"<up>", p.commandMoveUpOneStep},
		{"<down>", p.commandMoveDownOneStep},
		{"C-8", p.commandBackspace},
		{regexp.MustCompile(`[^\d]*(\d*)k$`), p.commandMoveUp},
		{regexp.MustCompile(`[^\d]*(\d*)j$`), p.commandMoveDown},
		{regexp.MustCompile(`[^\d]*(\d*)h$`), p.commandMoveLeft},
		{regexp.MustCompile(`[^\d]*(\d*)l$`), p.commandMoveRight},
		{regexp.MustCompile(`[^\d]*(\d*)dd$`), p.commandCut},
		{regexp.MustCompile(`[^\d]*(\d*)p$`), p.commandPaste},
	}
}

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
	ret := new(EditorView)
	ret.Editor = p
	ret.Block = &p.Block

	ret.Prepare()

	return ret
}

func (p *EditorView) Prepare() {
	p.Lines = []*EditorLine{}
	p.TextFgColor = termui.ThemeAttr("par.text.fg")
	p.TextBgColor = termui.ThemeAttr("par.text.bg")
	p.IsModifiable = true
	p.Mode = EditorModeNone

	p.PrepareCommandMode()
	p.PrepareInputMode()

	p.InputModeCursor = NewEditorViewCursor(p)

	p.InputModeAppendNewLine(p.InputModeCursor)

	p.ActionGroup = NewEditorActionGroup(p)

	p.isDisplayEditorLineNumber = true
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
