package ui

import (
	"regexp"
	"strconv"
)

type CommandModeCommandHandler func(inputModeCursor *EditorViewCursor)

type EditorCommandModeCommand struct {
	MatchKey interface{}
	Handler  CommandModeCommandHandler
}

var (
	_commandMatchKeyEnterInputModeBackward = byte('i')
	_commandMatchKeyEnterInputModeForward  = byte('a')
	_commandMatchKeyEnterLastLineMode      = byte(':')
	_commandMatchKeyUndo                   = byte('u')
	_commandMatchKeyRedo                   = "C-r"
	_commandMatchKeyMoveUpOneStep          = "<up>"
	_commandMatchKeyMoveDownOneStep        = "<down>"
	_commandMatchKeyBackspace              = "C-8"
	_commandMatchKeyMoveUp                 = regexp.MustCompile(`[^\d]*(\d*)k$`)
	_commandMatchKeyMoveDown               = regexp.MustCompile(`[^\d]*(\d*)j$`)
	_commandMatchKeyMoveLeft               = regexp.MustCompile(`[^\d]*(\d*)h$`)
	_commandMatchKeyMoveRight              = regexp.MustCompile(`[^\d]*(\d*)l$`)
	_commandMatchKeyCut                    = regexp.MustCompile(`[^\d]*(\d*)dd$`)
	_commandMatchKeyPaste                  = regexp.MustCompile(`[^\d]*(\d*)p$`)
)

func (p *EditorView) PrepareCommandMode() {
	p.CommandModeCommands = []EditorCommandModeCommand{
		{_commandMatchKeyEnterInputModeBackward, p.commandEnterInputModeBackward},
		{_commandMatchKeyEnterInputModeForward, p.commandEnterInputModeForward},
		{_commandMatchKeyEnterLastLineMode, p.commandEnterLastLineMode},
		{_commandMatchKeyUndo, p.commandUndo},
		{_commandMatchKeyRedo, p.commandRedo},
		{_commandMatchKeyMoveUpOneStep, p.commandMoveUpOneStep},
		{_commandMatchKeyMoveDownOneStep, p.commandMoveDownOneStep},
		{_commandMatchKeyBackspace, p.commandBackspace},
		{_commandMatchKeyMoveUp, p.commandMoveUp},
		{_commandMatchKeyMoveDown, p.commandMoveDown},
		{_commandMatchKeyMoveLeft, p.commandMoveLeft},
		{_commandMatchKeyMoveRight, p.commandMoveRight},
		{_commandMatchKeyCut, p.commandCut},
		{_commandMatchKeyPaste, p.commandPaste},
	}
}

func (p *EditorView) CommandModeEnter(inputModeCursor *EditorViewCursor) {
	p.Mode = EditorCommandMode
	p.CommandModeCommandStack = ""
	if nil == inputModeCursor.Line() {
		inputModeCursor.cellOffXVertical = 0
		inputModeCursor.CellOffX = 0

	} else {
		if inputModeCursor.CellOffX >= len(inputModeCursor.Line().Cells) {
			if 0 == len(inputModeCursor.Line().Cells) {
				inputModeCursor.CellOffX = 0
			} else {
				inputModeCursor.CellOffX = len(inputModeCursor.Line().Cells) - 1
			}
		}
	}
}

