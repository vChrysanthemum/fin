package ui

import (
	"bufio"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gizak/termui"
)

type ClearScreenBuffer struct {
	buf termui.Buffer
}

func NewClearScreenBuffer() *ClearScreenBuffer {
	buf := termui.NewBuffer()
	min := image.Point{0, 0}
	max := image.Point{termui.TermWidth(), termui.TermHeight()}
	buf.SetArea(image.Rectangle{min, max})
	buf.Fill(' ', termui.ColorDefault, termui.ColorDefault)
	return &ClearScreenBuffer{
		buf: buf,
	}
}

func (p *ClearScreenBuffer) Buffer() termui.Buffer {
	return p.buf
}

func (p *ClearScreenBuffer) RefreshArea() {
	min := image.Point{0, 0}
	max := image.Point{termui.TermWidth() - 1, termui.TermHeight() - 1}
	p.buf.SetArea(image.Rectangle{min, max})
}

func GetFileContent(path string) ([]byte, error) {
	path = filepath.Join(GlobalOption.ResBaseDir, "project", GlobalOption.ProjectName, path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(bufio.NewReader(file))
}
