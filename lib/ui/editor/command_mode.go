package editor

import "github.com/gizak/termui"

func (p *Editor) PrepareCommandMode() {
	p.CommandModeBuf = p.NewLine()
}

func (p *Editor) CommandModeQuit() {
	p.isCommandModeBufDirty = true
	p.CommandModeBuf.CleanData()
}

func (p *Editor) CommandModeEnter() {
	p.Mode = EDITOR_COMMAND_MODE
	p.CommandModeBuf.CleanData()
	p.CommandModeWrite(":")
}

func (p *Editor) CommandModeWrite(keyStr string) {
	p.isCommandModeBufDirty = true
	p.CommandModeBuf.ContentStartX = p.InnerArea.Min.X
	p.CommandModeBuf.ContentStartY = p.CommandModeBufAreaY

	if "<enter>" == keyStr {
		p.CommandModeQuit()
		p.NormalModeEnter()

	} else if "C-8" == keyStr {
		p.CommandModeBuf.Backspace()

	} else if "<left>" == keyStr {
		p.CursorLocation.MoveCursorNRuneLeft(1)

	} else if "<right>" == keyStr {
		p.CursorLocation.MoveCursorNRuneRight(1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.CommandModeBuf.Write(keyStr)
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
