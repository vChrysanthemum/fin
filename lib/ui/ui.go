package ui

import (
	"os"
	"path/filepath"

	"github.com/gizak/termui"
)

var GClearScreenBuffer *ClearScreenBuffer

var GlobalOption = Option{
	ResBaseDir: filepath.Join(os.Getenv("HOME"), ".in"),
}

type Option struct {
	ResBaseDir string
}

func Init(option Option) {
	GlobalOption = option
}

func init() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	GClearScreenBuffer = NewClearScreenBuffer()

	termui.ColorMap = map[string]termui.Attribute{
		"fg":           termui.ColorWhite,
		"bg":           termui.ColorDefault,
		"border.fg":    termui.ColorWhite,
		"label.fg":     termui.ColorWhite,
		"par.fg":       termui.ColorYellow,
		"par.label.bg": termui.ColorWhite,
	}

}
