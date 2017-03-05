package ui

import (
	"regexp"
	"strconv"
)

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
	for _, cmd := range p.Editor.CommandModeCommands {
		switch cmd.MatchKey.(type) {
		case *regexp.Regexp:
			if true == cmd.MatchKey.(*regexp.Regexp).Match([]byte(p.CommandModeCommandStack)) {
				cmd.Handler(cmd.MatchKey, inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		case byte:
			if p.CommandModeCommandStack[len(p.CommandModeCommandStack)-1] == cmd.MatchKey.(byte) {
				cmd.Handler(cmd.MatchKey, inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		case string:
			matchkey := cmd.MatchKey.(string)
			if len(p.CommandModeCommandStack) >= len(matchkey) &&
				p.CommandModeCommandStack[len(p.CommandModeCommandStack)-len(matchkey):] == matchkey {
				cmd.Handler(cmd.MatchKey, inputModeCursor)
				p.CommandModeCommandStack = ""
				return
			}
		}
	}
}

func (p *EditorView) commandEnterInputModeBackward(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.InputModeEnter(inputModeCursor)
}

func (p *EditorView) commandEnterInputModeForward(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	if len(inputModeCursor.Line().Cells) > 0 {
		inputModeCursor.CellOffX++
	}
	p.InputModeEnter(inputModeCursor)
}

func (p *EditorView) commandEnterLastLineMode(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.Editor.LastLineModeEnter()
}

func (p *EditorView) commandUndo(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.ActionGroup.Undo(inputModeCursor)
}

func (p *EditorView) commandRedo(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.ActionGroup.Redo(inputModeCursor)
}

func (p *EditorView) commandMoveUpOneStep(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.CommandModeMoveCursorUp(inputModeCursor, 1)
}

func (p *EditorView) commandMoveDownOneStep(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.CommandModeMoveCursorDown(inputModeCursor, 1)
}

func (p *EditorView) commandBackspace(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), 1)
}

func (p *EditorView) commandMoveUp(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CommandModeMoveCursorUp(inputModeCursor, n)
	} else {
		p.CommandModeMoveCursorUp(inputModeCursor, 1)
	}
}

func (p *EditorView) commandMoveDown(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.CommandModeMoveCursorDown(inputModeCursor, n)
	} else {
		p.CommandModeMoveCursorDown(inputModeCursor, 1)
	}
}

func (p *EditorView) commandMoveLeft(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), n)
	} else {
		p.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), 1)
	}
}

func (p *EditorView) commandMoveRight(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
	n, err := strconv.Atoi(string(_n[1]))
	if nil == err {
		p.MoveCursorRight(inputModeCursor, inputModeCursor.Line(), n)
	} else {
		p.MoveCursorRight(inputModeCursor, inputModeCursor.Line(), 1)
	}
}

func (p *EditorView) commandCut(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
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

func (p *EditorView) commandPaste(matchKey interface{}, inputModeCursor *EditorViewCursor) {
	var (
		copyNum int
		err     error
	)

	_n := matchKey.(*regexp.Regexp).FindSubmatch([]byte(p.CommandModeCommandStack))
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
