package ui

import (
	rw "github.com/mattn/go-runewidth"
)

func FormatStringWithWidth(src string, width int) string {
	tail := width - rw.StringWidth(src)
	if tail > 0 {
		return src + string(make([]byte, tail))
	}
	return src
}
