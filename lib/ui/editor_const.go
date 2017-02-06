package ui

import "errors"

const (
	EditorModeNone = iota
	EditorNormalMode
	EditorEditMode
	EditorCommandMode

	EditorActionStatePrepareWrite = iota
	EditorActionStateWrite

	EditorActionTypeInsert = iota
	EditorActionTypeDelete
)

var (
	EditorErrNotModifiable = errors.New("Cannot make changes, 'modifiable' is off")
)
