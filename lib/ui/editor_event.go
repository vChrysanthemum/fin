package ui

import (
	"fin/ui/utils"
	"time"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

func (p *Editor) handleKeyEvent(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	switch keyStr {
	case "<escape>":
		p.ActionGroup.makeStatePrepareWrite()

		switch p.EditorView.Mode {
		case EditorCommandMode:
			isQuitActiveMode = true
			utils.UISetCursor(-1, -1)

		case EditorInputMode:
			p.CommandModeEnter(p.InputModeCursor)

		case EditorLastLineMode:
			p.LastLineModeQuit()
			p.CommandModeEnter(p.InputModeCursor)
		}

		p.isShouldRefreshLastLineModeBuf = true

	default:
		switch p.Mode {
		case EditorLastLineMode:
			p.LastLineModeWrite(p.InputModeCursor, p.LastLineModeCursor, keyStr)
		default:
			if len(p.LastLineModeBuf.Data) > 0 {
				p.LastLineModeBuf.CleanData(p.LastLineModeCursor.EditorCursor)
			}
			p.ActionGroup.Write(p.InputModeCursor, keyStr)
		}

	}

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
			p.isShouldRefreshInputModeBuf = false
			p.isShouldRefreshLastLineModeBuf = false

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
