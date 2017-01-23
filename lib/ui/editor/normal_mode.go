package editor

import (
	"regexp"
	"strconv"
)

type NormalModeCommandHandler func()

type NormalModeCommand struct {
	MatchRegexp *regexp.Regexp
	Handler     NormalModeCommandHandler
}

var (
	_commandMatchRegexpMoveTop               = regexp.MustCompile(`[^\d]*(\d*)k$`)
	_commandMatchRegexpMoveBottom            = regexp.MustCompile(`[^\d]*(\d*)j$`)
	_commandMatchRegexpMoveLeft              = regexp.MustCompile(`[^\d]*(\d*)h$`)
	_commandMatchRegexpMoveRight             = regexp.MustCompile(`[^\d]*(\d*)l$`)
	_commandMatchRegexpEnterEditModeBackward = regexp.MustCompile(`i$`)
	_commandMatchRegexpEnterEditModeForward  = regexp.MustCompile(`a$`)
)

func (p *Editor) PrepareNormalMode() {
	p.NormalModeCommands = []NormalModeCommand{
		{_commandMatchRegexpMoveTop, p.commandMoveTop},
		{_commandMatchRegexpMoveBottom, p.commandMoveBottom},
		{_commandMatchRegexpMoveLeft, p.commandMoveLeft},
		{_commandMatchRegexpMoveRight, p.commandMoveRight},
		{_commandMatchRegexpEnterEditModeBackward, p.commandEnterEditModeBackward},
		{_commandMatchRegexpEnterEditModeForward, p.commandEnterEditModeForward},
	}
}

func (p *Editor) NormalModeEnter() {
	p.Mode = EDITOR_NORMAL_MODE
}

func (p *Editor) NormalModeWrite(keyStr string) {
	p.NormalModeCommandStack += keyStr
	for _, cmd := range p.NormalModeCommands {
		if true == cmd.MatchRegexp.Match([]byte(p.NormalModeCommandStack)) {
			cmd.Handler()
			p.NormalModeCommandStack = ""
		}
	}
}

func (p *Editor) commandMoveTop() {
	if p.OffXCellIndex > p.offXCellIndexForVerticalMoveCursor {
		p.offXCellIndexForVerticalMoveCursor = p.OffXCellIndex
	}

	_n := _commandMatchRegexpMoveTop.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CursorLocation.MoveCursorNRuneTop(n)
	} else {
		p.CursorLocation.MoveCursorNRuneTop(1)
	}

	if p.offXCellIndexForVerticalMoveCursor > p.OffXCellIndex {
		if p.offXCellIndexForVerticalMoveCursor >= len(p.CurrentLine.Cells) {
			if 0 == len(p.CurrentLine.Cells) {
				p.OffXCellIndex = 0
			} else {
				p.OffXCellIndex = len(p.CurrentLine.Cells) - 1
			}
		} else {
			p.OffXCellIndex = p.offXCellIndexForVerticalMoveCursor
		}
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	}
}

func (p *Editor) commandMoveBottom() {
	if p.OffXCellIndex > p.offXCellIndexForVerticalMoveCursor {
		p.offXCellIndexForVerticalMoveCursor = p.OffXCellIndex
	}

	_n := _commandMatchRegexpMoveBottom.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CursorLocation.MoveCursorNRuneBottom(n)
	} else {
		p.CursorLocation.MoveCursorNRuneBottom(1)
	}

	if p.offXCellIndexForVerticalMoveCursor > p.OffXCellIndex {
		if p.offXCellIndexForVerticalMoveCursor >= len(p.CurrentLine.Cells) {
			if 0 == len(p.CurrentLine.Cells) {
				p.OffXCellIndex = 0
			} else {
				p.OffXCellIndex = len(p.CurrentLine.Cells) - 1
			}
		} else {
			p.OffXCellIndex = p.offXCellIndexForVerticalMoveCursor
		}
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	}
}

func (p *Editor) commandMoveLeft() {
	_n := _commandMatchRegexpMoveLeft.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CursorLocation.MoveCursorNRuneLeft(n)
	} else {
		p.CursorLocation.MoveCursorNRuneLeft(1)
	}
}

func (p *Editor) commandMoveRight() {
	_n := _commandMatchRegexpMoveRight.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CursorLocation.MoveCursorNRuneRight(n)
	} else {
		p.CursorLocation.MoveCursorNRuneRight(1)
	}
}

func (p *Editor) commandEnterEditModeBackward() {
	p.EditModeEnter()
}

func (p *Editor) commandEnterEditModeForward() {
	if p.OffXCellIndex > 0 {
		p.OffXCellIndex += 1
	}
	p.EditModeEnter()
}
