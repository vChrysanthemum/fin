package script

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	lua "github.com/yuin/gopher-lua"
)

func (p *Script) Log(L *lua.LState) int {
	params := L.GetTop()
	var contents []string
	for i := 1; i <= params; i++ {
		contents = append(contents, L.ToString(i))
	}
	log.Println(contents)
	return 0
}

func (p *Script) Sleep(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	tm := L.ToInt(1)
	if tm < 0 {
		tm = 0
	}
	time.Sleep(time.Duration(tm) * time.Millisecond)
	return 0
}

func (p *Script) SetInterval(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	tm := L.ToInt(1)
	if tm < 0 {
		tm = 0
	}
	callback := L.ToFunction(2)

	sigKey := uuid.NewV4().String()
	cancel := make(chan bool, 0)
	p.CancelSigs[sigKey] = cancel

	go func(_cancel chan bool, _tm int, _L *lua.LState, _callback *lua.LFunction) {
		for {
			select {
			case <-_cancel:
				return

			default:
				time.Sleep(time.Duration(_tm) * time.Millisecond)
				if err := p.luaCallByParam(_L, lua.P{
					Fn:      _callback,
					NRet:    0,
					Protect: true,
				}); err != nil {
					log.Println(err)
					panic(err)
				}
			}
		}
	}(cancel, tm, L, callback)

	L.Push(lua.LString(sigKey))
	return 1
}

func (p *Script) SetTimeout(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	tm := L.ToInt(1)
	if tm < 0 {
		tm = 0
	}
	callback := L.ToFunction(2)

	sigKey := uuid.NewV4().String()
	cancel := make(chan bool, 0)
	p.CancelSigs[sigKey] = cancel

	go func(_script *Script, _cancel chan bool,
		_tm int, _sigKey string, _L *lua.LState, _callback *lua.LFunction) {

		select {
		case <-_cancel:
			return

		default:
			time.Sleep(time.Duration(_tm) * time.Millisecond)
			if err := p.luaCallByParam(_L, lua.P{
				Fn:      _callback,
				NRet:    0,
				Protect: true,
			}); err != nil {
				log.Println(err)
				panic(err)
			}
			delete(_script.CancelSigs, _sigKey)
			close(_cancel)
			return
		}
	}(p, cancel, tm, sigKey, L, callback)

	L.Push(lua.LString(sigKey))
	return 1
}

func (p *Script) SendCancelSig(L *lua.LState) int {
	if L.GetTop() < 1 {
		return 0
	}

	sigKey := L.ToString(1)
	cancel, ok := p.CancelSigs[sigKey]
	if false == ok {
		return 0
	}
	close(cancel)
	delete(p.CancelSigs, sigKey)
	return 0
}

func (p *Script) luaCallByParam(L *lua.LState, cp lua.P, args ...lua.LValue) error {
	p.LuaCallByParamLocker.Lock()
	defer func() {
		p.LuaCallByParamLocker.Unlock()
		if rcv := recover(); nil != rcv {
			log.Println(rcv)
		}
	}()
	return L.CallByParam(cp, args...)
}
