package ui

import (
	"in/utils"

	"github.com/gizak/termui"
	lua "github.com/yuin/gopher-lua"
)

func (p *Script) luaFuncUIRerender(L *lua.LState) int {
	defer utils.RecoverPanic()
	p.page.SetActiveNode(nil)
	p.page.Rerender()
	return 0
}

func (p *Script) luaFuncWindowWidth(L *lua.LState) int {
	defer utils.RecoverPanic()
	L.Push(lua.LNumber(termui.TermWidth()))
	return 1
}

func (p *Script) luaFuncWindowHeight(L *lua.LState) int {
	defer utils.RecoverPanic()
	L.Push(lua.LNumber(termui.TermHeight()))
	return 1
}

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
	defer utils.RecoverPanic()
	if L.GetTop() < 1 {
		return 0
	}

	content := L.ToString(1)
	page, err := Parse(content)
	if nil != err {
		return 0
	}

	err = page.Render()
	if nil != err {
		return 0
	}

	nodeSelect, _ := page.IdToNodeMap["SelectConfirm"]
	if nil == nodeSelect {
		return 0
	}

	nodeSelectData := nodeSelect.Data.(*NodeSelect)

	p.page.ClearActiveNode(nodeSelect)
	page.uiRender()
	p.page.SetActiveNode(nodeSelect)

	nodeSelectData.DisableQuit = true
	nodeSelect.WaitKeyPressEnter()

	L.Push(lua.LString(nodeSelectData.NodeDataGetValue()))

	page.Clear()

	p.page.SetActiveNode(nil)
	p.page.Rerender()

	return 1
}
