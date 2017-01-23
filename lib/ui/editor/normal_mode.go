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
	_commandMatchRegexpMoveTop       = regexp.MustCompile(`[^\d]*(\d*)k$`)
	_commandMatchRegexpMoveBottom    = regexp.MustCompile(`[^\d]*(\d*)j$`)
	_commandMatchRegexpMoveLeft      = regexp.MustCompile(`[^\d]*(\d*)h$`)
	_commandMatchRegexpMoveRight     = regexp.MustCompile(`[^\d]*(\d*)l$`)
	_commandMatchRegexpEnterEditMode = regexp.MustCompile(`i$`)
)

func (p *Editor) PrepareNormalMode() {
	p.NormalModeCommands = []NormalModeCommand{
		{_commandMatchRegexpMoveTop, p.commandMoveTop},
		{_commandMatchRegexpMoveBottom, p.commandMoveBottom},
		{_commandMatchRegexpMoveLeft, p.commandMoveLeft},
		{_commandMatchRegexpMoveRight, p.commandMoveRight},
		{_commandMatchRegexpEnterEditMode, p.commandEnterEditMode},
	}
}

func (p *Editor) NormalModeEnter() {
	p.Mode = EDITOR_NORMAL_MODE
	p.ModeWrite = p.NormalModeWrite
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
}

func (p *Editor) commandMoveBottom() {
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

func (p *Editor) commandEnterEditMode() {
	p.EditModeEnter()
}
