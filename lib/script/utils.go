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

func (p *Script) SetInterval(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	tm := L.ToInt(1)
	callback := L.ToFunction(2)

	sigKey := uuid.NewV4().String()
	cancel := make(chan bool, 0)
	p.CancelSigs[sigKey] = cancel

	go func(_cancel chan bool, _tm int, _L *lua.LState, _callback *lua.LFunction) {
		for {
			select {
			case <-_cancel:
				goto END

			default:
				time.Sleep(time.Duration(_tm) * time.Millisecond)
				if err := _L.CallByParam(lua.P{
					Fn:      _callback,
					NRet:    0,
					Protect: true,
				}); err != nil {
					log.Println(err)
					panic(err)
				}
			}
		}
	END:
	}(cancel, tm, L, callback)

	L.Push(lua.LString(sigKey))
	return 1
}

func (p *Script) SetTimeout(L *lua.LState) int {
	if L.GetTop() < 2 {
		return 0
	}

	tm := L.ToInt(1)
	callback := L.ToFunction(2)

	sigKey := uuid.NewV4().String()
	cancel := make(chan bool, 0)
	p.CancelSigs[sigKey] = cancel

	go func(_cancel chan bool, _tm int, _sigKey string, _L *lua.LState, _callback *lua.LFunction) {
		select {
		case <-_cancel:
			goto END

		default:
			time.Sleep(time.Duration(_tm) * time.Millisecond)
			if err := _L.CallByParam(lua.P{
				Fn:      _callback,
				NRet:    0,
				Protect: true,
			}); err != nil {
				log.Println(err)
				panic(err)
			}
		}
	END:
		delete(p.CancelSigs, _sigKey)
	}(cancel, tm, sigKey, L, callback)

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
	cancel <- true
	delete(p.CancelSigs, sigKey)
	return 0
}
