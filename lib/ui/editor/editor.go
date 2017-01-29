package editor

import (
	uiutils "fin/ui/utils"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

type EditorMode int

type Editor struct {
	Mode EditorMode

	Buf *termui.Buffer

	isDisplayLineNumber bool

	// NormalMode
	NormalModeCommands     []NormalModeCommand
	NormalModeCommandStack string

	Lines []*Line

	CommandModeContent *Line

	CurrentLineIndex int

	termui.Block
	EditModeContentAreaHeight    int
	CommandModeContentAreaY      int
	CommandModeContentAreaHeight int

	TextFgColor             termui.Attribute
	TextBgColor             termui.Attribute
	DisplayLinesTopIndex    int
	DisplayLinesBottomIndex int
	*CursorLocation

	isEditModeContentDirty    bool
	isCommandModeContentDirty bool
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

func (p *Editor) CurrentLine() *Line {
	if p.CurrentLineIndex >= len(p.Lines) {
		return nil
	}
	return p.Lines[p.CurrentLineIndex]
}

func (p *Editor) UpdateCurrentLineData(line string) {
	p.CurrentLine().Data = []byte(line)
}

func (p *Editor) Write(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	if 0 == len(p.Lines) {
		p.EditModeAppendNewLine()
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
			p.CommandModeQuit()
			p.NormalModeEnter()
			return
		}

	default:
		switch p.Mode {
		case EDITOR_MODE_NONE:
		case EDITOR_EDIT_MODE:
			p.Editor.refreshCommandModeBuf()
			p.EditModeWrite(keyStr)

		case EDITOR_NORMAL_MODE:
			p.Editor.refreshCommandModeBuf()
			p.NormalModeWrite(keyStr)

		case EDITOR_COMMAND_MODE:
			p.CommandModeWrite(keyStr)
		}
	}

	return
}

func (p *Editor) drawBorder() {
	if true == p.Block.Border {
		p.Block.DrawBorder(*p.Buf)
		p.Block.DrawBorderLabel(*p.Buf)
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}
}

func (p *Editor) refreshEditModeBuf() {
	if 0 == p.EditModeContentAreaHeight {
		return
	}

	defer func() {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}()

	p.drawBorder()

	if false == p.isEditModeContentDirty {
		return
	}

	p.isEditModeContentDirty = false

	var (
		finalX, finalY int
		y, x, n, w, k  int
		dx, dy         int
		line           *Line
		pageLastLine   int
		linePrefix     string
		ok             bool
		builtLinesMark map[int]bool = make(map[int]bool, 0)
	)

REFRESH_BEGIN:
	p.Buf.Fill(' ', uiutils.COLOR_DEFAULT, uiutils.COLOR_DEFAULT)

	p.drawBorder()

	p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
	if p.DisplayLinesTopIndex >= len(p.Lines) {
		p.DisplayLinesBottomIndex = p.DisplayLinesTopIndex
		p.DisplayLinesTopIndex = len(p.Lines) - 1
		return
	}

	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	dx, dy = 0, p.EditModeContentAreaHeight
	pageLastLine = p.DisplayLinesTopIndex
	for k = p.DisplayLinesTopIndex; k < len(p.Lines); k++ {
		line = p.Lines[k]
		if _, ok = builtLinesMark[k]; false == ok {
			line.Cells = DefaultRawTextBuilder.Build(string(line.Data), p.TextFgColor, p.TextBgColor)
			builtLinesMark[k] = true
		}

		if y >= p.EditModeContentAreaHeight {
			if p.CurrentLineIndex == line.Index {
				p.DisplayLinesTopIndex += 1
				goto REFRESH_BEGIN
			} else {
				return
			}
		}

		p.DisplayLinesBottomIndex = k

		linePrefix = line.getLinePrefix(k, pageLastLine)
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
				if y >= p.EditModeContentAreaHeight {
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

func (p *Editor) refreshCommandModeBuf() {
	if false == p.isCommandModeContentDirty {
		return
	}

	defer func() {
		for point, c := range p.Buf.CellMap {
			termbox.SetCell(point.X, point.Y, c.Ch, toTmAttr(c.Fg), toTmAttr(c.Bg))
		}
	}()

	p.isCommandModeContentDirty = false

	var x, y, n int

	maxY := p.CommandModeContentAreaY + p.CommandModeContentAreaHeight
	for x = p.Buf.Area.Min.X + 1; x < p.Buf.Area.Max.X-1; x++ {
		for y = p.CommandModeContentAreaY; y < maxY; y++ {
			p.Buf.Set(x, y, termui.Cell{' ', p.TextFgColor, p.TextBgColor, 0, 0, 0})
		}
	}

	p.CommandModeContent.Cells =
		DefaultRawTextBuilder.Build(string(p.CommandModeContent.Data), p.TextFgColor, p.TextBgColor)

	x = p.Block.InnerArea.Min.X
	y = p.CommandModeContentAreaY
	n = 0
	for n < len(p.CommandModeContent.Cells) {
		p.Buf.Set(x, y, p.CommandModeContent.Cells[n])
		p.CommandModeContent.Cells[n].X, p.CommandModeContent.Cells[n].Y = x, y
		x += p.CommandModeContent.Cells[n].Width()
		n += 1
	}
}

func (p *Editor) RefreshBuf() {
	p.refreshEditModeBuf()

	p.refreshCommandModeBuf()
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		if 0 == len(p.Lines) {
			p.EditModeAppendNewLine()
		}
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.Buf.IfNotRenderByTermUI = true
		p.CommandModeContentAreaY = p.Block.InnerArea.Max.Y - 1
		p.CommandModeContentAreaHeight = 1
		p.EditModeContentAreaHeight = p.Block.InnerArea.Dy() - p.CommandModeContentAreaHeight
		p.UIRender()
	} else {
		p.UIRender()
	}

	return *p.Buf
}

func (p *Editor) ActiveMode() {
	p.isEditModeContentDirty = true
	p.EditModeEnter()
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
}

func (p *Editor) UnActiveMode() {
	p.isEditModeContentDirty = true
	p.Mode = EDITOR_MODE_NONE
	p.CursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
