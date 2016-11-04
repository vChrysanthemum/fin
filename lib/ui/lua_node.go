package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) luaFuncGetNodePointer(L *lua.LState) int {
	nodeId := L.ToString(1)

	var (
		node *Node
		ok   bool
	)

	node, ok = p.page.IdToNodeMap[nodeId]

	if true == ok {
		luaNode := L.NewUserData()
		luaNode.Value = node
		L.Push(luaNode)
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (p *Script) luaFuncGetNodeHtmlData(L *lua.LState) int {
	var (
		node *Node
		ok   bool
	)

	lv := L.ToUserData(1)
	if nil == lv || nil == lv.Value {
		L.Push(lua.LNil)
		return 1
	}

	node, ok = lv.Value.(*Node)
	if false == ok || nil == node {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(node.HtmlData))
	}

	return 1
}
