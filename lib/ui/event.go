package ui

import "github.com/gizak/termui"

func (p *Page) registerHandles() {
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})

	termui.Handle("/sys/kbd", func(e termui.Event) {
		if nil != p.ActiveNode && nil != p.ActiveNode.KeyPress {
			p.ActiveNode.KeyPress(e)
			return
		}

		keyStr := e.Data.(termui.EvtKbd).KeyStr

		var nodeFocus *Node

		if "<tab>" == keyStr {
			if nil != p.FocusNode {
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.UnFocusMode {
					nodeFocus.UnFocusMode()
				}
			}

			if nil == p.FocusNode {
				p.FocusNode = p.WorkingNodes.Front()
			} else {
				if nil != p.FocusNode.Next() {
					p.FocusNode = p.FocusNode.Next()
				} else {
					p.FocusNode = p.WorkingNodes.Front()
				}
			}

			if nil != p.FocusNode {
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.FocusMode {
					nodeFocus.FocusMode()
				}
			}
		}

		if "<enter>" == keyStr {
			if nil != p.FocusNode {
				nodeFocus = p.FocusNode.Value.(*Node)
				if nil != nodeFocus.FocusMode {
					nodeFocus.UnFocusMode()
				}

				p.ActiveNode = p.FocusNode.Value.(*Node)
			}
		}
	})
}

func (p *Page) pushWorkingNode(node *Node) {
	p.WorkingNodes.PushBack(node)
	p.FocusNode = p.WorkingNodes.Back()
	p.ActiveNode = node
}

func (p *Node) quitActiveMode() {
	if nil != p.page.FocusNode {
		nodeFocus := p.page.FocusNode.Value.(*Node)
		if nil != nodeFocus.FocusMode {
			nodeFocus.UnFocusMode()
		}
	}
	p.page.ActiveNode = nil
}
