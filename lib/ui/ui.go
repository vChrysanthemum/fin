package ui

import (
	"fin/ui/utils"
	"os"
	"path/filepath"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

var (
	GCurrentRenderPage *Page
	GClearScreenBuffer *ClearScreenBuffer
	GTermboxEvents     = make(chan termbox.Event, 200)
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
		"fg":           utils.COLOR_WHITE,
		"bg":           utils.COLOR_DEFAULT,
		"border.fg":    utils.COLOR_WHITE,
		"label.fg":     utils.COLOR_WHITE,
		"par.fg":       utils.COLOR_YELLOW,
		"par.label.bg": utils.COLOR_WHITE,
	}

}

func PrepareUI() {
	err := termui.MiniInit()
	termbox.SetOutputMode(termbox.Output256)
	if err != nil {
		panic(err)
	}
	GClearScreenBuffer = NewClearScreenBuffer()
}

func uiClear(startY, endY int) {
	if 0 == startY && -1 == endY {
		termui.Clear()
	} else {
		GClearScreenBuffer.Buf.Area.Min.X = 0
		GClearScreenBuffer.Buf.Area.Max.X = termui.TermWidth()
		GClearScreenBuffer.Buf.Area.Min.Y = startY
		GClearScreenBuffer.Buf.Area.Max.Y = endY
		utils.UIRender(GClearScreenBuffer)
	}
}
