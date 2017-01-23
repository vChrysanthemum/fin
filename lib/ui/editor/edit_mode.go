package editor

import uiutils "fin/ui/utils"

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
		uiutils.UIRender(p.Editor)
	}
}

func (p *Editor) EditModeEnter() {
	p.Mode = EDITOR_EDIT_MODE
	p.ModeWrite = p.EditModeWrite
	if nil == p.CurrentLine {
		p.CurrentLine = p.InitNewLine()
	}
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	uiutils.UIRender(p.Editor)
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
		uiutils.UIRender(p.Editor)
		return
	}

	if "C-8" == keyStr {
		p.CurrentLine.Backspace()
		p.RefreshContent()
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
		uiutils.UIRender(p.Editor)
		return
	}

	p.CurrentLine.Write(keyStr)
	p.RefreshContent()
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	uiutils.UIRender(p.Editor)
	return
}
