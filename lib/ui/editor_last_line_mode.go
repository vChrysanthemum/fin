package ui

import "github.com/gizak/termui"

func (p *Editor) PrepareLastLineMode() {
	editorView := p.NewEditorView()
	p.LastLineModeBuf = p.NewLine(editorView)
	p.PrepareLastLineModeCommand()
}

func (p *Editor) LastLineModeQuit() {
}

func (p *Editor) LastLineModeEnter() {
	p.Mode = EditorLastLineMode
	p.LastLineModeBuf.CleanData(p.LastLineModeCursor.EditorCursor)
	p.LastLineModeWrite(p.InputModeCursor, p.LastLineModeCursor, ":")
}

func (p *Editor) LastLineModeWrite(
	inputModeCursor *EditorViewCursor,
	lastLineModeCursor *EditorCommandCursor,
	keyStr string) {
	p.LastLineModeBuf.ContentStartX = p.InnerArea.Min.X
	p.LastLineModeBuf.ContentStartY = p.LastLineModeBufAreaY()

	if "<enter>" == keyStr {
		p.ExecLastLineCommand()
		p.LastLineModeQuit()
		p.CommandModeEnter(inputModeCursor)

	} else if "C-8" == keyStr {
		p.LastLineModeBuf.LastLineModeBackspace(lastLineModeCursor)

	} else if "<left>" == keyStr {
		p.MoveCommandCursorLeft(lastLineModeCursor, p.LastLineModeBuf, 1)

	} else if "<right>" == keyStr {
		p.MoveCommandCursorRight(lastLineModeCursor, p.LastLineModeBuf, 1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.LastLineModeBuf.Write(lastLineModeCursor.EditorCursor, keyStr)
	}

	p.isShouldRefreshLastLineModeBuf = true
}

func (p *Editor) RefreshLastLineModeBuf(lastLineModeCursor *EditorCommandCursor) {
	var x, y, n int

	maxY := p.LastLineModeBufAreaY() + p.LastLineModeBufAreaHeight
	for x = p.Buf.Area.Min.X + 1; x < p.Buf.Area.Max.X-1; x++ {
		for y = p.LastLineModeBufAreaY(); y < maxY; y++ {
			p.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0, 0})
		}
	}

	p.LastLineModeBuf.Cells =
		DefaultRawTextBuilder.Build(p.LastLineModeBuf.Data, p.TextFgColor, p.TextBgColor)

	x = p.Block.InnerArea.Min.X
	y = p.LastLineModeBufAreaY()
	n = 0
	for n < len(p.LastLineModeBuf.Cells) {
		p.Buf.Set(x, y, p.LastLineModeBuf.Cells[n])
		p.LastLineModeBuf.Cells[n].X, p.LastLineModeBuf.Cells[n].Y = x, y
		x += p.LastLineModeBuf.Cells[n].UIWidth
		n++
	}
}

func (p *Editor) CommandShowMsg(msg string) {
	p.isShouldRefreshLastLineModeBuf = true
	p.LastLineModeBuf.Data = []byte(msg)
}
