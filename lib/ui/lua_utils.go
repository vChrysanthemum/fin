package ui

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) LuaFuncLog(L *lua.LState) int {
	content := L.ToString(1)
	log.Println(content)
	return 0
}
