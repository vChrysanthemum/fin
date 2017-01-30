package ui

import "github.com/gizak/termui"

func (p *Editor) PrepareEditorCommandMode() {
	p.EditorCommandModeBuf = p.NewEditorLine()
}

func (p *Editor) EditorCommandModeQuit() {
	p.isEditorCommandModeBufDirty = true
	p.EditorCommandModeBuf.CleanData()
}

func (p *Editor) EditorCommandModeEnter() {
	p.Mode = EDITOR_COMMAND_MODE
	p.EditorCommandModeBuf.CleanData()
	p.EditorCommandModeWrite(":")
}

func (p *Editor) EditorCommandModeWrite(keyStr string) {
	p.isEditorCommandModeBufDirty = true
	p.EditorCommandModeBuf.ContentStartX = p.InnerArea.Min.X
	p.EditorCommandModeBuf.ContentStartY = p.EditorCommandModeBufAreaY

	if "<enter>" == keyStr {
		p.EditorCommandModeQuit()
		p.EditorNormalModeEnter()

	} else if "C-8" == keyStr {
		p.EditorCommandModeBuf.Backspace()

	} else if "<left>" == keyStr {
		p.EditorCursorLocation.MoveCursorNRuneLeft(1)

	} else if "<right>" == keyStr {
		p.EditorCursorLocation.MoveCursorNRuneRight(1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.EditorCommandModeBuf.Write(keyStr)
	}

	p.isShouldRefreshEditorCommandModeBuf = true
}

func (p *Editor) RefreshEditorCommandModeBuf() {
	if false == p.isEditorCommandModeBufDirty {
		return
	}

	p.isEditorCommandModeBufDirty = false

	var x, y, n int

	maxY := p.EditorCommandModeBufAreaY + p.EditorCommandModeBufAreaHeight
	for x = p.Buf.Area.Min.X + 1; x < p.Buf.Area.Max.X-1; x++ {
		for y = p.EditorCommandModeBufAreaY; y < maxY; y++ {
			p.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0})
		}
	}

	p.EditorCommandModeBuf.Cells =
		DefaultRawTextBuilder.Build(string(p.EditorCommandModeBuf.Data), p.TextFgColor, p.TextBgColor)

	x = p.Block.InnerArea.Min.X
	y = p.EditorCommandModeBufAreaY
	n = 0
	for n < len(p.EditorCommandModeBuf.Cells) {
		p.Buf.Set(x, y, p.EditorCommandModeBuf.Cells[n])
		p.EditorCommandModeBuf.Cells[n].X, p.EditorCommandModeBuf.Cells[n].Y = x, y
		x += p.EditorCommandModeBuf.Cells[n].Width()
		n += 1
	}
}
