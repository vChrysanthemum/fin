package ui

import (
	uiutils "fin/ui/utils"
	"os"
	"path/filepath"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

var (
	GCurrentRenderPage *Page
	GClearScreenBuffer *ClearScreenBuffer
)

var GlobalOption = Option{
	ResBaseDir:  filepath.Join(os.Getenv("HOME"), ".fin"),
	ProjectPath: filepath.Join(os.Getenv("HOME"), ".fin", "project", "traveller"),
}

type Option struct {
	ResBaseDir  string
	ProjectPath string
}

func Init(option Option) {
	GlobalOption = option
}

func init() {
	termui.ColorMap = map[string]termui.Attribute{
		"fg":           termui.ColorWhite,
		"bg":           termui.ColorDefault,
		"border.fg":    termui.ColorWhite,
		"label.fg":     termui.ColorWhite,
		"par.fg":       termui.ColorYellow,
		"par.label.bg": termui.ColorWhite,
	}

}

func PrepareUI() {
	err := termui.Init()
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		panic(err)
	}
	GClearScreenBuffer = NewClearScreenBuffer()
	registerHandles()
}

func uiClear() {
	uiutils.UIRender(GClearScreenBuffer)
}
