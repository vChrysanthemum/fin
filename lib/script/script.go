package script

import (
	"os"
	"path/filepath"
	"sync"

	luajson "github.com/layeh/gopher-json"
	lua "github.com/yuin/gopher-lua"
)

var GlobalOption = Option{
	ResBaseDir:  filepath.Join(os.Getenv("HOME"), ".fin"),
	ProjectPath: filepath.Join(os.Getenv("HOME"), ".fin", "project", "traveller"),
}

type Option struct {
	ResBaseDir  string
	ProjectPath string
}

func Init(option Option) {
	GlobalOption = option
}

type Script struct {
	CancelSigs           map[string](chan bool)
	LuaCallByParamLocker *sync.RWMutex
}

func NewScript(LuaCallByParamLocker *sync.RWMutex) *Script {
	ret := new(Script)
	ret.CancelSigs = make(map[string](chan bool), 0)
	ret.LuaCallByParamLocker = LuaCallByParamLocker
	return ret
}

func (p *Script) RegisterBaseTable(L *lua.LState, baseTable *lua.LTable) {
	L.SetField(baseTable, "Quit", L.NewFunction(p.Quit))
	L.SetField(baseTable, "Log", L.NewFunction(p.Log))

	L.SetField(baseTable, "ResBaseDir", lua.LString(
		filepath.Join(GlobalOption.ResBaseDir),
	))

	L.SetField(baseTable, "Sleep", L.NewFunction(p.Sleep))
	L.SetField(baseTable, "SetInterval", L.NewFunction(p.SetInterval))
	L.SetField(baseTable, "SetTimeout", L.NewFunction(p.SetTimeout))
	L.SetField(baseTable, "SendCancelSig", L.NewFunction(p.SendCancelSig))

	L.SetField(baseTable, "WriteContentToFile", L.NewFunction(p.WriteContentToFile))
	L.SetField(baseTable, "ReadContentFromFile", L.NewFunction(p.ReadContentFromFile))

	L.SetField(baseTable, "OpenDB", L.NewFunction(p.OpenDB))
	L.SetField(baseTable, "CloseDB", L.NewFunction(p.CloseDB))
	L.SetField(baseTable, "DBQuery", L.NewFunction(p.DBQuery))
	L.SetField(baseTable, "DBRowsNext", L.NewFunction(p.DBRowsNext))
	L.SetField(baseTable, "DBRowsClose", L.NewFunction(p.DBRowsClose))
	L.SetField(baseTable, "DBExec", L.NewFunction(p.DBExec))
	L.SetField(baseTable, "DBResultLastInsertId", L.NewFunction(p.DBResultLastInsertId))
	L.SetField(baseTable, "DBResultRowsAffected", L.NewFunction(p.DBResultRowsAffected))

	L.SetField(baseTable, "NewRWMutex", L.NewFunction(p.NewRWMutex))
	L.SetField(baseTable, "RWMutexLock", L.NewFunction(p.RWMutexLock))
	L.SetField(baseTable, "RWMutexUnlock", L.NewFunction(p.RWMutexUnlock))

	L.SetField(baseTable, "NewUnixSockClient", L.NewFunction(p.NewUnixSockClient))
	L.SetField(baseTable, "UnixSockGet", L.NewFunction(p.UnixSockGet))
}

func (p *Script) RegisterScript(L *lua.LState) {
	luajson.Preload(L)
}
