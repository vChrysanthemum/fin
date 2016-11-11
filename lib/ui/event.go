package ui

import "github.com/gizak/termui"

func (p *Page) registerHandles() {
	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		GClearScreenBuffer.RefreshArea()
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.DefaultEvtStream.ResetHandlers()
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd", func(e termui.Event) {
		if nil != p.ActiveNode && nil != p.ActiveNode.KeyPress {
			p.ActiveNode.KeyPress(e)
			return
		}

		keyStr := e.Data.(termui.EvtKbd).KeyStr

		var nodeFocus *Node

		// 切换ActiveNode
		if "<tab>" == keyStr ||
			"<up>" == keyStr || "<down>" == keyStr ||
			"<left>" == keyStr || "<right>" == keyStr {

			if nil != p.FocusNode {
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.UnFocusMode {
					nodeFocus.UnFocusMode()
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
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.FocusMode {
					nodeFocus.FocusMode()
				}
			}
		}

		// 确认ActiveNode
		if "<enter>" == keyStr {
			if nil != p.FocusNode {
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.FocusMode {
					nodeFocus.UnFocusMode()
				}

				p.SetActiveNode(p.FocusNode.Value.(*Node))
			}
		}
	})
}

func (p *Page) pushWorkingNode(node *Node) {
	p.WorkingNodes.PushBack(node)
	p.FocusNode = p.WorkingNodes.Back()
	p.NodeActiveAfterRender = node
}

func (p *Node) QuitActiveMode() {
	if nil != p.page.FocusNode {
		nodeFocus := p.page.FocusNode.Value.(*Node)
		if nil != nodeFocus.FocusMode {
			nodeFocus.FocusMode()
		}
	}
	p.page.ActiveNode = nil
}

func (p *Page) SetActiveNode(node *Node) {
	if nil != p.ActiveNode && p.ActiveNode != node && nil != p.ActiveNode.ActiveMode {
		p.ActiveNode.UnActiveMode()
	}
	p.ActiveNode = node
	if nil != p.ActiveNode && nil != p.ActiveNode.ActiveMode {
		p.ActiveNode.ActiveMode()
	}
}
