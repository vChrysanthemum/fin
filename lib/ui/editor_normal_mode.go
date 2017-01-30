package ui

import (
	"regexp"
	"strconv"
)

type EditorNormalModeCommandHandler func()

type EditorNormalModeCommand struct {
	MatchRegexp *regexp.Regexp
	Handler     EditorNormalModeCommandHandler
}

var (
	_commandMatchRegexpMoveTop                     = regexp.MustCompile(`[^\d]*(\d*)k$`)
	_commandMatchRegexpMoveBottom                  = regexp.MustCompile(`[^\d]*(\d*)j$`)
	_commandMatchRegexpMoveLeft                    = regexp.MustCompile(`[^\d]*(\d*)h$`)
	_commandMatchRegexpMoveRight                   = regexp.MustCompile(`[^\d]*(\d*)l$`)
	_commandMatchRegexpEnterEditorEditModeBackward = regexp.MustCompile(`i$`)
	_commandMatchRegexpEnterEditorEditModeForward  = regexp.MustCompile(`a$`)
	_commandMatchRegexpEnterEditorCommandMode      = regexp.MustCompile(`:$`)
)

func (p *Editor) PrepareEditorNormalMode() {
	p.EditorNormalModeCommands = []EditorNormalModeCommand{
		{_commandMatchRegexpMoveTop, p.commandMoveTop},
		{_commandMatchRegexpMoveBottom, p.commandMoveBottom},
		{_commandMatchRegexpMoveLeft, p.commandMoveLeft},
		{_commandMatchRegexpMoveRight, p.commandMoveRight},
		{_commandMatchRegexpEnterEditorEditModeBackward, p.commandEnterEditorEditModeBackward},
		{_commandMatchRegexpEnterEditorEditModeForward, p.commandEnterEditorEditModeForward},
		{_commandMatchRegexpEnterEditorCommandMode, p.commandEnterEditorCommandMode},
	}
}

func (p *Editor) EditorNormalModeEnter() {
	p.Mode = EDITOR_NORMAL_MODE
	p.EditorCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
}

func (p *Editor) EditorNormalModeWrite(keyStr string) {
	p.EditorNormalModeCommandStack += keyStr
	for _, cmd := range p.EditorNormalModeCommands {
		if true == cmd.MatchRegexp.Match([]byte(p.EditorNormalModeCommandStack)) {
			cmd.Handler()
			p.EditorNormalModeCommandStack = ""
		}
	}
}

func (p *Editor) commandMoveTop() {
	if p.EditorEditModeOffXCellIndex > p.offXCellIndexForVerticalMoveCursor {
		p.offXCellIndexForVerticalMoveCursor = p.EditorEditModeOffXCellIndex
	}

	_n := _commandMatchRegexpMoveTop.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.EditorCursorLocation.MoveCursorNRuneTop(n)
	} else {
		p.EditorCursorLocation.MoveCursorNRuneTop(1)
	}

	if p.offXCellIndexForVerticalMoveCursor > p.EditorEditModeOffXCellIndex {
		if p.offXCellIndexForVerticalMoveCursor >= len(p.Editor.CurrentLine().Cells) {
			if 0 == len(p.Editor.CurrentLine().Cells) {
				p.EditorEditModeOffXCellIndex = 0
			} else {
				p.EditorEditModeOffXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
			}
		} else {
			p.EditorEditModeOffXCellIndex = p.offXCellIndexForVerticalMoveCursor
		}
		p.EditorCursorLocation.RefreshCursorByEditorLine(p.Editor.CurrentLine())
	}
}

func (p *Editor) commandMoveBottom() {
	if p.EditorEditModeOffXCellIndex > p.offXCellIndexForVerticalMoveCursor {
		p.offXCellIndexForVerticalMoveCursor = p.EditorEditModeOffXCellIndex
	}

	_n := _commandMatchRegexpMoveBottom.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.EditorCursorLocation.MoveCursorNRuneBottom(n)
	} else {
		p.EditorCursorLocation.MoveCursorNRuneBottom(1)
	}

	if p.offXCellIndexForVerticalMoveCursor > p.EditorEditModeOffXCellIndex {
		if p.offXCellIndexForVerticalMoveCursor >= len(p.Editor.CurrentLine().Cells) {
			if 0 == len(p.Editor.CurrentLine().Cells) {
				p.EditorEditModeOffXCellIndex = 0
			} else {
				p.EditorEditModeOffXCellIndex = len(p.Editor.CurrentLine().Cells) - 1
			}
		} else {
			p.EditorEditModeOffXCellIndex = p.offXCellIndexForVerticalMoveCursor
		}
		p.EditorCursorLocation.RefreshCursorByEditorLine(p.Editor.CurrentLine())
	}
}

func (p *Editor) commandMoveLeft() {
	_n := _commandMatchRegexpMoveLeft.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.EditorCursorLocation.MoveCursorNRuneLeft(n)
	} else {
		p.EditorCursorLocation.MoveCursorNRuneLeft(1)
	}
}

func (p *Editor) commandMoveRight() {
	_n := _commandMatchRegexpMoveRight.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.EditorCursorLocation.MoveCursorNRuneRight(n)
	} else {
		p.EditorCursorLocation.MoveCursorNRuneRight(1)
	}
}

func (p *Editor) commandEnterEditorEditModeBackward() {
	p.EditorEditModeEnter()
}

func (p *Editor) commandEnterEditorEditModeForward() {
	if len(p.CurrentLine().Cells) > 0 {
		p.EditorEditModeOffXCellIndex += 1
	}
	p.EditorEditModeEnter()
}

func (p *Editor) commandEnterEditorCommandMode() {
	p.EditorCommandModeEnter()
}
