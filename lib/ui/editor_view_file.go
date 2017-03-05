package ui

import (
	"bufio"
	"os"

	termbox "github.com/nsf/termbox-go"
)

func (p *EditorView) LoadFile(filePath string) error {
	p.Prepare()
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

func (p *EditorView) SaveFile() error {
	if "" == p.FilePath {
		return nil
	}

	f, err := os.OpenFile(p.FilePath, os.O_RDWR|os.O_CREATE, 0777)
	if nil != err {
		return err
	}
	defer f.Close()

	for k, _ := range p.Lines {
		f.Write(p.Lines[k].Data)
		f.Write([]byte("\n"))
	}

	return nil
}
