package ui

func (p *Editor) LoadFile(filePath string) error {
	return p.EditorView.LoadFile(filePath)
}
