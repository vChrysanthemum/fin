package ui

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
