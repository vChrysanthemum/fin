package ui

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) LuaFuncLog(L *lua.LState) int {
	params := L.GetTop()
	var contents []string
	for i := 1; i <= params; i++ {
		contents = append(contents, L.ToString(i))
	}
	log.Println(contents)
	return 0
}
