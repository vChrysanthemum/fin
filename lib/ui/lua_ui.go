package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) luaFuncWindowConfirm(L *lua.LState) int {
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

	uiRender(page.Bufferers...)

	p.page.SetActiveNode(nodeSelect)

	nodeSelectData.DisableQuit = true
	nodeSelect.WaitKeyPressEnter()

	L.Push(lua.LString(nodeSelectData.GetValue()))

	page.Clear()

	p.page.SetActiveNode(nil)
	p.page.Rerender()

	return 1
}
