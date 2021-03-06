package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

type NodeTerminal struct {
	*Node
	*Terminal
	ActiveModeBorderColor   termui.Attribute
	CommandPrefix           string
	NewCommand              *TerminalLine
	CommandHistory          []string
	CurrentCommandLineIndex int
}

func (p *Node) InitNodeTerminal() {
	nodeTerminal := new(NodeTerminal)
	nodeTerminal.Node = p
	nodeTerminal.Terminal = NewTerminal()
	nodeTerminal.ActiveModeBorderColor = ColorActiveModeBorderFg
	nodeTerminal.CommandPrefix = "> "
	nodeTerminal.PrepareNewCommand()

	p.Data = nodeTerminal
	p.KeyPress = nodeTerminal.KeyPress

	p.UIBuffer = nodeTerminal.Terminal
	p.UIBlock = &nodeTerminal.Terminal.Block
	p.Display = &p.UIBlock.Display

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Border = true

	p.isWorkNode = true

	return
}

func (p *NodeTerminal) KeyPress(keyStr string) (isExecNormalKeyPressWork bool) {
	isExecNormalKeyPressWork = true
	defer func() {
		if len(p.Node.KeyPressHandlers) > 0 {
			for _, v := range p.Node.KeyPressHandlers {
				v.Args = append(v.Args, keyStr)
				v.Handler(p.Node, v.Args...)
			}
		}
	}()

	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
		return
	}

	// 禁止删除一行
	if "C-8" == keyStr && (nil == p.Terminal.Cursor.Line || len(p.Terminal.Cursor.Line.Data) <= len(p.CommandPrefix)) {
		utils.Beep()
		p.Terminal.Cursor.ResumeCursor()
		p.Node.UIRender()
		return
	}

	if "<left>" == keyStr || "<right>" == keyStr {
		return
	}

	if "<up>" == keyStr || "<down>" == keyStr {
		if len(p.CommandHistory) > 0 {
			if "<up>" == keyStr {
				p.CurrentCommandLineIndex--
				if p.CurrentCommandLineIndex <= 0 {
					p.CurrentCommandLineIndex = 0
				}
			} else if "<down>" == keyStr {
				p.CurrentCommandLineIndex++
				if p.CurrentCommandLineIndex >= len(p.CommandHistory) {
					p.CurrentCommandLineIndex = len(p.CommandHistory)
				}
			}

			if len(p.CommandHistory) <= p.CurrentCommandLineIndex {
				p.Terminal.UpdateCursorLineData(p.CommandPrefix)
			} else {
				p.Terminal.UpdateCursorLineData(p.CommandPrefix + p.CommandHistory[p.CurrentCommandLineIndex])
			}

			p.Node.UIRender()
		}
		return
	}

	if "C-c" == keyStr {
		p.Terminal.UpdateCursorLineData(p.CommandPrefix)
		p.Node.UIRender()
		return
	}

	// 获取新的命令行
	if "<enter>" == keyStr {
		if nil != p.Terminal.Cursor.Line {
			p.NewCommand = p.Terminal.Cursor.Line
			if nil != p.NewCommand &&
				nil != p.NewCommand.Data &&
				len(p.NewCommand.Data) > len(p.CommandPrefix) &&
				"" != string(p.NewCommand.Data[len(p.CommandPrefix):]) {
				p.CommandHistory = append(p.CommandHistory, string(p.NewCommand.Data[len(p.CommandPrefix):]))
			}
			p.CurrentCommandLineIndex = len(p.CommandHistory)
		}

		if len(p.Node.KeyPressEnterHandlers) > 0 {
			for _, v := range p.Node.KeyPressEnterHandlers {
				v.Handler(p.Node, v.Args...)
			}
		}

		p.PrepareNewCommand()
		p.Node.UIRender()
		return
	}

	p.Terminal.Write(keyStr)
	p.Node.UIRender()
	return
}

func (p *NodeTerminal) PrepareNewCommand() {
	p.Terminal.WriteNewLine(p.CommandPrefix)
}

func (p *NodeTerminal) PopNewCommand() (ret []byte) {
	if nil == p.NewCommand {
		return
	}

	ret = p.NewCommand.Data
	p.NewCommand = nil
	if len(p.CommandPrefix) > 0 {
		return ret[len(p.CommandPrefix):]
	}
	return ret
}

func (p *NodeTerminal) WriteString(data string) {
	p.Terminal.Cursor.Line.Write(data)
}

func (p *NodeTerminal) WriteNewLine(line string) {
	p.Terminal.WriteNewLine(line)
	p.Terminal.Cursor.Line = p.InitNewLine()
}

func (p *NodeTerminal) ClearLines() {
	p.NewCommand = nil
	p.Terminal.ClearLines()
}

func (p *NodeTerminal) ClearCommandHistory() {
	p.CommandHistory = make([]string, 0)
}

func (p *NodeTerminal) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = ColorFocusModeBorderFg
		p.Node.UIRender()
	}
}

func (p *NodeTerminal) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.UIRender()
	}
}

func (p *NodeTerminal) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = ColorActiveModeBorderFg
	}
	p.Terminal.ActiveMode()
	p.Node.UIRender()
}

func (p *NodeTerminal) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode && true == p.Node.UIBlock.Border {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Terminal.UnActiveMode()
		p.Node.UIRender()
	}
}
