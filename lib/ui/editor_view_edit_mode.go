package ui

import "github.com/gizak/termui"

func (p *EditorView) PrepareEditMode() {
}

func (p *EditorView) EditModeEnter(editModeCursor *EditorViewCursor) {
	if false == p.IsModifiable {
		p.Editor.CommandShowError(EditorErrNotModifiable)
		p.NormalModeEnter(editModeCursor)

	} else {
		editModeCursor.CellOffXVertical = 0
		p.Mode = EditorEditMode
	}
}

func (p *EditorView) EditModeWrite(editModeCursor *EditorViewCursor, keyStr string) {
	if "<enter>" == keyStr {
		p.EditModeAppendNewLine(editModeCursor)

	} else if "C-8" == keyStr {
		editModeCursor.Line().EditModeBackspace(editModeCursor)

	} else {
		editModeCursor.Line().Write(editModeCursor.EditorCursor, keyStr)
	}

	p.isShouldRefreshEditModeBuf = true
}

func (p *EditorView) RefreshEditModeBuf(editModeCursor *EditorViewCursor) {
	if 0 == p.EditModeBufAreaHeight {
		return
	}

	var (
		finalX, finalY     int
		y, x, n, w, k      int
		dx, dy             int
		line               *EditorLine
		pageLastEditorLine int
		linePrefix         string
		ok                 bool
		builtLinesMark     = make(map[int]bool, 0)
	)

REFRESH_BEGIN:
	for x = p.InnerArea.Min.X; x < p.InnerArea.Max.X; x++ {
		for y = p.InnerArea.Min.Y; y < p.InnerArea.Max.Y; y++ {
			p.Editor.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0})
		}
	}

	editModeCursor.DisplayLinesBottomIndex = editModeCursor.DisplayLinesTopIndex
	if editModeCursor.DisplayLinesTopIndex >= len(p.Lines) {
		editModeCursor.DisplayLinesBottomIndex = editModeCursor.DisplayLinesTopIndex
		editModeCursor.DisplayLinesTopIndex = len(p.Lines) - 1
		return
	}

	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	dx, dy = 0, p.EditModeBufAreaHeight
	pageLastEditorLine = editModeCursor.DisplayLinesTopIndex
	for k = editModeCursor.DisplayLinesTopIndex; k < len(p.Lines); k++ {
		line = p.Lines[k]
		if _, ok = builtLinesMark[k]; false == ok {
			line.Cells = DefaultRawTextBuilder.Build(string(line.Data), p.TextFgColor, p.TextBgColor)
			builtLinesMark[k] = true
		}

		if y >= p.EditModeBufAreaHeight {
			if editModeCursor.LineIndex == line.Index {
				editModeCursor.DisplayLinesTopIndex++
				goto REFRESH_BEGIN
			} else {
				return
			}
		}

		editModeCursor.DisplayLinesBottomIndex = k

		linePrefix = line.getEditorLinePrefix(k, pageLastEditorLine)
		line.ContentStartX = len(linePrefix) + p.Block.InnerArea.Min.X
		line.ContentStartY = y + p.Block.InnerArea.Min.Y
		x = 0
		for _, v := range linePrefix {
			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Editor.Buf.Set(finalX, finalY, termui.Cell{rune(v), p.TextFgColor, p.TextBgColor, finalX, finalY, 0})
			x++
		}

		dx = p.Block.InnerArea.Dx() - len(linePrefix)
		x, n = 0, 0
		for n < len(line.Cells) {
			w = line.Cells[n].Width()
			if x+w > dx {
				x = 0
				y++
				// 输出一行未完成 且 超过内容区域
				if y >= p.EditModeBufAreaHeight {
					editModeCursor.DisplayLinesTopIndex++
					goto REFRESH_BEGIN
				}

				continue
			}

			finalX = line.ContentStartX + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Editor.Buf.Set(finalX, finalY, line.Cells[n])
			line.Cells[n].X, line.Cells[n].Y = finalX, finalY

			n++
			x += w
		}

		y++
	}

	for ; y < dy; y++ {
		finalX = p.Block.InnerArea.Min.X
		finalY = p.Block.InnerArea.Min.Y + y
		p.Editor.Buf.Set(finalX, finalY, termui.Cell{'~', p.TextFgColor, p.TextBgColor, finalX, finalY, 0})
	}
}
