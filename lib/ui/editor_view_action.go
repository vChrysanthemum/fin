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
	Apply(editModeCursor *EditorViewCursor, keyStr string)
	Redo(editModeCursor *EditorViewCursor)
	Undo(editModeCursor *EditorViewCursor)
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

func (p *EditorActionGroup) Write(editModeCursor *EditorViewCursor, keyStr string) {
	switch p.EditorView.Mode {
	case EditorEditMode:
		switch keyStr {
		case "<left>":
			editModeCursor.CellOffXVertical = 0
			p.makeStatePrepareWrite()
			p.EditorView.MoveCursorLeft(editModeCursor, editModeCursor.Line(), 1)

		case "<right>":
			editModeCursor.CellOffXVertical = 0
			p.makeStatePrepareWrite()
			p.EditorView.MoveCursorRight(editModeCursor, editModeCursor.Line(), 1)

		case "<up>":
			p.makeStatePrepareWrite()
			p.EditorView.EditModeMoveCursorUp(editModeCursor, 1)

		case "<down>":
			p.makeStatePrepareWrite()
			p.EditorView.EditModeMoveCursorDown(editModeCursor, 1)

		default:
			if "<space>" == keyStr {
				keyStr = " "
			} else if "<tab>" == keyStr {
				keyStr = "\t"
			}

			if false == p.EditorView.IsModifiable {
				p.EditorView.Editor.CommandShowError(EditorErrNotModifiable)
				p.EditorView.NormalModeEnter(editModeCursor)

			} else {
				editModeCursor.CellOffXVertical = 0

				if EditorActionStateWrite != p.State {
					p.AllocNewEditorActionInsert(editModeCursor)
					p.State = EditorActionStateWrite
				}
				p.CurrentUndoAction.Value.(EditorAction).Apply(editModeCursor, keyStr)

				p.EditorView.EditModeWrite(editModeCursor, keyStr)
			}
		}

	case EditorNormalMode:
		p.EditorView.NormalModeWrite(p.EditorView.EditModeCursor, keyStr)
	}
}

func (p *EditorActionGroup) Undo(editModeCursor *EditorViewCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentUndoAction {
		return
	}

	p.CurrentUndoAction.Value.(EditorAction).Undo(editModeCursor)
	p.CurrentRedoAction = p.CurrentUndoAction
	p.CurrentUndoAction = p.CurrentUndoAction.Prev()
	p.EditorView.isShouldRefreshEditModeBuf = true

	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = editModeCursor.LineIndex
	}
}

func (p *EditorActionGroup) Redo(editModeCursor *EditorViewCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentRedoAction {
		return
	}

	p.CurrentRedoAction.Value.(EditorAction).Redo(editModeCursor)
	p.CurrentUndoAction = p.CurrentRedoAction
	p.CurrentRedoAction = p.CurrentRedoAction.Next()
	p.EditorView.isShouldRefreshEditModeBuf = true

	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = editModeCursor.LineIndex
	}
}
