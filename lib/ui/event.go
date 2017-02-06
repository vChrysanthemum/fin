package ui

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"
	"github.com/nsf/termbox-go"
)

func handleEvent(ev termbox.Event) (isContinueLoop, isRerender bool) {
	switch ev.Type {
	case termbox.EventKey:
		return handleEventKey(ev), false
	case termbox.EventResize:
		if nil != GCurrentRenderPage {
			return true, true
		}
		return true, false
	}

	return true, false
}

func handleEventKey(ev termbox.Event) (isContinueLoop bool) {
	isContinueLoop = true
	keyStr := evt2KeyStr(ev)
	GCurrentRenderPage.KeyPressHandleLocker.Lock()
	defer GCurrentRenderPage.KeyPressHandleLocker.Unlock()

	if "C-c" == keyStr {
		termui.Close()
		isContinueLoop = false
		return
	}

	if nil != GCurrentRenderPage.ActiveNode {
		isExecNormalKeyPressWork := true
		if nil != GCurrentRenderPage.ActiveNode.KeyPress {
			isExecNormalKeyPressWork = GCurrentRenderPage.ActiveNode.KeyPress(keyStr)
		}

		// 关于 ActiveNode 的一般性操作
		if true == isExecNormalKeyPressWork && nil != GCurrentRenderPage.ActiveNode {
			if "<escape>" == keyStr {
				GCurrentRenderPage.ActiveNode.QuitActiveMode()
				return
			}
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

			} else if true == IsVimKeyPressRight(keyStr) {
				if nil != node.FocusRightNode {
					GCurrentRenderPage.FocusNode = node.FocusRightNode
				}

			} else if true == IsVimKeyPressLeft(keyStr) {
				if nil != node.FocusLeftNode {
					GCurrentRenderPage.FocusNode = node.FocusLeftNode
				}

			} else if true == IsVimKeyPressDown(keyStr) {
				if nil != node.FocusBottomNode {
					GCurrentRenderPage.FocusNode = node.FocusBottomNode
				}

			} else if true == IsVimKeyPressUp(keyStr) {
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

	return
}

func consumeMoreEvents() (isContinueLoop, isRerender bool) {
	for {
		select {
		case ev := <-GTermboxEvents:
			if true == isRerender {
				isContinueLoop, _ = handleEvent(ev)
			} else {
				isContinueLoop, isRerender = handleEvent(ev)
			}
			if false == isContinueLoop {
				return isContinueLoop, isRerender
			}

		default:
			return true, false
		}
	}
}

func MainLoop() {
	go func() {
		for {
			GTermboxEvents <- termbox.PollEvent()
		}
	}()

	var (
		isContinueLoop, isRerender bool
	)

	for {
		select {
		case ev := <-GTermboxEvents:
			isContinueLoop, isRerender = handleEvent(ev)
			if false == isContinueLoop {
				return
			}

			if true == isRerender {
				isContinueLoop, _ = consumeMoreEvents()
			} else {
				isContinueLoop, isRerender = consumeMoreEvents()
			}

			if false == isContinueLoop {
				return
			}

			if true == isRerender {
				uiClear(0, -1)
				termui.Body.Width = ev.Width
				GCurrentRenderPage.ReRender()
			}
		}
	}
}

func (p *Page) pushWorkingNode(node *Node) {
	if false == node.CheckIfDisplay() {
		return
	}
	if uiBuffer, ok := node.UIBuffer.(*extra.Tabpane); true == ok {
		if true == uiBuffer.IsHideMenu {
			return
		}
	}
	p.WorkingNodes.PushBack(node)
	p.FocusNode = p.WorkingNodes.Back()
}

func (p *Node) QuitActiveMode() {
	p.page.ActiveNodeAfterReRender = nil

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
	p.ActiveNodeAfterReRender = node
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
