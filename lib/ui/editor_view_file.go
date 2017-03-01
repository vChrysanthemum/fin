package ui

import (
	"bufio"
	"os"

	termbox "github.com/nsf/termbox-go"
)

func (p *EditorView) LoadFile(filePath string) error {
	p = p.Editor.NewEditorView()
	p.Editor.EditorView = p
	p.FilePath = filePath

	f, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	if nil != err {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p.AppendLineData(scanner.Bytes())
	}

	p.isShouldRefreshInputModeBuf = true
	p.Editor.Buffer()
	termbox.Flush()

	return nil
}
