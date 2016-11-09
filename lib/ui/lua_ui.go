package ui

import (
	"github.com/gizak/termui"
	lua "github.com/yuin/gopher-lua"
)

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
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

	p.page.SetActiveNode(nodeSelect)

	termui.Render(page.Bufferers...)

	NodeSelectOption := nodeSelect.Data.(*NodeSelect).WaitResult()
	L.Push(lua.LString(NodeSelectOption.Value))

	p.page.SetActiveNode(nil)
	p.page.Refresh()

	return 1
}
