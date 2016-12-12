package ui

import uuid "github.com/satori/go.uuid"

func (p *Node) RegisterLuaActiveModeHandler(handler NodeJobHandler, args ...interface{}) string {
	key := uuid.NewV4().String()
	p.LuaActiveModeHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RemoveLuaActiveModeHandler(key string) {
	delete(p.LuaActiveModeHandlers, key)
}

func (p *Node) RegisterKeyPressHandler(handler NodeJobHandler, args ...interface{}) string {
	key := uuid.NewV4().String()
	p.KeyPressHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RemoveKeyPressHandler(key string) {
	delete(p.KeyPressHandlers, key)
}

func (p *Node) RegisterKeyPressEnterHandler(handler NodeJobHandler, args ...interface{}) string {
	key := uuid.NewV4().String()
	p.KeyPressEnterHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RemoveKeyPressEnterHandler(key string) {
	delete(p.KeyPressEnterHandlers, key)
}
