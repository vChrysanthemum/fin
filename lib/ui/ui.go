package ui

import (
	"os"
	"path/filepath"
)

var GlobalOption = Option{
	LuaResBaseDir: filepath.Join(os.Getenv("HOME"), ".in/lua/"),
}

type Option struct {
	LuaResBaseDir string
}

func Init(option Option) {
	GlobalOption = option
}
