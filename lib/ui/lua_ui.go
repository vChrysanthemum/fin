package ui

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
	content := L.ToString(1)
	log.Println(content)
	page, err := Parse(content)
	if nil != err {
		return 0
	}
	page.DumpNodesHtmlData()

	return 0
}
