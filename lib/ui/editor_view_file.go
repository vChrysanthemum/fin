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

	p.Lines = []*EditorLine{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p.AppendLineWithData(scanner.Bytes())
	}

	if 0 == len(p.Lines) {
		p.InputModeAppendNewLine(p.InputModeCursor)
	}

	p.isShouldRefreshInputModeBuf = true
	p.Editor.Buffer()
	termbox.Flush()

	return nil
}
