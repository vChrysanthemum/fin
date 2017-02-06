package ui

import (
	"container/list"
	"fin/ui/utils"
)

type EditorActionGroup struct {
	*Editor
	State             int
	CurrentUndoAction *list.Element
	CurrentRedoAction *list.Element
	Actions           *list.List
}

type EditorAction interface {
	Apply(editModeCursor *EditorCursor, keyStr string)
	Redo(editModeCursor *EditorCursor)
	Undo(editModeCursor *EditorCursor)
}

func NewEditorActionGroup(editor *Editor) *EditorActionGroup {
	ret := &EditorActionGroup{
		Editor:  editor,
		State:   EditorActionStatePrepareWrite,
		Actions: list.New(),
	}

	return ret
}

func (p *EditorActionGroup) makeStatePrepareWrite() {
	p.State = EditorActionStatePrepareWrite
}

func (p *EditorActionGroup) Write(editModeCursor *EditorCursor, keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	switch keyStr {
	case "<escape>":
		p.makeStatePrepareWrite()

		switch p.Mode {
		case EditorNormalMode:
			isQuitActiveMode = true
			utils.UISetCursor(-1, -1)

		case EditorEditMode:
			p.NormalModeEnter(editModeCursor)

		case EditorCommandMode:
			p.CommandModeQuit()
			p.NormalModeEnter(editModeCursor)
		}

		p.isShouldRefreshCommandModeBuf = true

	default:
		switch p.Mode {
		case EditorModeNone:

		case EditorEditMode:
			switch keyStr {
			case "<left>":
				editModeCursor.CellOffXVertical = 0
				p.makeStatePrepareWrite()
				p.MoveCursorNRuneLeft(editModeCursor, editModeCursor.Line(), 1)

			case "<right>":
				editModeCursor.CellOffXVertical = 0
				p.makeStatePrepareWrite()
				p.MoveCursorNRuneRight(editModeCursor, editModeCursor.Line(), 1)

			case "<up>":
				p.makeStatePrepareWrite()
				p.EditModeMoveCursorNRuneUp(editModeCursor, 1)

			case "<down>":
				p.makeStatePrepareWrite()
				p.EditModeMoveCursorNRuneDown(editModeCursor, 1)

			default:
				if "<space>" == keyStr {
					keyStr = " "
				} else if "<tab>" == keyStr {
					keyStr = "\t"
				}

				if false == p.Editor.IsModifiable {
					p.CommandShowError(EditorErrNotModifiable)
					p.NormalModeEnter(editModeCursor)

				} else {
					editModeCursor.CellOffXVertical = 0

					if EditorActionStateWrite != p.State {
						p.AllocNewEditorActionInsert(editModeCursor)
						p.State = EditorActionStateWrite
					}
					p.CurrentUndoAction.Value.(EditorAction).Apply(editModeCursor, keyStr)

					p.EditModeWrite(editModeCursor, keyStr)
				}
			}

		case EditorNormalMode:
			p.makeStatePrepareWrite()
			p.NormalModeWrite(p.EditModeCursor, keyStr)

		case EditorCommandMode:
			p.makeStatePrepareWrite()
			p.CommandModeWrite(p.EditModeCursor, p.CommandModeCursor, keyStr)
		}
	}

	return
}

func (p *EditorActionGroup) Undo(editModeCursor *EditorCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentUndoAction {
		return
	}

	p.CurrentUndoAction.Value.(EditorAction).Undo(editModeCursor)
	p.CurrentRedoAction = p.CurrentUndoAction
	p.CurrentUndoAction = p.CurrentUndoAction.Prev()
	p.Editor.isShouldRefreshEditModeBuf = true

	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = editModeCursor.LineIndex
	}
}

func (p *EditorActionGroup) Redo(editModeCursor *EditorCursor) {
	if p.Actions.Len() <= 0 || nil == p.CurrentRedoAction {
		return
	}

	p.CurrentRedoAction.Value.(EditorAction).Redo(editModeCursor)
	p.CurrentUndoAction = p.CurrentRedoAction
	p.CurrentRedoAction = p.CurrentRedoAction.Next()
	p.Editor.isShouldRefreshEditModeBuf = true

	if editModeCursor.LineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = editModeCursor.LineIndex
	}
}
