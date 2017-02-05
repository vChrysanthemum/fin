package ui

import "container/list"

type EditorActionInsert struct {
	Editor            *Editor
	EditorActionGroup *EditorActionGroup
	StartCellOffX     int
	StartLineIndex    int
	InsertData        []string
	DeletedData       []string
}

func (p *EditorActionGroup) AllocNewEditorActionInsert(editModeCursor *EditorCursor) *EditorActionInsert {
	ret := &EditorActionInsert{
		Editor:            p.Editor,
		EditorActionGroup: p,
		StartCellOffX:     editModeCursor.CellOffX,
		StartLineIndex:    editModeCursor.LineIndex,
	}

	if nil == p.CurrentUndoAction && p.Actions.Len() > 0 {
		p.Actions = list.New()
	}

	if nil != p.CurrentUndoAction {
		for e := p.Actions.Back(); e != p.CurrentUndoAction; e = p.Actions.Back() {
			p.Actions.Remove(e)
		}
	}
	p.CurrentUndoAction = p.Actions.PushBack(ret)
	p.CurrentRedoAction = nil
	return ret
}

func (p *EditorActionInsert) Apply(editModeCursor *EditorCursor, keyStr string) {
	if "C-8" == keyStr {
		if len(p.InsertData) > 0 {
			p.InsertData = p.InsertData[:len(p.InsertData)-1]

		} else {
			if editModeCursor.CellOffX == 0 && 1 == len(p.Editor.Lines) {

			} else if editModeCursor.CellOffX == 0 && len(p.Editor.Lines) > 1 {
				p.DeletedData = append([]string{"<enter>"}, p.DeletedData...)

			} else if editModeCursor.CellOffX > 0 {
				cursor := editModeCursor
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
					p.StartCellOffX = len(p.Editor.Lines[p.StartLineIndex].Cells)
				}
			}
		}

	} else {
		p.InsertData = append(p.InsertData, keyStr)
	}
}

func (p *EditorActionInsert) Redo(editModeCursor *EditorCursor) {
	if p.StartLineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = p.StartLineIndex
		p.Editor.RefreshEditModeBuf(editModeCursor)
	}

	var (
		offStart, offEnd int
		n                = 0
		lineIndex        = p.StartLineIndex
		line             *EditorLine
	)

	line = p.Editor.Lines[lineIndex]

	if 0 == len(p.Editor.Lines[p.StartLineIndex].Cells) || 0 == p.StartCellOffX {
		offStart = 0
	} else {
		offStart = p.Editor.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].BytesOff +
			len(string(p.Editor.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].Ch))
	}
	offEnd = offStart

	for n < len(p.DeletedData) {
		for _, ch := range p.DeletedData[n:] {
			if "<enter>" == ch {
				p.Editor.EditModeReduceLine(lineIndex + 1)
				line.CutAway(offStart, offEnd)

				offEnd = offStart
				line = p.Editor.Lines[lineIndex]
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

	editModeCursor.LineIndex = p.StartLineIndex
	editModeCursor.CellOffX = p.StartCellOffX
	for _, ch := range p.InsertData {
		if "<enter>" == ch {
			p.Editor.EditModeAppendNewLine(editModeCursor)
		} else {
			editModeCursor.Line().Write(editModeCursor, ch)
		}
	}

	editModeCursor.LineIndex = p.StartLineIndex
	editModeCursor.CellOffX = p.StartCellOffX

	line = editModeCursor.Line()
	if len(line.Cells) > 0 && editModeCursor.CellOffX >= len(line.Cells) {
		editModeCursor.CellOffX = len(line.Cells) - 1
	}

	p.Editor.isShouldRefreshEditModeBuf = true
}

func (p *EditorActionInsert) Undo(editModeCursor *EditorCursor) {
	if p.StartLineIndex > editModeCursor.DisplayLinesBottomIndex {
		editModeCursor.DisplayLinesTopIndex = p.StartLineIndex
		p.Editor.RefreshEditModeBuf(editModeCursor)
	}

	var (
		offStart, offEnd int
		n                = 0
		lineIndex        = p.StartLineIndex
		line             *EditorLine
	)

	line = p.Editor.Lines[lineIndex]

	if 0 == len(p.Editor.Lines[p.StartLineIndex].Cells) || 0 == p.StartCellOffX {
		offStart = 0
	} else {
		offStart = p.Editor.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].BytesOff +
			len(string(p.Editor.Lines[p.StartLineIndex].Cells[p.StartCellOffX-1].Ch))
	}
	offEnd = offStart

	for n < len(p.InsertData) {
		for _, ch := range p.InsertData[n:] {
			if "<enter>" == ch {
				p.Editor.EditModeReduceLine(lineIndex + 1)
				line.CutAway(offStart, offEnd)

				offEnd = offStart
				line = p.Editor.Lines[lineIndex]
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

	editModeCursor.LineIndex = p.StartLineIndex
	editModeCursor.CellOffX = p.StartCellOffX
	for _, ch := range p.DeletedData {
		if "<enter>" == ch {
			p.Editor.EditModeAppendNewLine(editModeCursor)
		} else {
			editModeCursor.Line().Write(editModeCursor, ch)
		}
	}

	editModeCursor.LineIndex = p.StartLineIndex
	editModeCursor.CellOffX = p.StartCellOffX

	line = editModeCursor.Line()
	if len(line.Cells) > 0 && editModeCursor.CellOffX >= len(line.Cells) {
		editModeCursor.CellOffX = len(line.Cells) - 1
	}

	p.Editor.isShouldRefreshEditModeBuf = true
}
