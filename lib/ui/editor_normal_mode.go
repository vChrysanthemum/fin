package ui

import (
	"regexp"
	"strconv"
)

type NormalModeCommandHandler func(editModeCursor *EditorCursor)

type EditorNormalModeCommand struct {
	MatchKey interface{}
	Handler  NormalModeCommandHandler
}

var (
	_commandMatchKeyEnterEditModeBackward = byte('i')
	_commandMatchKeyEnterEditModeForward  = byte('a')
	_commandMatchKeyEnterCommandMode      = byte(':')
	_commandMatchKeyUndo                  = byte('u')
	_commandMatchKeyRedo                  = "C-r"
	_commandMatchKeyMoveUpOneStep         = "<up>"
	_commandMatchKeyMoveDownOneStep       = "<down>"
	_commandMatchKeyBackspace             = "C-8"
	_commandMatchKeyMoveUp                = regexp.MustCompile(`[^\d]*(\d*)k$`)
	_commandMatchKeyMoveDown              = regexp.MustCompile(`[^\d]*(\d*)j$`)
	_commandMatchKeyMoveLeft              = regexp.MustCompile(`[^\d]*(\d*)h$`)
	_commandMatchKeyMoveRight             = regexp.MustCompile(`[^\d]*(\d*)l$`)
)

func (p *Editor) PrepareNormalMode() {
	p.NormalModeCommands = []EditorNormalModeCommand{
		{_commandMatchKeyEnterEditModeBackward, p.commandEnterEditModeBackward},
		{_commandMatchKeyEnterEditModeForward, p.commandEnterEditModeForward},
		{_commandMatchKeyEnterCommandMode, p.commandEnterCommandMode},
		{_commandMatchKeyUndo, p.commandUndo},
		{_commandMatchKeyRedo, p.commandRedo},
		{_commandMatchKeyMoveUpOneStep, p.commandMoveUpOneStep},
		{_commandMatchKeyMoveDownOneStep, p.commandMoveDownOneStep},
		{_commandMatchKeyBackspace, p.commandBackspace},
		{_commandMatchKeyMoveUp, p.commandMoveUp},
		{_commandMatchKeyMoveDown, p.commandMoveDown},
		{_commandMatchKeyMoveLeft, p.commandMoveLeft},
		{_commandMatchKeyMoveRight, p.commandMoveRight},
	}
}

func (p *Editor) NormalModeEnter(editModeCursor *EditorCursor) {
	p.Mode = EditorNormalMode
	p.NormalModeCommandStack = ""
	if editModeCursor.CellOffX >= len(editModeCursor.Line().Cells) {
		if 0 == len(editModeCursor.Line().Cells) {
			editModeCursor.CellOffX = 0
		} else {
			editModeCursor.CellOffX = len(editModeCursor.Line().Cells) - 1
		}
	}
}

func (p *Editor) NormalModeWrite(editModeCursor *EditorCursor, keyStr string) {
	p.NormalModeCommandStack += keyStr
	for _, cmd := range p.NormalModeCommands {
		switch cmd.MatchKey.(type) {
		case *regexp.Regexp:
			if true == cmd.MatchKey.(*regexp.Regexp).Match([]byte(p.NormalModeCommandStack)) {
				cmd.Handler(editModeCursor)
				p.NormalModeCommandStack = ""
				return
			}
		case byte:
			if p.NormalModeCommandStack[len(p.NormalModeCommandStack)-1] == cmd.MatchKey.(byte) {
				cmd.Handler(editModeCursor)
				p.NormalModeCommandStack = ""
				return
			}
		case string:
			matchkey := cmd.MatchKey.(string)
			if len(p.NormalModeCommandStack) >= len(matchkey) {
				if p.NormalModeCommandStack[len(p.NormalModeCommandStack)-len(matchkey):] == matchkey {
					cmd.Handler(editModeCursor)
					p.NormalModeCommandStack = ""
					return
				}
			}
		}
	}
}

func (p *Editor) commandEnterEditModeBackward(editModeCursor *EditorCursor) {
	p.EditModeEnter(editModeCursor)
}

func (p *Editor) commandEnterEditModeForward(editModeCursor *EditorCursor) {
	if len(editModeCursor.Line().Cells) > 0 {
		editModeCursor.CellOffX++
	}
	p.EditModeEnter(editModeCursor)
}

func (p *Editor) commandEnterCommandMode(editModeCursor *EditorCursor) {
	p.CommandModeEnter()
}

func (p *Editor) commandUndo(editModeCursor *EditorCursor) {
	p.ActionGroup.Undo(editModeCursor)
}

func (p *Editor) commandRedo(editModeCursor *EditorCursor) {
	p.ActionGroup.Redo(editModeCursor)
}

func (p *Editor) commandMoveUpOneStep(editModeCursor *EditorCursor) {
	p.NormalModeMoveCursorNRuneUp(editModeCursor, 1)
}

func (p *Editor) commandMoveDownOneStep(editModeCursor *EditorCursor) {
	p.NormalModeMoveCursorNRuneDown(editModeCursor, 1)
}

func (p *Editor) commandBackspace(editModeCursor *EditorCursor) {
	p.MoveCursorNRuneLeft(editModeCursor, editModeCursor.Line(), 1)
}

func (p *Editor) commandMoveUp(editModeCursor *EditorCursor) {
	_n := _commandMatchKeyMoveUp.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.NormalModeMoveCursorNRuneUp(editModeCursor, n)
	} else {
		p.NormalModeMoveCursorNRuneUp(editModeCursor, 1)
	}
}

func (p *Editor) commandMoveDown(editModeCursor *EditorCursor) {
	_n := _commandMatchKeyMoveDown.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.NormalModeMoveCursorNRuneDown(editModeCursor, n)
	} else {
		p.NormalModeMoveCursorNRuneDown(editModeCursor, 1)
	}
}

func (p *Editor) commandMoveLeft(editModeCursor *EditorCursor) {
	_n := _commandMatchKeyMoveLeft.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneLeft(editModeCursor, editModeCursor.Line(), n)
	} else {
		p.MoveCursorNRuneLeft(editModeCursor, editModeCursor.Line(), 1)
	}
}

func (p *Editor) commandMoveRight(editModeCursor *EditorCursor) {
	_n := _commandMatchKeyMoveRight.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorNRuneRight(editModeCursor, editModeCursor.Line(), n)
	} else {
		p.MoveCursorNRuneRight(editModeCursor, editModeCursor.Line(), 1)
	}
}
