package editor

import (
	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

func str2runes(s string) []rune {
	return []rune(s)
}

func toTmAttr(x termui.Attribute) termbox.Attribute {
	return termbox.Attribute(x)
}

func (p *Editor) UIRender() {
	p.RefreshBuf()

	switch p.Mode {
	case EDITOR_EDIT_MODE:
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
	case EDITOR_NORMAL_MODE:
		p.CursorLocation.RefreshCursorByLine(p.CurrentLine())
	case EDITOR_COMMAND_MODE:
		p.CursorLocation.RefreshCursorByLine(p.CommandModeContent)
	}
}
