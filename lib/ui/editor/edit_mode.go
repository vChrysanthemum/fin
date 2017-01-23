package editor

func (p *Editor) PrepareEditMode() {
}

func (p *Editor) EditModeQuit() {
	if p.CursorLocation.OffXCellIndex >= len(p.CurrentLine.Cells) {
		if 0 == len(p.CurrentLine.Cells) {
			p.CursorLocation.OffXCellIndex = 0
		} else {
			p.CursorLocation.OffXCellIndex = len(p.CurrentLine.Cells) - 1
		}
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	}
}

func (p *Editor) EditModeEnter() {
	p.offXCellIndexForVerticalMoveCursor = 0
	p.Mode = EDITOR_EDIT_MODE
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
}

func (p *Editor) EditModeWrite(keyStr string) {
	if "<space>" == keyStr {
		keyStr = " "
	}

	if "<tab>" == keyStr {
		keyStr = "\t"
	}

	if "<enter>" == keyStr {
		p.CurrentLine = p.InitNewLine()
		p.RefreshContent()
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
		return
	}

	if "C-8" == keyStr {
		p.CurrentLine.Backspace()
		p.RefreshContent()
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
		return
	}

	p.CurrentLine.Write(keyStr)
	p.RefreshContent()
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	return
}
