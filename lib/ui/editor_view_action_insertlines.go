package ui

type EditorActionInsertLines struct {
	EditorView        *EditorView
	EditorActionGroup *EditorActionGroup
	StartCellOffX     int
	StartLineIndex    int
	CopyNum           int
	Data              []EditorLine
}

func (p *EditorActionGroup) NewEditorActionInsertLines() *EditorActionInsertLines {
	ret := &EditorActionInsertLines{
		EditorView:        p.EditorView,
		EditorActionGroup: p,
	}

	return ret
}

func (p *EditorActionInsertLines) Apply(inputModeCursor *EditorViewCursor, param ...interface{}) {
	insertLines := param[0].([]EditorLine)
	p.CopyNum = param[1].(int)
	p.StartCellOffX = inputModeCursor.CellOffX
	p.StartLineIndex = inputModeCursor.LineIndex

	p.Data = []EditorLine{}
	for k := range insertLines {
		p.Data = append(p.Data, p.EditorView.TmpLinesBuf.Lines[k])
	}

	lineIndex := p.StartLineIndex + 1
	for i := 0; i < p.CopyNum; i++ {
		p.EditorView.InsertLines(lineIndex, p.Data)
		lineIndex += len(p.Data)
	}
}

func (p *EditorActionInsertLines) Redo(inputModeCursor *EditorViewCursor) {
	lineIndex := p.StartLineIndex + 1
	for i := 0; i < p.CopyNum; i++ {
		p.EditorView.InsertLines(lineIndex, p.Data)
		lineIndex += len(p.Data)
	}
}

func (p *EditorActionInsertLines) Undo(inputModeCursor *EditorViewCursor) {
	p.EditorView.RemoveLines(p.StartLineIndex+1, len(p.Data)*p.CopyNum)
}
