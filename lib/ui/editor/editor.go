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

	termui.Block

	TextFgColor       termui.Attribute
	TextBgColor       termui.Attribute
	DisplayLinesRange [2]int
	*CursorLocation
}

func NewEditor() *Editor {
	ret := &Editor{
		Lines:             []*Line{},
		Block:             *termui.NewBlock(),
		TextFgColor:       termui.ThemeAttr("par.text.fg"),
		TextBgColor:       termui.ThemeAttr("par.text.bg"),
		DisplayLinesRange: [2]int{0, 1},
	}
	ret.Mode = EDITOR_MODE_NONE
	ret.ModeWrite = nil
	ret.PrepareNormalMode()
	ret.PrepareEditMode()
	ret.PrepareCommandMode()
	ret.CursorLocation = NewCursorLocation(ret)
	return ret
}

func (p *Editor) Text() []*Line {
	var printLines []*Line
	for k, line := range p.Lines {
		if k < p.DisplayLinesRange[0] {
			continue
		}
		if k >= p.DisplayLinesRange[1] {
			continue
		}
		printLines = append(printLines, line)
	}
	return printLines
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

func (p *Editor) Buffer() termui.Buffer {
	buf := p.Block.Buffer()

	fg, bg := p.TextFgColor, p.TextBgColor
	lines := p.Text()
	for k, _ := range lines {
		lines[k].Cells = termui.DefaultTxBuilder.Build(string(lines[k].Data), fg, bg)
	}

	finalX, finalY := 0, 0
	y, x, n, w := 0, 0, 0, 0
	for _, line := range lines {
		line.ContentStartY = y
		n = 0
		for n < len(line.Cells) {
			w = line.Cells[n].Width()
			if x+w > p.InnerArea.Dx() {
				x = 0
				y++
				if y >= p.InnerArea.Dy() {
					goto BUFFER_END
				}

				continue
			}

			finalX = p.Block.InnerArea.Min.X + x
			finalY = p.Block.InnerArea.Min.Y + y
			buf.Set(finalX, finalY, line.Cells[n])
			line.Cells[n].X, line.Cells[n].Y = finalX, finalY

			n++
			x += w
		}

		x = 0
		y++
		if y >= p.InnerArea.Dy() {
			goto BUFFER_END
		}
	}

BUFFER_END:
	return buf
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
