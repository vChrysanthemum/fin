package ui

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
	switch p.Mode {
	case EDITOR_EDIT_MODE:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EDITOR_NORMAL_MODE:
		p.EditModeCursor.RefreshCursorByEditorLine(p.EditModeCursor.Line())
	case EDITOR_COMMAND_MODE:
		p.CommandModeCursor.RefreshCursorByEditorLine(p.CommandModeBuf)
	}
}
