package ui

type EditorCursor struct {
	CellOffX int
}

func NewEditorCursor() *EditorCursor {
	ret := &EditorCursor{
		CellOffX: 0,
	}
	return ret
}
