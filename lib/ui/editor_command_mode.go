package ui

import "github.com/gizak/termui"

func (p *Editor) PrepareCommandMode() {
	p.CommandModeBuf = p.NewLine()
}

func (p *Editor) CommandModeQuit() {
	p.CommandModeBuf.CleanData(p.CommandModeCursor)
}

func (p *Editor) CommandModeEnter() {
	p.Mode = EditorCommandMode
	p.CommandModeBuf.CleanData(p.CommandModeCursor)
	p.CommandModeWrite(p.EditModeCursor, p.CommandModeCursor, ":")
}

func (p *Editor) CommandModeWrite(editModeCursor, commandModeCursor *EditorCursor, keyStr string) {
	p.CommandModeBuf.ContentStartX = p.InnerArea.Min.X
	p.CommandModeBuf.ContentStartY = p.CommandModeBufAreaY

	if "<enter>" == keyStr {
		p.CommandModeQuit()
		p.NormalModeEnter(editModeCursor)

	} else if "C-8" == keyStr {
		p.CommandModeBuf.CommandModeBackspace(commandModeCursor)

	} else if "<left>" == keyStr {
		p.MoveCursorNRuneLeft(commandModeCursor, p.CommandModeBuf, 1)

	} else if "<right>" == keyStr {
		p.MoveCursorNRuneRight(commandModeCursor, p.CommandModeBuf, 1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.CommandModeBuf.Write(commandModeCursor, keyStr)
	}

	p.isShouldRefreshCommandModeBuf = true
}

func (p *Editor) RefreshCommandModeBuf(commandModeCursor *EditorCursor) {
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
		n++
	}
}
