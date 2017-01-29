package editor

func (p *Editor) PrepareCommandMode() {
	p.CommandModeContent = p.NewLine()
}

func (p *Editor) CommandModeQuit() {
	p.isCommandModeContentDirty = true
	p.CommandModeContent.CleanData()
}

func (p *Editor) CommandModeEnter() {
	p.Mode = EDITOR_COMMAND_MODE
	p.CommandModeContent.CleanData()
	p.CommandModeWrite(":")
}

func (p *Editor) CommandModeWrite(keyStr string) {
	p.isCommandModeContentDirty = true
	p.CommandModeContent.ContentStartX = p.InnerArea.Min.X
	p.CommandModeContent.ContentStartY = p.CommandModeContentAreaY

	if "<enter>" == keyStr {
		p.CommandModeQuit()
		p.NormalModeEnter()

	} else if "C-8" == keyStr {
		p.CommandModeContent.Backspace()

	} else if "<left>" == keyStr {
		p.CursorLocation.MoveCursorNRuneLeft(1)

	} else if "<right>" == keyStr {
		p.CursorLocation.MoveCursorNRuneRight(1)

	} else {
		if "<space>" == keyStr {
			keyStr = " "
		} else if "<tab>" == keyStr {
			keyStr = "\t"
		}
		p.CommandModeContent.Write(keyStr)
	}

	p.UIRender()
}
