package ui

type EditorActionRemoveLines struct {
	EditorView        *EditorView
	EditorActionGroup *EditorActionGroup
	StartCellOffX     int
	StartLineIndex    int
	DeletedData       []*EditorLine
}

func (p *EditorActionGroup) NewEditorActionRemoveLines() *EditorActionRemoveLines {
	ret := &EditorActionRemoveLines{
		EditorView:        p.EditorView,
		EditorActionGroup: p,
	}

	return ret
}

func (p *EditorActionRemoveLines) Apply(inputModeCursor *EditorViewCursor, param ...interface{}) {
	p.StartCellOffX = inputModeCursor.CellOffX
	p.StartLineIndex = inputModeCursor.LineIndex

	linesNum := param[0].(int)
	lineIndex := inputModeCursor.LineIndex

	if lineIndex+linesNum > len(p.EditorView.Lines) {
		linesNum = len(p.EditorView.Lines) - lineIndex
	}

	p.DeletedData = make([]*EditorLine, linesNum)
	copy(p.DeletedData, p.EditorView.Lines[lineIndex:lineIndex+linesNum])

	p.EditorView.RemoveLines(lineIndex, linesNum)
}

func (p *EditorActionRemoveLines) Redo(inputModeCursor *EditorViewCursor) {
	p.EditorView.RemoveLines(p.StartLineIndex, len(p.DeletedData))
}

func (p *EditorActionRemoveLines) Undo(inputModeCursor *EditorViewCursor) {
	p.EditorView.AppendLines(p.StartLineIndex, p.DeletedData)
}
