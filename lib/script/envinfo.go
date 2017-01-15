package script

import (
	"runtime"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) GetMemAlloc(L *lua.LState) int {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	L.Push(lua.LNumber(stats.Alloc))
	return 1
}
