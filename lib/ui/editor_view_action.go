package ui

import "container/list"

type EditorActionGroup struct {
	EditorView        *EditorView
	State             int
	CurrentUndoAction *list.Element
	CurrentRedoAction *list.Element
	Actions           *list.List
}

type EditorAction interface {
	Apply(inputModeCursor *EditorViewCursor, param ...interface{})
	Redo(inputModeCursor *EditorViewCursor)
	Undo(inputModeCursor *EditorViewCursor)
}

func NewEditorActionGroup(editorView *EditorView) *EditorActionGroup {
	ret := &EditorActionGroup{
		EditorView: editorView,
		State:      EditorActionStatePrepareWrite,
		Actions:    list.New(),
	}

	return ret
}

func (p *EditorActionGroup) makeStatePrepareWrite() {
	p.State = EditorActionStatePrepareWrite
}

func (p *EditorActionGroup) Write(inputModeCursor *EditorViewCursor, keyStr string) {
	switch p.EditorView.Mode {
	case EditorInputMode:
		switch keyStr {
		case "<left>":
			inputModeCursor.cellOffXVertical = 0
			p.makeStatePrepareWrite()
			p.EditorView.MoveCursorLeft(inputModeCursor, inputModeCursor.Line(), 1)

		case "<right>":
			inputModeCursor.cellOffXVertical = 0
			p.makeStatePrepareWrite()
			p.EditorView.MoveCursorRight(inputModeCursor, inputModeCursor.Line(), 1)

		case "<up>":
			p.makeStatePrepareWrite()
			p.EditorView.InputModeMoveCursorUp(inputModeCursor, 1)

		case "<down>":
			p.makeStatePrepareWrite()
			p.EditorView.InputModeMoveCursorDown(inputModeCursor, 1)

		default:
			if "<space>" == keyStr {
				keyStr = " "
			} else if "<tab>" == keyStr {
				keyStr = "\t"
			}

			if false == p.EditorView.IsModifiable {
				p.EditorView.Editor.CommandShowMsg(EditorErrNotModifiable.Error())
				p.EditorView.CommandModeEnter(inputModeCursor)

			} else {
				inputModeCursor.cellOffXVertical = 0

				if EditorActionStateWrite != p.State {
					p.EditorView.ActionGroup.AppendEditorAction(
						p.EditorView.ActionGroup.NewEditorActionInsert(inputModeCursor))
					p.State = EditorActionStateWrite
				}
				p.CurrentUndoAction.Value.(EditorAction).Apply(inputModeCursor, keyStr)
			}
		}

	case EditorCommandMode:
		p.EditorView.CommandModeWrite(p.EditorView.InputModeCursor, keyStr)
	}
}

func (p *EditorActionGroup) AppendEditorAction(action EditorAction) {
	if nil == p.CurrentUndoAction && p.Actions.Len() > 0 {
		p.Actions = list.New()
	}

	if nil != p.CurrentUndoAction {
		for e := p.Actions.Back(); e != p.CurrentUndoAction; e = p.Actions.Back() {
			p.Actions.Remove(e)
		}
	}
	p.CurrentUndoAction = p.Actions.PushBack(action)
	p.CurrentRedoAction = nil
}

func (p *EditorActionGroup) Undo(inputModeCursor *EditorViewCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentUndoAction {
		return
	}

	p.CurrentUndoAction.Value.(EditorAction).Undo(inputModeCursor)
	p.CurrentRedoAction = p.CurrentUndoAction
	p.CurrentUndoAction = p.CurrentUndoAction.Prev()
	p.EditorView.isShouldRefreshInputModeBuf = true

	if inputModeCursor.LineIndex > inputModeCursor.DisplayLinesBottomIndex {
		inputModeCursor.DisplayLinesTopIndex = inputModeCursor.LineIndex
	}
}

func (p *EditorActionGroup) Redo(inputModeCursor *EditorViewCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentRedoAction {
		return
	}

	p.CurrentRedoAction.Value.(EditorAction).Redo(inputModeCursor)
	p.CurrentUndoAction = p.CurrentRedoAction
	p.CurrentRedoAction = p.CurrentRedoAction.Next()
	p.EditorView.isShouldRefreshInputModeBuf = true

	if inputModeCursor.LineIndex > inputModeCursor.DisplayLinesBottomIndex {
		inputModeCursor.DisplayLinesTopIndex = inputModeCursor.LineIndex
	}
}
