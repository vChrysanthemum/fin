package editor

import (
	uiutils "fin/ui/utils"
	"sync"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	Mode EditorMode

	Buf *termui.Buffer

	// NormalMode
	NormalModeCommands     []NormalModeCommand
	NormalModeCommandStack string

	FirstLine, LastLine, CurrentLine *Line

	LinesLocker sync.RWMutex
	Lines       []*Line

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
			uiutils.UISetCursor(-1, -1)
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
		switch p.Mode {
		case EDITOR_MODE_NONE:
		case EDITOR_EDIT_MODE:
			p.EditModeWrite(keyStr)
		case EDITOR_NORMAL_MODE:
			p.NormalModeWrite(keyStr)
		case EDITOR_COMMAND_MODE:
			p.CommandModeWrite(keyStr)
		}
	}

	return
}

func (p *Editor) RefreshBorder() {
	if true == p.Block.Border {
		p.Block.DrawBorder(*p.Buf)
		p.Block.DrawBorderLabel(*p.Buf)
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}
}

func (p *Editor) RefreshContent() {
	if 0 == p.Block.InnerArea.Dy() {
		return
	}

	defer func() {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}()

	fg, bg := p.TextFgColor, p.TextBgColor

	var (
		finalX, finalY int
		y, x, n, w, k  int
		dx, dy         int
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

	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	dx, dy = 0, p.Block.InnerArea.Dy()
	pageLastLine = p.DisplayLinesTopIndex
	for k = p.DisplayLinesTopIndex; k < len(p.Lines); k++ {
		line = p.Lines[k]

		if y >= p.Block.InnerArea.Dy() {
			if p.CurrentLine == line {
				p.DisplayLinesTopIndex += 1
				goto REFRESH_BEGIN
			} else {
				return
			}
		}

		p.DisplayLinesBottomIndex = k
		line.Cells = DefaultRawTextBuilder.Build(string(line.Data), fg, bg)

		linePrefix = line.getLinePrefix(k, pageLastLine)
		line.ContentStartX = len(linePrefix) + p.Block.InnerArea.Min.X
		line.ContentStartY = y + p.Block.InnerArea.Min.Y
		x = 0
		for _, v := range linePrefix {
			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, termui.Cell{rune(v), fg, bg, finalX, finalY})
			//termbox.SetCell(finalX, finalY, v, toTmAttr(fg), toTmAttr(bg))
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
				if y >= p.Block.InnerArea.Dy() {
					p.DisplayLinesTopIndex += 1
					goto REFRESH_BEGIN
				}

				continue
			}

			finalX = line.ContentStartX + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, line.Cells[n])
			termbox.SetCell(finalX, finalY, line.Cells[n].Ch, toTmAttr(line.Cells[n].Fg), toTmAttr(line.Cells[n].Bg))
			line.Cells[n].X, line.Cells[n].Y = finalX, finalY

			n++
			x += w
		}

		y++
	}

	for ; y < dy; y++ {
		finalX = p.Block.InnerArea.Min.X
		finalY = p.Block.InnerArea.Min.Y + y
		p.Buf.Set(finalX, finalY, termui.Cell{'~', fg, bg, finalX, finalY})
	}
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.CurrentLine = p.InitNewLine()
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.Buf.IfNotRenderByTermUI = true
		p.RefreshBorder()
		p.RefreshContent()
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
	} else {
		p.RefreshBorder()
	}

	return *p.Buf
}

func (p *Editor) ActiveMode() {
	p.EditModeEnter()
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine)
}

func (p *Editor) UnActiveMode() {
	p.Mode = EDITOR_MODE_NONE
	p.CursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
