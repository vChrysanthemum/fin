package ui

import "github.com/gizak/termui"

func (p *Editor) PrepareEditorCommandMode() {
	p.CommandModeBuf = p.NewLine()
}

func (p *Editor) EditorCommandModeQuit() {
	p.isCommandModeBufDirty = true
	p.CommandModeBuf.CleanData(p.CommandModeCursorLocation)
}

func (p *Editor) EditorCommandModeEnter() {
	p.Mode = EDITOR_COMMAND_MODE
	p.CommandModeBuf.CleanData(p.CommandModeCursorLocation)
	p.EditorCommandModeWrite(":")
}

func (p *Editor) EditorCommandModeWrite(keyStr string) {
	p.isCommandModeBufDirty = true
	p.CommandModeBuf.ContentStartX = p.InnerArea.Min.X
	p.CommandModeBuf.ContentStartY = p.CommandModeBufAreaY

	if "<enter>" == keyStr {
		p.EditorCommandModeQuit()
		p.EditorNormalModeEnter()

	} else if "C-8" == keyStr {
		p.CommandModeBuf.Backspace(p.CommandModeCursorLocation)

	} else if "<left>" == keyStr {
		p.MoveCursorNRuneLeft(p.CommandModeCursorLocation, p.CommandModeBuf, 1)

	} else if "<right>" == keyStr {
		p.MoveCursorNRuneRight(p.CommandModeCursorLocation, p.CommandModeBuf, 1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.CommandModeBuf.Write(p.CommandModeCursorLocation, keyStr)
	}

	p.isShouldRefreshCommandModeBuf = true
}

func (p *Editor) RefreshCommandModeBuf() {
	if false == p.isCommandModeBufDirty {
		return
	}

	p.isCommandModeBufDirty = false

	var x, y, n int

	maxY := p.CommandModeBufAreaY + p.CommandModeBufAreaHeight
	for x = p.Buf.Area.Min.X + 1; x < p.Buf.Area.Max.X-1; x++ {
		for y = p.CommandModeBufAreaY; y < maxY; y++ {
			p.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0})
		}
	}

	p.CommandModeBuf.Cells =
		DefaultRawTextBuilder.Build(string(p.CommandModeBuf.Data), p.TextFgColor, p.TextBgColor)

	x = p.Block.InnerArea.Min.X
	y = p.CommandModeBufAreaY
	n = 0
	for n < len(p.CommandModeBuf.Cells) {
		p.Buf.Set(x, y, p.CommandModeBuf.Cells[n])
		p.CommandModeBuf.Cells[n].X, p.CommandModeBuf.Cells[n].Y = x, y
		x += p.CommandModeBuf.Cells[n].Width()
		n += 1
	}
}
