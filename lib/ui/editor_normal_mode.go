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
	p.EditModeCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
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
	if p.EditModeCursorLocation.OffXCellIndex > p.EditModeCursorLocation.OffXCellIndexVertical {
		p.EditModeCursorLocation.OffXCellIndexVertical = p.EditModeCursorLocation.OffXCellIndex
	}

	_n := _commandMatchRegexpMoveTop.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneTop(p.EditModeCursorLocation, n)
	} else {
		p.MoveCursorNRuneTop(p.EditModeCursorLocation, 1)
	}

	if p.EditModeCursorLocation.OffXCellIndexVertical > p.EditModeCursorLocation.OffXCellIndex {
		if p.EditModeCursorLocation.OffXCellIndexVertical >= len(p.CurrentLine().Cells) {
			if 0 == len(p.CurrentLine().Cells) {
				p.EditModeCursorLocation.OffXCellIndex = 0
			} else {
				p.EditModeCursorLocation.OffXCellIndex = len(p.CurrentLine().Cells) - 1
			}
		} else {
			p.EditModeCursorLocation.OffXCellIndex = p.EditModeCursorLocation.OffXCellIndexVertical
		}
		p.EditModeCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
	}
}

func (p *Editor) commandMoveBottom() {
	if p.EditModeCursorLocation.OffXCellIndex > p.EditModeCursorLocation.OffXCellIndexVertical {
		p.EditModeCursorLocation.OffXCellIndexVertical = p.EditModeCursorLocation.OffXCellIndex
	}

	_n := _commandMatchRegexpMoveBottom.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneBottom(p.EditModeCursorLocation, n)
	} else {
		p.MoveCursorNRuneBottom(p.EditModeCursorLocation, 1)
	}

	if p.EditModeCursorLocation.OffXCellIndexVertical > p.EditModeCursorLocation.OffXCellIndex {
		if p.EditModeCursorLocation.OffXCellIndexVertical >= len(p.CurrentLine().Cells) {
			if 0 == len(p.CurrentLine().Cells) {
				p.EditModeCursorLocation.OffXCellIndex = 0
			} else {
				p.EditModeCursorLocation.OffXCellIndex = len(p.CurrentLine().Cells) - 1
			}
		} else {
			p.EditModeCursorLocation.OffXCellIndex = p.EditModeCursorLocation.OffXCellIndexVertical
		}
		p.EditModeCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
	}
}

func (p *Editor) commandMoveLeft() {
	_n := _commandMatchRegexpMoveLeft.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneLeft(p.EditModeCursorLocation, p.CurrentLine(), n)
	} else {
		p.MoveCursorNRuneLeft(p.EditModeCursorLocation, p.CurrentLine(), 1)
	}
}

func (p *Editor) commandMoveRight() {
	_n := _commandMatchRegexpMoveRight.FindSubmatch([]byte(p.EditorNormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneRight(p.EditModeCursorLocation, p.CurrentLine(), n)
	} else {
		p.MoveCursorNRuneRight(p.EditModeCursorLocation, p.CurrentLine(), 1)
	}
}

func (p *Editor) commandEnterEditorEditModeBackward() {
	p.EditorEditModeEnter()
}

func (p *Editor) commandEnterEditorEditModeForward() {
	if len(p.CurrentLine().Cells) > 0 {
		p.EditModeCursorLocation.OffXCellIndex += 1
	}
	p.EditorEditModeEnter()
}

func (p *Editor) commandEnterEditorCommandMode() {
	p.EditorCommandModeEnter()
}
