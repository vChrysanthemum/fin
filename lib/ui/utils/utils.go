package utils

import (
	"fmt"
	"image"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

var (
	GUIRenderLocker sync.RWMutex
)

func FormatStringWithWidth(src string, width int) string {
	tail := width - rw.StringWidth(src)
	if tail > 0 {
		return src + string(make([]byte, tail))
	}
	return src
}

func StringToBool(str string, defaultVal bool) bool {
	if "true" == str {
		return true
	} else if "false" == str {
		return false
	}
	return defaultVal
}

func ColorToTermuiAttribute(colorsStr string, defaultColor termui.Attribute) termui.Attribute {
	if "" == colorsStr {
		return defaultColor
	}

	colors := strings.Split(colorsStr, "|")

	var color termui.Attribute
	for _, colorStr := range colors {
		tmp, err := strconv.ParseInt(colorStr, 0, 0)
		if nil == err {
			color = termui.Attribute(int(tmp))
			continue
		}

		switch colorStr {
		case "white":
			color |= COLOR_WHITE
		case "black":
			color |= COLOR_BLACK
		case "red":
			color |= COLOR_RED
		case "green":
			color |= COLOR_GREEN
		case "yellow":
			color |= COLOR_YELLOW
		case "blue":
			color |= COLOR_BLUE
		case "magenta":
			color |= COLOR_MAGENTA
		case "cyan":
			color |= COLOR_CYAN
		case "gray":
			color |= COLOR_GRAY
		case "bold":
			color |= termui.AttrBold
		case "underline":
			color |= termui.AttrUnderline
		case "reverse":
			color |= termui.AttrReverse
		}
	}

	return color
}

func MaxInt(data ...int) int {
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

func Beep() {
	fmt.Println("\a")
}

func CalculateTextHeight(text string, widthLimited int) (height int) {
	buf := []byte(text)
	var ch rune
	x, w := 0, 0
	height = 1
	for len(buf) > 0 {
		ch, w = utf8.DecodeRune(buf)
		buf = buf[w:]
		if ch == '\n' || x+w > widthLimited {
			x = 0 // set x = 0
			height += 1
			continue
		}

		x += rw.RuneWidth(ch)
	}

	return
}

func CalculateTextLastPosition(text string, innerArea image.Rectangle) (resultX, resultY int) {
	buf := []byte(text)

	var ch rune
	y, x, n, w := 0, 0, 0, 0
	for y < innerArea.Dy() && n < len(buf) {
		ch, w = utf8.DecodeRune(buf)
		buf = buf[w:]
		if ch == '\n' || x+w > innerArea.Dx() {
			y++
			x = 0 // set x = 0
			if ch == '\n' {
				n++
			}

			continue
		}

		n += w
		x += rw.RuneWidth(ch)
	}

	resultX = innerArea.Min.X + x
	resultY = innerArea.Min.Y + y

	return
}

func UIRender(bs ...termui.Bufferer) {
	GUIRenderLocker.Lock()
	termui.Render(bs...)
	GUIRenderLocker.Unlock()
}

func UISetCursor(x, y int) {
	termbox.SetCursor(x, y)
}
