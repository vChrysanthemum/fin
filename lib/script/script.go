package script

import (
	"os"
	"path/filepath"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

var (
	GUIRenderLocker sync.RWMutex
)

var GlobalOption = Option{
	ResBaseDir: filepath.Join(os.Getenv("HOME"), ".in"),
}

type Option struct {
	ResBaseDir string
}

func Init(option Option) {
	GlobalOption = option
}

type Script struct {
	CancelSigs map[string](chan bool)
}

func (p *Script) RegisterInLuaTable(L *lua.LState, table *lua.LTable) {
	L.SetField(table, "Log", L.NewFunction(p.Log))

	L.SetField(table, "SetInterval", L.NewFunction(p.SetInterval))
	L.SetField(table, "SetTimeout", L.NewFunction(p.SetTimeout))
	L.SetField(table, "SendCancelSig", L.NewFunction(p.SendCancelSig))

	L.SetField(table, "OpenDB", L.NewFunction(p.OpenDB))
	L.SetField(table, "CloseDB", L.NewFunction(p.CloseDB))
	L.SetField(table, "DBQuery", L.NewFunction(p.DBQuery))
	L.SetField(table, "DBRowsNext", L.NewFunction(p.DBRowsNext))
	L.SetField(table, "DBRowsClose", L.NewFunction(p.DBRowsClose))
	L.SetField(table, "DBExec", L.NewFunction(p.DBExec))
	L.SetField(table, "DBResultLastInsertId", L.NewFunction(p.DBResultLastInsertId))
	L.SetField(table, "DBResultRowsAffected", L.NewFunction(p.DBResultRowsAffected))
}
