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

		if nil == p.FocusNode {
			p.FocusNode = p.WorkingNodes.Front()
		}

		if nil != p.FocusNode {
			keyStr := e.Data.(termui.EvtKbd).KeyStr
			if "<enter>" == keyStr {
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
