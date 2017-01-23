package editor

func (p *Editor) PrepareCommandMode() {
}

func (p *Editor) CommandModeEnter() {
	p.Mode = EDITOR_EDIT_MODE
}

func (p *Editor) CommandModeWrite(keyStr string) {
}
