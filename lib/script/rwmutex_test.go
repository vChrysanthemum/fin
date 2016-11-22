package script

import (
	"fmt"
	"in/script"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

func TestRWMutex(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	var (
		err error
		s   script.Script
	)

	L := lua.NewState()
	s.RegisterScript(L)
	luaBase := L.NewTable()
	L.SetGlobal("base", luaBase)
	s.RegisterBaseTable(L, luaBase)

	err = L.DoFile(filepath.Join(GlobalOption.ResBaseDir, "lua/script/core.lua"))
	assert.Nil(t, err)

	content := fmt.Sprintf(`
	local rwmutex = require("rwmutex")
	local locker = rwmutex.NewRWMutex()
	locker:Lock()
	locker:Unlock()
	`)
	err = L.DoString(content)
	assert.Nil(t, err)
}
