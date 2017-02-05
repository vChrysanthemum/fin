package ui

import (
	"time"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
)

func (p *Editor) handleKeyEvent(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	if 0 == len(p.Lines) {
		p.EditModeAppendNewLine(p.EditModeCursor)
	}

	isQuitActiveMode = p.ActionGroup.Write(p.EditModeCursor, keyStr)

	if "<escape>" == keyStr || "<enter>" == keyStr {
		p.KeyEventsResultIsQuitActiveMode <- isQuitActiveMode
	}

	return
}

// return:
// 			bool		isQuitActiveMode
func (p *Editor) consumeMoreKeyEvents() bool {
	var keyStr string
	for {
		select {
		case keyStr = <-p.KeyEvents:
			if true == p.handleKeyEvent(keyStr) {
				return true
			}
		default:
			return false
		}
	}
}

func (p *Editor) RegisterKeyEventHandlers() {
	go func() {
		var keyStr string
		for {
			p.isShouldRefreshEditModeBuf = false
			p.isShouldRefreshCommandModeBuf = false

			select {
			case keyStr = <-p.KeyEvents:
				if true == p.handleKeyEvent(keyStr) {
					break
				}

				// 尽可能合并请求
				time.Sleep(time.Duration(4) * time.Microsecond)

				if true == p.consumeMoreKeyEvents() {
					break
				}
			}

			p.RefreshBuf()
			p.RefreshCursorByEditorLine()
			p.RefreshBuf()

			termui.RenderLock.Lock()
			termbox.Flush()
			termui.RenderLock.Unlock()
		}
	}()
}

func (p *Editor) Write(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false
	if "<escape>" == keyStr || "<enter>" == keyStr {
		p.KeyEvents <- keyStr
		isQuitActiveMode = <-p.KeyEventsResultIsQuitActiveMode
	} else {
		p.KeyEvents <- keyStr
	}
	return
}