func (p *EditorView) CommandModeWrite(inputModeCursor *EditorViewCursor, keyStr string) {
	p.CommandModeCommandStack += keyStr
	for _, cmd := range p.CommandModeCommands {
		switch cmd.MatchKey.(type) {
		case *regexp.Regexp:
			if true == cmd.MatchKey.(*regexp.Regexp).Match([]byte(p.CommandModeCommandStack)) {
				cmd.Handler(inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		case byte:
			if p.CommandModeCommandStack[len(p.CommandModeCommandStack)-1] == cmd.MatchKey.(byte) {
				cmd.Handler(inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		case string:
			matchkey := cmd.MatchKey.(string)
			if len(p.CommandModeCommandStack) >= len(matchkey) &&
				p.CommandModeCommandStack[len(p.CommandModeCommandStack)-len(matchkey):] == matchkey {
				cmd.Handler(inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		}
	}
}

func (p *EditorView) commandEnterInputModeBackward(inputModeCursor *EditorViewCursor) {
	p.InputModeEnter(inputModeCursor)
}

func (p *EditorView) commandEnterInputModeForward(inputModeCursor *EditorViewCursor) {
	if len(inputModeCursor.Line().Cells) > 0 {
		inputModeCursor.CellOffX++
	}
	p.InputModeEnter(inputModeCursor)
}

func (p *EditorView) commandEnterLastLineMode(inputModeCursor *EditorViewCursor) {
	p.Editor.LastLineModeEnter()
}

func (p *EditorView) commandUndo(inputModeCursor *EditorViewCursor) {
	p.ActionGroup.Undo(inputModeCursor)
}

func (p *EditorView) commandRedo(inputModeCursor *EditorViewCursor) {
	p.ActionGroup.Redo(inputModeCursor)
}

func (p *EditorView) commandMoveUpOneStep(inputModeCursor *EditorViewCursor) {
	p.CommandModeMoveCursorUp(inputModeCursor, 1)
}

func (p *EditorView) commandMoveDownOneStep(inputModeCursor *EditorViewCursor) {
	p.CommandModeMoveCursorDown(inputModeCursor, 1)
}

func (p *EditorView) commandBackspace(inputModeCursor *EditorViewCursor) {
	p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), 1)
}

func (p *EditorView) commandMoveUp(inputModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveUp.FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CommandModeMoveCursorUp(inputModeCursor, n)
	} else {
		p.CommandModeMoveCursorUp(inputModeCursor, 1)
	}
}

func (p *EditorView) commandMoveDown(inputModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveDown.FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CommandModeMoveCursorDown(inputModeCursor, n)
	} else {
		p.CommandModeMoveCursorDown(inputModeCursor, 1)
	}
}

func (p *EditorView) commandMoveLeft(inputModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveLeft.FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), n)
	} else {
		p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), 1)
	}
}

func (p *EditorView) commandMoveRight(inputModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyMoveRight.FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorRight(inputModeCursor, inputModeCursor.Line(), n)
	} else {
		p.MoveCursorRight(inputModeCursor, inputModeCursor.Line(), 1)
	}
}

func (p *EditorView) commandCut(inputModeCursor *EditorViewCursor) {
	_n := _commandMatchKeyCut.FindSubmatch([]byte(p.CommandModeCommandStack))
	linesNum, err := strconv.Atoi(string(_n[1]))
	if nil != err {
		linesNum = 1
	}

	if linesNum <= 0 || inputModeCursor.LineIndex >= len(p.Lines) ||
		(1 == len(p.Lines) && 0 == len(p.Lines[0].Data)) {
		return
	}

	p.Editor.ClipboardLines.CleanLines()
	if linesNum+inputModeCursor.LineIndex < len(p.Lines) {
		p.Editor.ClipboardLines.CopyLines(p.Lines[inputModeCursor.LineIndex : inputModeCursor.LineIndex+linesNum])
	} else {
		p.Editor.ClipboardLines.CopyLines(p.Lines[inputModeCursor.LineIndex:])
	}

	p.ActionGroup.AppendEditorAction(p.ActionGroup.NewEditorActionRemoveLines())
	p.ActionGroup.CurrentUndoAction.Value.(EditorAction).Apply(inputModeCursor, linesNum)

	if inputModeCursor.LineIndex >= len(p.Lines) {
		inputModeCursor.LineIndex = len(p.Lines) - 1
	}
	p.MoveCursorLeftmost(inputModeCursor, inputModeCursor.Line())

	p.isShouldRefreshInputModeBuf = true
}

func (p *EditorView) commandPaste(inputModeCursor *EditorViewCursor) {
	var (
		copyNum int
		err     error
	)

	_n := _commandMatchKeyPaste.FindSubmatch([]byte(p.CommandModeCommandStack))
	if len(_n) > 1 {
		copyNum, err = strconv.Atoi(string(_n[1]))
		if nil != err {
			copyNum = 1
		}
	} else {
		copyNum = 1
	}

	p.ActionGroup.AppendEditorAction(p.ActionGroup.NewEditorActionInsertLines())
	p.ActionGroup.CurrentUndoAction.Value.(EditorAction).Apply(inputModeCursor, p.Editor.ClipboardLines.Lines, copyNum)

	inputModeCursor.LineIndex++
	p.MoveCursorLeftmost(inputModeCursor, inputModeCursor.Line())
	p.isShouldRefreshInputModeBuf = true
}
