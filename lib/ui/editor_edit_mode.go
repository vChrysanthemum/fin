package ui

import "github.com/gizak/termui"

func (p *Editor) PrepareEditorEditMode() {
	p.isEditorEditModeBufDirty = true
}

func (p *Editor) EditorEditModeQuit() {
	if p.EditorCursorLocation.EditorEditModeOffXCellIndex >= len(p.CurrentLine().Cells) {
		if 0 == len(p.CurrentLine().Cells) {
			p.EditorCursorLocation.EditorEditModeOffXCellIndex = 0
		} else {
			p.EditorCursorLocation.EditorEditModeOffXCellIndex = len(p.CurrentLine().Cells) - 1
		}
		p.EditorCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
	}
}

func (p *Editor) EditorEditModeEnter() {
	p.offXCellIndexForVerticalMoveCursor = 0
	p.Mode = EDITOR_EDIT_MODE
	p.EditorCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
}

func (p *Editor) EditorEditModeWrite(keyStr string) {
	p.isEditorEditModeBufDirty = true

	if "<enter>" == keyStr {
		p.EditorEditModeAppendNewLine()

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

	p.isShouldRefreshEditorEditModeBuf = true
}

func (p *Editor) RefreshEditorEditModeBuf() {
	if 0 == p.EditorEditModeBufAreaHeight {
		return
	}

	if false == p.isEditorEditModeBufDirty {
		return
	}

	p.isEditorEditModeBufDirty = false

	var (
		finalX, finalY     int
		y, x, n, w, k      int
		dx, dy             int
		line               *EditorLine
		pageLastEditorLine int
		linePrefix         string
		ok                 bool
		builtLinesMark     map[int]bool = make(map[int]bool, 0)
	)

REFRESH_BEGIN:
	for x = p.InnerArea.Min.X; x < p.InnerArea.Max.X; x++ {
		for y = p.InnerArea.Min.Y; y < p.InnerArea.Max.Y; y++ {
			p.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0})
		}
	}

	p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
	if p.DisplayLinesTopIndex >= len(p.Lines) {
		p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
		p.DisplayLinesTopIndex = len(p.Lines) - 1
		return
	}

	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	dx, dy = 0, p.EditorEditModeBufAreaHeight
	pageLastEditorLine = p.DisplayLinesTopIndex
	for k = p.DisplayLinesTopIndex; k < len(p.Lines); k++ {
		line = p.Lines[k]
		if _, ok = builtLinesMark[k]; false == ok {
			line.Cells = DefaultRawTextBuilder.Build(string(line.Data), p.TextFgColor, p.TextBgColor)
			builtLinesMark[k] = true
		}

		if y >= p.EditorEditModeBufAreaHeight {
			if p.CurrentLineIndex == line.Index {
				p.DisplayLinesTopIndex += 1
				goto REFRESH_BEGIN
			} else {
				return
			}
		}

		p.DisplayLinesBottomIndex = k

		linePrefix = line.getEditorLinePrefix(k, pageLastEditorLine)
		line.ContentStartX = len(linePrefix) + p.Block.InnerArea.Min.X
		line.ContentStartY = y + p.Block.InnerArea.Min.Y
		x = 0
		for _, v := range linePrefix {
			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, termui.Cell{rune(v), p.TextFgColor, p.TextBgColor, finalX, finalY, 0})
			x += 1
		}

		dx = p.Block.InnerArea.Dx() - len(linePrefix)
		x, n = 0, 0
		for n < len(line.Cells) {
			w = line.Cells[n].Width()
			if x+w > dx {
				x = 0
				y++
				// 输出一行未完成 且 超过内容区域
				if y >= p.EditorEditModeBufAreaHeight {
					p.DisplayLinesTopIndex += 1
					goto REFRESH_BEGIN
				}

				continue
			}

			finalX = line.ContentStartX + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, line.Cells[n])
			line.Cells[n].X, line.Cells[n].Y = finalX, finalY

			n++
			x += w
		}

		y++
	}

	for ; y < dy; y++ {
		finalX = p.Block.InnerArea.Min.X
		finalY = p.Block.InnerArea.Min.Y + y
		p.Buf.Set(finalX, finalY, termui.Cell{'~', p.TextFgColor, p.TextBgColor, finalX, finalY, 0})
	}
}
