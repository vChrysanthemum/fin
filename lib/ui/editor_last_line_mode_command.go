package ui

import (
	"fmt"
	"regexp"
)

type LastLineModeCommandHandler func(matchKey interface{})

type EditorLastLineModeCommand struct {
	MatchKey interface{}
	Handler  LastLineModeCommandHandler
}

func (p *Editor) PrepareLastLineModeCommand() {
	p.LastLineModeCommands = []EditorLastLineModeCommand{
		{byte('w'), p.lastLineCommandSaveFile},
	}
}

func (p *Editor) lastLineCommandSaveFile(matchKey interface{}) {
	err := p.EditorView.SaveFile()
	if nil == err {
		p.CommandShowMsg(fmt.Sprintf("%v saved", p.EditorView.FilePath))
	} else {
		p.CommandShowMsg(fmt.Sprintf("%v", err.Error()))
	}
}

func (p *Editor) ExecLastLineCommand() {
	if len(p.LastLineModeBuf.Data) <= 1 {
		return
	}
	p.LastLineModeBuf.Data = p.LastLineModeBuf.Data[1:]

	for _, cmd := range p.Editor.LastLineModeCommands {
		switch cmd.MatchKey.(type) {
		case *regexp.Regexp:
			if true == cmd.MatchKey.(*regexp.Regexp).Match(p.LastLineModeBuf.Data) {
				cmd.Handler(cmd.MatchKey)
				return
			}
		case byte:
			if p.LastLineModeBuf.Data[0] == cmd.MatchKey.(byte) {
				cmd.Handler(cmd.MatchKey)
				return
			}
		case string:
			matchkey := cmd.MatchKey.(string)
			if len(p.LastLineModeBuf.Data) >= len(matchkey) &&
				string(p.LastLineModeBuf.Data) == matchkey {
				cmd.Handler(cmd.MatchKey)
				return
			}
		}
	}
}
