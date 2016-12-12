package ui

import "github.com/gizak/termui"

func registerHandles() {
	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		GCurrentRenderPage.Rerender()
	})

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.DefaultEvtStream.ResetHandlers()
		termui.StopLoop()
		termui.Close()
	})

	termui.Handle("/sys/kbd", func(e termui.Event) {
		GCurrentRenderPage.KeyPressHandleLocker.Lock()
		defer GCurrentRenderPage.KeyPressHandleLocker.Unlock()
		keyStr := e.Data.(termui.EvtKbd).KeyStr

		if nil != GCurrentRenderPage.ActiveNode {
			if nil != GCurrentRenderPage.ActiveNode.KeyPress {
				GCurrentRenderPage.ActiveNode.KeyPress(e)
			}

			// GCurrentRenderPage.ActiveNode.KeyPress后，GCurrentRenderPage.ActiveNode有可能为nil
			if nil == GCurrentRenderPage.ActiveNode {
				return
			}

			if len(GCurrentRenderPage.ActiveNode.KeyPressHandlers) > 0 {
				for _, v := range GCurrentRenderPage.ActiveNode.KeyPressHandlers {
					v.Args = append(v.Args, e)
					v.Handler(GCurrentRenderPage.ActiveNode, v.Args...)
				}
			}

			if "<escape>" == keyStr {
				GCurrentRenderPage.ActiveNode.QuitActiveMode()
				return
			}

			return
		}

		// 切换ActiveNode
		if "<tab>" == keyStr ||
			"<up>" == keyStr || "<down>" == keyStr ||
			"<left>" == keyStr || "<right>" == keyStr ||
			"h" == keyStr || "j" == keyStr || "k" == keyStr || "l" == keyStr {

			if nil != GCurrentRenderPage.FocusNode {
				nodeDataUnFocusModer, ok := GCurrentRenderPage.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer)
				if true == ok {
					nodeDataUnFocusModer.NodeDataUnFocusMode()
				}
			}

			if nil == GCurrentRenderPage.FocusNode {
				GCurrentRenderPage.FocusNode = GCurrentRenderPage.WorkingNodes.Front()
			} else {
				node := GCurrentRenderPage.FocusNode.Value.(*Node)
				if "<tab>" == keyStr {
					if nil != GCurrentRenderPage.FocusNode.Next() {
						GCurrentRenderPage.FocusNode = GCurrentRenderPage.FocusNode.Next()
					} else {
						GCurrentRenderPage.FocusNode = GCurrentRenderPage.WorkingNodes.Front()
					}

				} else if "<right>" == keyStr || "l" == keyStr {
					if nil != GCurrentRenderPage.FocusNode.Next() {
						GCurrentRenderPage.FocusNode = GCurrentRenderPage.FocusNode.Next()
					}

				} else if "<left>" == keyStr || "h" == keyStr {
					// "<left>" == keyStr
					if nil != GCurrentRenderPage.FocusNode.Prev() {
						GCurrentRenderPage.FocusNode = GCurrentRenderPage.FocusNode.Prev()
					}

				} else if "<down>" == keyStr || "j" == keyStr {
					if nil != node.FocusBottomNode {
						GCurrentRenderPage.FocusNode = node.FocusBottomNode
					}

				} else if "<up>" == keyStr || "k" == keyStr {
					if nil != node.FocusTopNode {
						GCurrentRenderPage.FocusNode = node.FocusTopNode
					}
				}
			}

			if nil != GCurrentRenderPage.FocusNode {
				nodeDataFocusModer, ok := GCurrentRenderPage.FocusNode.Value.(*Node).Data.(NodeDataFocusModer)
				if true == ok {
					nodeDataFocusModer.NodeDataFocusMode()
				}
			}

		} else if "<enter>" == keyStr {
			// 确认ActiveNode
			if nil != GCurrentRenderPage.FocusNode {
				nodeDataUnFocusModer, ok := GCurrentRenderPage.FocusNode.Value.(*Node).Data.(NodeDataUnFocusModer)
				if true == ok {
					nodeDataUnFocusModer.NodeDataUnFocusMode()
				}

				GCurrentRenderPage.SetActiveNode(GCurrentRenderPage.FocusNode.Value.(*Node))
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

func (p *Page) ClearActiveNode() {
	if nil != p.ActiveNode {
		if nodeDataUnActiveModer, ok := p.ActiveNode.Data.(NodeDataUnActiveModer); true == ok {
			nodeDataUnActiveModer.NodeDataUnActiveMode()
		}
	}
	p.ActiveNode = nil
}

func (p *Page) SetActiveNode(node *Node) {
	p.ClearActiveNode()
	p.ActiveNodeAfterRerender = node
	p.ActiveNode = node
	if nil != p.ActiveNode {
		if len(p.ActiveNode.LuaActiveModeHandlers) > 0 {
			for _, v := range p.ActiveNode.LuaActiveModeHandlers {
				v.Handler(p.ActiveNode, v.Args...)
			}
		}
		if nodeDataActiveModer, ok := p.ActiveNode.Data.(NodeDataActiveModer); true == ok {
			nodeDataActiveModer.NodeDataActiveMode()
		}
	}
}
