package editor

import (
	uiutils "fin/ui/utils"
	"sync"

	"github.com/gizak/termui"
)

type EditorMode int
type EditorModeWrite func(keyStr string)

type Editor struct {
	Mode      EditorMode
	ModeWrite EditorModeWrite

	// NormalMode
	NormalModeCommands     []NormalModeCommand
	NormalModeCommandStack string

	FirstLine, LastLine, CurrentLine *Line

	LinesLocker sync.RWMutex
	Lines       []*Line

	Buf *termui.Buffer
	termui.Block

	TextFgColor             termui.Attribute
	TextBgColor             termui.Attribute
	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int
	*CursorLocation

	isDisplayLineNumber bool
}

func NewEditor() *Editor {
	ret := &Editor{
		Lines:                []*Line{},
		Block:                *termui.NewBlock(),
		TextFgColor:          termui.ThemeAttr("par.text.fg"),
		TextBgColor:          termui.ThemeAttr("par.text.bg"),
		DisplayLinesTopIndex: 0,
	}
	ret.Mode = EDITOR_MODE_NONE
	ret.ModeWrite = nil
	ret.PrepareNormalMode()
	ret.PrepareEditMode()
	ret.PrepareCommandMode()
	ret.CursorLocation = NewCursorLocation(ret)
	ret.isDisplayLineNumber = true
	return ret
}

func (p *Editor) UpdateCurrentLineData(line string) {
	p.CurrentLine.Data = []byte(line)
}

func (p *Editor) Write(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	if 0 == len(p.Lines) {
		p.CurrentLine = p.InitNewLine()
	}

	switch keyStr {
	case "<escape>":
		if EDITOR_NORMAL_MODE == p.Mode {
			isQuitActiveMode = true
			return
		}

		if EDITOR_EDIT_MODE == p.Mode {
			p.EditModeQuit()
			p.NormalModeEnter()
			return
		}

		if EDITOR_COMMAND_MODE == p.Mode {
			p.NormalModeEnter()
			return
		}

	default:
		p.ModeWrite(keyStr)
	}

	return
}

func (p *Editor) RefreshContent() {
	if 0 == p.Block.InnerArea.Dy() {
		return
	}

	fg, bg := p.TextFgColor, p.TextBgColor

	var (
		finalX, finalY int
		y, x, n, w, k  int
		line           *Line
		pageLastLine   int
		linePrefix     string
	)

REFRESH_BEGIN:
	p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
	if p.DisplayLinesTopIndex >= len(p.Lines) {
		p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
		p.DisplayLinesTopIndex = len(p.Lines) - 1
		return
	}

	pageLastLine = p.DisplayLinesTopIndex + p.Block.InnerArea.Dy()

	buf := p.Block.Buffer()
	p.Buf = &buf
	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	for k = p.DisplayLinesTopIndex; k < len(p.Lines); k++ {
		line = p.Lines[k]

		if y >= p.Block.InnerArea.Dy() {
			if p.CurrentLine == line {
				p.DisplayLinesTopIndex += 1
				goto REFRESH_BEGIN
			} else {
				goto REFRESH_END
			}
		}

		p.DisplayLinesBottomIndex = k
		line.Cells = termui.DefaultTxBuilder.Build(string(line.Data), fg, bg)

		linePrefix = line.getLinePrefix(k, pageLastLine)
		line.ContentStartX = len(linePrefix) + p.Block.InnerArea.Min.X
		line.ContentStartY = y + p.Block.InnerArea.Min.Y
		x = 0
		for _, v := range linePrefix {
			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, termui.Cell{rune(v), fg, bg, finalX, finalY})
			x += 1
		}

		x, n = 0, 0
		for n < len(line.Cells) {
			w = line.Cells[n].Width()
			if x+w > p.Block.InnerArea.Dx() {
				x = 0
				y++
				// 输出一行未完成 且 超过内容区域
				if y >= p.Block.InnerArea.Dy() {
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

REFRESH_END:
	return
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.CurrentLine = p.InitNewLine()
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.RefreshContent()
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	}

	p.Block.DrawBorder(*p.Buf)
	p.Block.DrawBorderLabel(*p.Buf)

	return *p.Buf
}

func (p *Editor) ActiveMode() {
	p.EditModeEnter()
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
}

func (p *Editor) UnActiveMode() {
	p.Mode = EDITOR_MODE_NONE
	p.ModeWrite = nil
	p.CursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
