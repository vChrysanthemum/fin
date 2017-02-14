package ui

import "errors"

const (
	EditorModeNone = iota
	EditorCommandMode
	EditorInputMode
	EditorLastLineMode

	EditorActionStatePrepareWrite = iota
	EditorActionStateWrite

	EditorActionTypeInsert = iota
	EditorActionTypeDelete

	EditorErrNotModifiableTypeLines = iota
	EditorErrNotModifiableTypeMark
)

var (
	EditorErrNotModifiable = errors.New("Cannot make changes, 'modifiable' is off")
)
