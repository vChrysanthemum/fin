package script

import (
	"os"
	"path/filepath"

	luajson "github.com/layeh/gopher-json"
	lua "github.com/yuin/gopher-lua"
)

var GlobalOption = Option{
	ResBaseDir:  filepath.Join(os.Getenv("HOME"), ".in"),
	ProjectName: "",
}

type Option struct {
	ResBaseDir  string
	ProjectName string
}

func Init(option Option) {
	GlobalOption = option
}

type Script struct {
	CancelSigs map[string](chan bool)
}

func NewScript() *Script {
	ret := new(Script)
	ret.CancelSigs = make(map[string](chan bool), 0)
	return ret
}

func (p *Script) RegisterBaseTable(L *lua.LState, baseTable *lua.LTable) {
	L.SetField(baseTable, "Log", L.NewFunction(p.Log))

	L.SetField(baseTable, "ResBaseDir", lua.LString(
		filepath.Join(GlobalOption.ResBaseDir),
	))

	L.SetField(baseTable, "SetInterval", L.NewFunction(p.SetInterval))
	L.SetField(baseTable, "SetTimeout", L.NewFunction(p.SetTimeout))
	L.SetField(baseTable, "SendCancelSig", L.NewFunction(p.SendCancelSig))

	L.SetField(baseTable, "OpenDB", L.NewFunction(p.OpenDB))
	L.SetField(baseTable, "CloseDB", L.NewFunction(p.CloseDB))
	L.SetField(baseTable, "DBQuery", L.NewFunction(p.DBQuery))
	L.SetField(baseTable, "DBRowsNext", L.NewFunction(p.DBRowsNext))
	L.SetField(baseTable, "DBRowsClose", L.NewFunction(p.DBRowsClose))
	L.SetField(baseTable, "DBExec", L.NewFunction(p.DBExec))
	L.SetField(baseTable, "DBResultLastInsertId", L.NewFunction(p.DBResultLastInsertId))
	L.SetField(baseTable, "DBResultRowsAffected", L.NewFunction(p.DBResultRowsAffected))
}

func (p *Script) RegisterScript(L *lua.LState) {
	luajson.Preload(L)
}
