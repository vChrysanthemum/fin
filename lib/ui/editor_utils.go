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
	if p.CurrentLineIndex > p.DisplayLinesBottomIndex {
		p.CurrentLineIndex = p.DisplayLinesBottomIndex
	}

	switch p.Mode {
	case EDITOR_EDIT_MODE:
		p.EditModeCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
	case EDITOR_NORMAL_MODE:
		p.EditModeCursorLocation.RefreshCursorByEditorLine(p.CurrentLine())
	case EDITOR_COMMAND_MODE:
		p.CommandModeCursorLocation.RefreshCursorByEditorLine(p.CommandModeBuf)
	}
}
