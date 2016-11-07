package ui

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
	title := L.ToString(1)
	log.Println(title)
	return 0
}
