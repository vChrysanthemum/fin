package utils

import (
	"fmt"
	"image"
	"log"
	"runtime/debug"
	"unicode/utf8"

	"github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
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

func ColorToTermuiAttribute(color string, defaultColor termui.Attribute) termui.Attribute {
	switch color {
	case "black":
		return termui.ColorBlack
	case "red":
		return termui.ColorRed
	case "green":
		return termui.ColorGreen
	case "yellow":
		return termui.ColorYellow
	case "blue":
		return termui.ColorBlue
	case "magenta":
		return termui.ColorMagenta
	case "cyan":
		return termui.ColorCyan
	case "white":
		return termui.ColorWhite
	}

	return defaultColor
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

		x += w
	}

	return
}

func CalculateTextLastPosition(text string, innerArea image.Rectangle) (resultX, resultY int) {
	buf := []byte(text)

	var ch rune
	y, x, n, w := 0, 0, 0, 0
	for y < innerArea.Dy() && n < len(buf) {
		ch, w = utf8.DecodeRune(buf)
		if ch == '\n' || x+w > innerArea.Dx() {
			y++
			x = 0 // set x = 0
			if ch == '\n' {
				n++
			}

			continue
		}

		n++
		x += w
	}

	resultX = innerArea.Min.X + x
	resultY = innerArea.Min.Y + y

	return
}

func UISetCursor(x, y int) {
	log.Println(x, y, string(debug.Stack()))
	termbox.SetCursor(x, y)
}
