package ui

type EditorActionInsert struct {
	EditorView        *EditorView
	EditorActionGroup *EditorActionGroup
	StartCellOffX     int
	StartLineIndex    int
	InsertData        []string
	DeletedData       []string
}

func (p *EditorActionGroup) NewEditorActionInsert(inputModeCursor *EditorViewCursor) *EditorActionInsert {
	ret := &EditorActionInsert{
		EditorView:        p.EditorView,
		EditorActionGroup: p,
		StartCellOffX:     inputModeCursor.CellOffX,
		StartLineIndex:    inputModeCursor.LineIndex,
	}

	return ret
}

func (p *EditorActionInsert) Apply(inputModeCursor *EditorViewCursor, param ...interface{}) {
	keyStr := param[0].(string)
	if "C-8" == keyStr {
		if len(p.InsertData) > 0 {
			p.InsertData = p.InsertData[:len(p.InsertData)-1]

		} else {
			if inputModeCursor.CellOffX == 0 && 1 == len(p.EditorView.Lines) {

			} else if inputModeCursor.CellOffX == 0 && len(p.EditorView.Lines) > 1 {
				p.DeletedData = append([]string{"<enter>"}, p.DeletedData...)

			} else if inputModeCursor.CellOffX > 0 {
				cursor := inputModeCursor
				line := cursor.Line()
				var str string
				if cursor.CellOffX >= len(line.Cells) {
					str = string(line.Data[line.Cells[cursor.CellOffX-1].BytesOff:])
				} else {
					str = string(line.Data[line.Cells[cursor.CellOffX-1].BytesOff:line.Cells[cursor.CellOffX].BytesOff])
				}
				p.DeletedData = append([]string{str}, p.DeletedData...)
			}

			p.StartCellOffX--
			if p.StartCellOffX < 0 {
				p.StartLineIndex--
				if p.StartLineIndex < 0 {
					p.StartLineIndex = 0
					p.StartCellOffX = 0
				} else {
					p.StartCellOffX = len(p.EditorView.Lines[p.StartLineIndex].Cells)
				}
			}
		}

	} else {
		p.InsertData = append(p.InsertData, keyStr)
	}

	p.EditorView.InputModeWrite(inputModeCursor, keyStr)
}

func (p *EditorActionInsert) Redo(inputModeCursor *EditorViewCursor) {
	if p.StartLineIndex > inputModeCursor.DisplayLinesBottomIndex {
		inputModeCursor.DisplayLinesTopIndex = p.StartLineIndex
		p.EditorView.RefreshInputModeBuf(inputModeCursor)
	}

	var (
		offStart, offEnd int
		n                = 0
		lineIndex        = p.StartLineIndex
		line             *EditorLine
	)

	line = p.EditorView.Lines[lineIndex]

	if 0 == len(p.EditorView.Lines[p.StartLineIndex].Cells) || 0 == p.StartCellOffX {
		offStart = 0
	} else {
		offStart = p.EditorView.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].BytesOff +
			len(string(p.EditorView.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].Ch))
	}
	offEnd = offStart

	for n < len(p.DeletedData) {
		for _, ch := range p.DeletedData[n:] {
			if "<enter>" == ch {
				p.EditorView.InputModeReduceLine(lineIndex + 1)
				line.CutAway(offStart, offEnd)

				offEnd = offStart
				line = p.EditorView.Lines[lineIndex]
				n++
				goto INSERT_DATA_NEXT_LINE

			} else {
				offEnd += len(ch)
				n++
			}
		}

		line.CutAway(offStart, offEnd)

	INSERT_DATA_NEXT_LINE:
	}

	inputModeCursor.LineIndex = p.StartLineIndex
	inputModeCursor.CellOffX = p.StartCellOffX
	for _, ch := range p.InsertData {
		if "<enter>" == ch {
			p.EditorView.InputModeAppendNewLine(inputModeCursor)
		} else {
			inputModeCursor.Line().Write(inputModeCursor.EditorCursor, ch)
		}
	}

	inputModeCursor.LineIndex = p.StartLineIndex
	inputModeCursor.CellOffX = p.StartCellOffX

	line = inputModeCursor.Line()
	if len(line.Cells) > 0 && inputModeCursor.CellOffX >= len(line.Cells) {
		inputModeCursor.CellOffX = len(line.Cells) - 1
	}

	p.EditorView.isShouldRefreshInputModeBuf = true
}

func (p *EditorActionInsert) Undo(inputModeCursor *EditorViewCursor) {
	if p.StartLineIndex > inputModeCursor.DisplayLinesBottomIndex {
		inputModeCursor.DisplayLinesTopIndex = p.StartLineIndex
		p.EditorView.RefreshInputModeBuf(inputModeCursor)
	}

	var (
		offStart, offEnd int
		n                = 0
		lineIndex        = p.StartLineIndex
		line             *EditorLine
	)

	line = p.EditorView.Lines[lineIndex]

	if 0 == len(p.EditorView.Lines[p.StartLineIndex].Cells) || 0 == p.StartCellOffX {
		offStart = 0
	} else {
		offStart = p.EditorView.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].BytesOff +
			len(string(p.EditorView.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].Ch))
	}
	offEnd = offStart

	for n < len(p.InsertData) {
		for _, ch := range p.InsertData[n:] {
			if "<enter>" == ch {
				p.EditorView.InputModeReduceLine(lineIndex + 1)
				line.CutAway(offStart, offEnd)

				offEnd = offStart
				line = p.EditorView.Lines[lineIndex]
				n++
				goto INSERT_DATA_NEXT_LINE

			} else {
				offEnd += len(ch)
				n++
			}
		}

		line.CutAway(offStart, offEnd)

	INSERT_DATA_NEXT_LINE:
	}

	inputModeCursor.LineIndex = p.StartLineIndex
	inputModeCursor.CellOffX = p.StartCellOffX
	for _, ch := range p.DeletedData {
		if "<enter>" == ch {
			p.EditorView.InputModeAppendNewLine(inputModeCursor)
		} else {
			inputModeCursor.Line().Write(inputModeCursor.EditorCursor, ch)
		}
	}

	inputModeCursor.LineIndex = p.StartLineIndex
	inputModeCursor.CellOffX = p.StartCellOffX

	line = inputModeCursor.Line()
	if len(line.Cells) > 0 && inputModeCursor.CellOffX >= len(line.Cells) {
		inputModeCursor.CellOffX = len(line.Cells) - 1
	}

	p.EditorView.isShouldRefreshInputModeBuf = true
}
