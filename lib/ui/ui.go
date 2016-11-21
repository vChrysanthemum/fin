package ui

import (
	uiutils "in/ui/utils"
	"os"
	"path/filepath"

	"github.com/gizak/termui"
)

var (
	GClearScreenBuffer *ClearScreenBuffer
)

var GlobalOption = Option{
	ResBaseDir:  filepath.Join(os.Getenv("HOME"), ".in"),
	ProjectName: "",
}

type Option struct {
	ResBaseDir  string
	ProjectName string
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
	if err != nil {
		panic(err)
	}
	GClearScreenBuffer = NewClearScreenBuffer()
}

func uiClear() {
	uiutils.UIRender(GClearScreenBuffer)
}
