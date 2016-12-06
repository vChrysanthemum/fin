package ui

import "github.com/gizak/termui"

func (p *Page) registerHandles() {
	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		p.Rerender()
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.DefaultEvtStream.ResetHandlers()
		termui.StopLoop()
		termui.Close()
	})

	termui.Handle("/sys/kbd", func(e termui.Event) {
		p.KeyPressHandleLocker.Lock()
		defer p.KeyPressHandleLocker.Unlock()
		keyStr := e.Data.(termui.EvtKbd).KeyStr

		if nil != p.ActiveNode {
			if nil != p.ActiveNode.KeyPress {
				p.ActiveNode.KeyPress(e)
			}

			// p.ActiveNode.KeyPress后，p.ActiveNode有可能为nil
			if nil == p.ActiveNode {
				return
			}

			if len(p.ActiveNode.KeyPressHandlers) > 0 {
				node := p.ActiveNode
				node.JobHanderLocker.RLock()
				for _, v := range p.ActiveNode.KeyPressHandlers {
					v.Args = append(v.Args, e)
					v.Handler(p.ActiveNode, v.Args...)
				}
				node.JobHanderLocker.RUnlock()
			}

			if "<escape>" == keyStr {
				p.ActiveNode.QuitActiveMode()
				return
			}

			return
		}

		// 切换ActiveNode
		if "<tab>" == keyStr ||
			"<up>" == keyStr || "<down>" == keyStr ||
			"<left>" == keyStr || "<right>" == keyStr ||
			"h" == keyStr || "j" == keyStr || "k" == keyStr || "l" == keyStr {

			if nil != p.FocusNode {
				if nodeDataUnFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer); true == ok {
					nodeDataUnFocusModer.NodeDataUnFocusMode()
				}
			}

			if nil == p.FocusNode {
				p.FocusNode = p.WorkingNodes.Front()
			} else {
				node := p.FocusNode.Value.(*Node)
				if "<tab>" == keyStr || "<right>" == keyStr || "l" == keyStr {
					if nil != p.FocusNode.Next() {
						p.FocusNode = p.FocusNode.Next()
					}

				} else if "<left>" == keyStr || "h" == keyStr {
					// "<left>" == keyStr
					if nil != p.FocusNode.Prev() {
						p.FocusNode = p.FocusNode.Prev()
					}

				} else if "<down>" == keyStr || "j" == keyStr {
					if nil != node.FocusBottomNode {
						p.FocusNode = node.FocusBottomNode
					}

				} else if "<up>" == keyStr || "k" == keyStr {
					if nil != node.FocusTopNode {
						p.FocusNode = node.FocusTopNode
					}
				}
			}

			if nil != p.FocusNode {
				if nodeDataFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataFocusModer); true == ok {
					nodeDataFocusModer.NodeDataFocusMode()
				}
			}

		} else if "<enter>" == keyStr {
			// 确认ActiveNode
			if nil != p.FocusNode {
				if nodeDataUnFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer); true == ok {
					nodeDataUnFocusModer.NodeDataUnFocusMode()
				}

				p.SetActiveNode(p.FocusNode.Value.(*Node))
			}
		}
	})
}

func (p *Page) pushWorkingNode(node *Node) {
	p.WorkingNodes.PushBack(node)
	p.FocusNode = p.WorkingNodes.Back()
}

func (p *Node) QuitActiveMode() {
	p.page.ActiveNodeAfterRerender = nil

	if nodeDataUnActiveModer, ok := p.Data.(NodeDataUnActiveModer); true == ok {
		nodeDataUnActiveModer.NodeDataUnActiveMode()
	}

	p.page.FocusNode = p.FocusThisNode

	if nil != p.page.FocusNode {
		if nodeDataFocusModer, ok := p.page.FocusNode.Value.(*Node).Data.(NodeDataFocusModer); true == ok {
			nodeDataFocusModer.NodeDataFocusMode()
		}
	}

	p.page.ActiveNode = nil
}

func (p *Page) ClearActiveNode(node *Node) {
	if nil != p.ActiveNode {
		if nodeDataUnActiveModer, ok := p.ActiveNode.Data.(NodeDataUnActiveModer); true == ok {
			nodeDataUnActiveModer.NodeDataUnActiveMode()
		}
	}
	p.ActiveNode = nil
}

func (p *Page) SetActiveNode(node *Node) {
	p.ClearActiveNode(node)
	p.ActiveNodeAfterRerender = node
	p.ActiveNode = node
	if nil != p.ActiveNode {
		if len(p.ActiveNode.LuaActiveModeHandlers) > 0 {
			_node := p.ActiveNode
			_node.JobHanderLocker.RLock()
			for _, v := range p.ActiveNode.LuaActiveModeHandlers {
				v.Handler(p.ActiveNode, v.Args...)
			}
			_node.JobHanderLocker.RUnlock()
		}
		if nodeDataActiveModer, ok := p.ActiveNode.Data.(NodeDataActiveModer); true == ok {
			nodeDataActiveModer.NodeDataActiveMode()
		}
	}
}
