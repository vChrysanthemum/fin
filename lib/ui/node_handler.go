package ui

import uuid "github.com/satori/go.uuid"

func (p *Node) RegisterKeyPressHandler(handler NodeJobHandler, args ...interface{}) string {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()

	key := uuid.NewV4().String()
	p.KeyPressHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RegisterLuaActiveModeHandler(handler NodeJobHandler, args ...interface{}) string {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()

	key := uuid.NewV4().String()
	p.LuaActiveModeHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RemoveLuaActiveModeHandler(key string) {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()
	delete(p.LuaActiveModeHandlers, key)
}

func (p *Node) RemoveKeyPressHandler(key string) {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()
	delete(p.KeyPressHandlers, key)
}

func (p *Node) RegisterKeyPressEnterHandler(handler NodeJobHandler, args ...interface{}) string {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()

	key := uuid.NewV4().String()
	p.KeyPressEnterHandlers[key] = NodeJob{
		Node:    p,
		Handler: handler,
		Args:    args,
	}
	return key
}

func (p *Node) RemoveKeyPressEnterHandler(key string) {
	p.JobHanderLocker.Lock()
	defer p.JobHanderLocker.Unlock()
	delete(p.KeyPressEnterHandlers, key)
}
