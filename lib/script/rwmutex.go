package script

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

func (p *Script) _getRWMutexFromUserData(L *lua.LState, lu *lua.LUserData) *sync.RWMutex {
	if nil == lu || nil == lu.Value {
		return nil
	}

	mutex, ok := lu.Value.(*sync.RWMutex)
	if false == ok || nil == mutex {
		return nil
	}

	return mutex
}

func (p *Script) NewRWMutex(L *lua.LState) int {
	mutex := L.NewUserData()
	mutex.Value = new(sync.RWMutex)
	L.Push(mutex)
	return 1
}

func (p *Script) RWMutexLock(L *lua.LState) int {
	if L.GetTop() < 1 {
		L.Push(lua.LFalse)
		return 1
	}

	lu := L.ToUserData(1)
	mutex := p._getRWMutexFromUserData(L, lu)
	if nil == mutex {
		L.Push(lua.LFalse)
		return 1
	}

	mutex.Lock()
	L.Push(lua.LTrue)
	return 1
}

func (p *Script) RWMutexUnlock(L *lua.LState) int {
	if L.GetTop() < 1 {
		L.Push(lua.LFalse)
		return 1
	}

	lu := L.ToUserData(1)
	mutex := p._getRWMutexFromUserData(L, lu)
	if nil == mutex {
		L.Push(lua.LFalse)
		return 1
	}

	mutex.Unlock()
	L.Push(lua.LTrue)
	return 1
}
