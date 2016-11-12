package ui

import "github.com/gizak/termui"

func (p *Page) registerHandles() {
	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		p.Rerender()
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.DefaultEvtStream.ResetHandlers()
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd", func(e termui.Event) {
		keyStr := e.Data.(termui.EvtKbd).KeyStr

		if nil != p.ActiveNode {
			if nil != p.ActiveNode.KeyPress {
				p.ActiveNode.KeyPress(e)
				return
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
			"<left>" == keyStr || "<right>" == keyStr {

			if nil != p.FocusNode {
				if nodeDataUnFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer); true == ok {
					nodeDataUnFocusModer.NodeDataUnFocusMode()
				}
			}

			if nil == p.FocusNode {
				p.FocusNode = p.WorkingNodes.Front()
			} else {
				if "<tab>" == keyStr || "<down>" == keyStr || "<right>" == keyStr {
					if nil != p.FocusNode.Next() {
						p.FocusNode = p.FocusNode.Next()
					} else {
						p.FocusNode = p.WorkingNodes.Front()
					}
				} else {
					// "<up>" == keyStr || "<left>" == keyStr
					if nil != p.FocusNode.Prev() {
						p.FocusNode = p.FocusNode.Prev()
					} else {
						p.FocusNode = p.WorkingNodes.Back()
					}
				}
			}

			if nil != p.FocusNode {
				if nodeDataFocusModer, ok := p.FocusNode.Value.(*Node).Data.(NodeDataFocusModer); true == ok {
					nodeDataFocusModer.NodeDataFocusMode()
				}
			}
		}

		// 确认ActiveNode
		if "<enter>" == keyStr {
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

	if nil != p.page.FocusNode {
		if nodeDataFocusModer, ok := p.page.FocusNode.Value.(*Node).Data.(NodeDataFocusModer); true == ok {
			nodeDataFocusModer.NodeDataFocusMode()
		}
	}

	p.page.ActiveNode = nil
}

func (p *Page) SetActiveNode(node *Node) {
	if nil != p.ActiveNode {
		if nodeDataUnActiveModer, ok := p.ActiveNode.Data.(NodeDataUnActiveModer); true == ok {
			nodeDataUnActiveModer.NodeDataUnActiveMode()
		}
	}
	p.ActiveNodeAfterRerender = node
	p.ActiveNode = node
	if nil != p.ActiveNode {
		if nodeDataActiveModer, ok := p.ActiveNode.Data.(NodeDataActiveModer); true == ok {
			nodeDataActiveModer.NodeDataActiveMode()
		}
	}
}
