package editor

func (p *Editor) PrepareCommandMode() {
}

func (p *Editor) CommandModeEnter() {
	p.Mode = EDITOR_EDIT_MODE
	p.ModeWrite = p.CommandModeWrite
}

func (p *Editor) CommandModeWrite(keyStr string) {
}
