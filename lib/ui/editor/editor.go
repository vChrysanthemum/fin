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

	TextFgColor          termui.Attribute
	TextBgColor          termui.Attribute
	DisplayLinesTopIndex int
	*CursorLocation
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
	return ret
}

func (p *Editor) UpdateCurrentLineData(line string) {
	p.CurrentLine.Data = []byte(line)
}

func (p *Editor) WriteNewLine(line string) {
	if 0 == len(p.Lines) {
		p.CurrentLine = p.InitNewLine()
	}

	// 如果上一行不为空，则启用新一行
	// 反之则利用上一行
	if len(p.CurrentLine.Data) > 0 {
		p.CurrentLine = p.InitNewLine()
	}

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
	fg, bg := p.TextFgColor, p.TextBgColor

	var (
		finalX, finalY int
		y, x, n, w     int
	)

START_REFRESH:
	buf := p.Block.Buffer()
	p.Buf = &buf
	finalX, finalY = 0, 0
	y, x, n, w = 0, 0, 0, 0
	for _, line := range p.Lines[p.DisplayLinesTopIndex:] {
		line.Cells = termui.DefaultTxBuilder.Build(string(line.Data), fg, bg)
		line.ContentStartY = y + p.Block.InnerArea.Min.Y
		n = 0
		for n < len(line.Cells) {
			w = line.Cells[n].Width()
			if x+w > p.Block.InnerArea.Dx() {
				x = 0
				y++
				if y > p.Block.InnerArea.Dy() {
					goto CHECK_LAST_LINE
				}

				continue
			}

			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			p.Buf.Set(finalX, finalY, line.Cells[n])
			line.Cells[n].X, line.Cells[n].Y = finalX, finalY

			n++
			x += w
		}

		x = 0
		y++

	CHECK_LAST_LINE:
		if y > p.Block.InnerArea.Dy() {
			if p.CurrentLine == line && p.DisplayLinesTopIndex < len(p.Lines)-1 {
				p.DisplayLinesTopIndex += 1
				goto START_REFRESH
			} else {
				goto REFRESH_END
			}
		}
	}

REFRESH_END:
	return
}

func (p *Editor) Buffer() termui.Buffer {
	if nil == p.Buf {
		buf := p.Block.Buffer()
		p.Buf = &buf
		p.RefreshContent()
	}

	p.Block.DrawBorder(*p.Buf)
	p.Block.DrawBorderLabel(*p.Buf)

	return *p.Buf
}

func (p *Editor) ActiveMode() {
	p.EditModeEnter()
	p.CursorLocation.IsDisplay = true
	p.CursorLocation.ResumeCursor()
}

func (p *Editor) UnActiveMode() {
	p.Mode = EDITOR_MODE_NONE
	p.ModeWrite = nil
	p.CursorLocation.IsDisplay = false
	uiutils.UISetCursor(-1, -1)
}
