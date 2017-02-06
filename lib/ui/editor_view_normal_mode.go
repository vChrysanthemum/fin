package ui

import (
	"regexp"
	"strconv"
)

type NormalModeCommandHandler func(editModeCursor *EditorViewCursor)

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

func (p *EditorView) PrepareNormalMode() {
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

func (p *EditorView) NormalModeEnter(editModeCursor *EditorViewCursor) {
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

func (p *EditorView) NormalModeWrite(editModeCursor *EditorViewCursor, keyStr string) {
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

func (p *EditorView) commandEnterEditModeBackward(editModeCursor *EditorViewCursor) {
	p.EditModeEnter(editModeCursor)
}

func (p *EditorView) commandEnterEditModeForward(editModeCursor *EditorViewCursor) {
	if len(editModeCursor.Line().Cells) > 0 {
		editModeCursor.CellOffX++
	}
	p.EditModeEnter(editModeCursor)
}

func (p *EditorView) commandEnterCommandMode(editModeCursor *EditorViewCursor) {
	p.Editor.CommandModeEnter()
}

func (p *EditorView) commandUndo(editModeCursor *EditorViewCursor) {
	p.ActionGroup.Undo(editModeCursor)
}

func (p *EditorView) commandRedo(editModeCursor *EditorViewCursor) {
	p.ActionGroup.Redo(editModeCursor)
}

func (p *EditorView) commandMoveUpOneStep(editModeCursor *EditorViewCursor) {
	p.NormalModeMoveCursorUp(editModeCursor, 1)
}

func (p *EditorView) commandMoveDownOneStep(editModeCursor *EditorViewCursor) {
	p.NormalModeMoveCursorDown(editModeCursor, 1)
}

func (p *EditorView) commandBackspace(editModeCursor *EditorViewCursor) {
	p.MoveCursorLeft(editModeCursor, editModeCursor.Line(), 1)
}

func (p *EditorView) commandMoveUp(editModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveUp.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.NormalModeMoveCursorUp(editModeCursor, n)
	} else {
		p.NormalModeMoveCursorUp(editModeCursor, 1)
	}
}

func (p *EditorView) commandMoveDown(editModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveDown.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.NormalModeMoveCursorDown(editModeCursor, n)
	} else {
		p.NormalModeMoveCursorDown(editModeCursor, 1)
	}
}

func (p *EditorView) commandMoveLeft(editModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveLeft.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorLeft(editModeCursor, editModeCursor.Line(), n)
	} else {
		p.MoveCursorLeft(editModeCursor, editModeCursor.Line(), 1)
	}
}

func (p *EditorView) commandMoveRight(editModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveRight.FindSubmatch([]byte(p.NormalModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorRight(editModeCursor, editModeCursor.Line(), n)
	} else {
		p.MoveCursorRight(editModeCursor, editModeCursor.Line(), 1)
	}
}
