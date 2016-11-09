package ui

import lua "github.com/yuin/gopher-lua"

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

	nodeSelectData := nodeSelect.Data.(*NodeSelect)

	uirender(page.Bufferers...)

	p.page.SetActiveNode(nodeSelect)

	nodeSelect.OnKeyPressEnter()
	nodeSelectData.DisableQuit = true
	L.Push(lua.LString(nodeSelectData.Children[nodeSelectData.SelectedOptionIndex].Value))

	page.Clear()

	p.page.SetActiveNode(nil)
	p.page.Rerender()

	return 1
}
