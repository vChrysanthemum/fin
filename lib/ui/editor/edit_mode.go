package editor

func (p *Editor) PrepareEditMode() {
	p.isEditModeContentDirty = true
}

func (p *Editor) EditModeQuit() {
	if p.CursorLocation.EditModeOffXCellIndex >= len(p.CurrentLine().Cells) {
		if 0 == len(p.CurrentLine().Cells) {
			p.CursorLocation.EditModeOffXCellIndex = 0
		} else {
			p.CursorLocation.EditModeOffXCellIndex = len(p.CurrentLine().Cells) - 1
		}
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
	}
}

func (p *Editor) EditModeEnter() {
	p.offXCellIndexForVerticalMoveCursor = 0
	p.Mode = EDITOR_EDIT_MODE
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
}

func (p *Editor) EditModeWrite(keyStr string) {
	p.isEditModeContentDirty = true

	if "<enter>" == keyStr {
		p.EditModeAppendNewLine()

	} else if "C-8" == keyStr {
		p.CurrentLine().Backspace()

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.CurrentLine().Write(keyStr)
	}

	p.UIRender()
}
