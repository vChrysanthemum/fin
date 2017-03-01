package ui

import lua "github.com/yuin/gopher-lua"

func (p *Script) _getNodeEditorPointerFromUserData(L *lua.LState, lu *lua.LUserData) *NodeEditor {
	if nil == lu || nil == lu.Value {
		return nil
	}

	var (
		node       *Node
		nodeEditor *NodeEditor
		ok         bool
	)

	node, ok = lu.Value.(*Node)
	if false == ok || nil == node {
		return nil
	}

	if nodeEditor, ok = node.Data.(*NodeEditor); false == ok {
		return nil
	}

	return nodeEditor
}

func (p *Script) luaFuncNodeEditorLoadFile(L *lua.LState) int {
	if L.GetTop() < 2 {
		L.Push(lua.LNil)
		return 1
	}

	nodeEditor := p._getNodeEditorPointerFromUserData(L, L.ToUserData(1))
	err := nodeEditor.Editor.LoadFile(L.ToString(2))
	if nil != err {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 0
}
