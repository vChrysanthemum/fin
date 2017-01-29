package editor

import (
	uiutils "fin/ui/utils"
	"time"

	"github.com/gizak/termui"
	termbox "github.com/nsf/termbox-go"
)

func (p *Editor) handleKeyEvent(keyStr string) (isQuitActiveMode bool) {
	isQuitActiveMode = false

	if 0 == len(p.Lines) {
		p.EditModeAppendNewLine()
	}

	switch keyStr {
	case "<escape>":
		switch p.Mode {
		case EDITOR_NORMAL_MODE:
			isQuitActiveMode = true
			uiutils.UISetCursor(-1, -1)

		case EDITOR_EDIT_MODE:
			p.EditModeQuit()
			p.NormalModeEnter()

		case EDITOR_COMMAND_MODE:
			p.CommandModeQuit()
			p.NormalModeEnter()
		}

	default:
		switch p.Mode {
		case EDITOR_MODE_NONE:

		case EDITOR_EDIT_MODE:
			p.isShouldRefreshCommandModeBuf = true
			p.EditModeWrite(keyStr)

		case EDITOR_NORMAL_MODE:
			p.isShouldRefreshCommandModeBuf = true
			p.NormalModeWrite(keyStr)

		case EDITOR_COMMAND_MODE:
			p.CommandModeWrite(keyStr)
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
			p.UIRender()

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
