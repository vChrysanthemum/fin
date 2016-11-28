package ui

import (
	"in/ui/editor"
	"in/ui/utils"

	"github.com/gizak/termui"
)

type NodeTerminal struct {
	*Node
	*editor.Editor
	ActiveModeBorderColor   termui.Attribute
	CommandPrefix           string
	NewCommand              *editor.Line
	CommandLines            []*editor.Line
	CurrentCommandLineIndex int
}

func (p *Node) InitNodeTerminal() *NodeTerminal {
	nodeTerminal := new(NodeTerminal)
	nodeTerminal.Node = p
	nodeTerminal.Editor = editor.NewEditor()
	nodeTerminal.ActiveModeBorderColor = COLOR_ACTIVE_MODE_BORDERFG
	nodeTerminal.CommandPrefix = "> "
	nodeTerminal.PrepareNewCommand()

	p.Data = nodeTerminal
	p.KeyPress = nodeTerminal.KeyPress

	p.uiBuffer = nodeTerminal.Editor
	p.UIBlock = &nodeTerminal.Editor.Block

	p.isShouldCalculateWidth = true
	p.isShouldCalculateHeight = false
	p.UIBlock.Border = true

	return nodeTerminal
}

func (p *NodeTerminal) KeyPress(e termui.Event) {
	keyStr := e.Data.(termui.EvtKbd).KeyStr
	if "<escape>" == keyStr {
		p.Node.QuitActiveMode()
		return
	}

	// 获取新的命令行
	if "<enter>" == keyStr {

		if (len(p.CommandLines) == 0 && nil != p.Editor.CurrentLine) ||
			(len(p.CommandLines) > 0 && p.CommandLines[len(p.CommandLines)-1] != p.Editor.CurrentLine) {

			p.NewCommand = p.Editor.CurrentLine
			p.CommandLines = append(p.CommandLines, p.NewCommand)
			p.CurrentCommandLineIndex = len(p.CommandLines) - 1
		}

		if len(p.Node.KeyPressEnterHandlers) > 0 {
			p.Node.JobHanderLocker.RLock()
			defer p.Node.JobHanderLocker.RUnlock()
			for _, v := range p.Node.KeyPressEnterHandlers {
				v.Handler(p.Node, v.Args...)
			}
		}

		p.PrepareNewCommand()
		p.Node.uiRender()
		return
	}

	// 禁止删除一行
	if "C-8" == keyStr && (nil == p.CurrentLine || len(p.CurrentLine.Data) <= len(p.CommandPrefix)) {
		utils.Beep()
		p.Editor.ResumeCursor()
		p.Node.uiRender()
		return
	}

	if "<up>" == keyStr || "<down>" == keyStr {
		if len(p.CommandLines) > 0 {
			if "<up>" == keyStr {
				p.CurrentCommandLineIndex -= 1
				if p.CurrentCommandLineIndex <= 0 {
					p.CurrentCommandLineIndex = 0
				}
			} else if "<down>" == keyStr {
				p.CurrentCommandLineIndex += 1
				if p.CurrentCommandLineIndex >= len(p.CommandLines) {
					p.CurrentCommandLineIndex = len(p.CommandLines)
				}
			}

			if len(p.CommandLines) == p.CurrentCommandLineIndex {
				p.Editor.UpdateCurrentLineData(p.CommandPrefix)
			} else {
				p.Editor.UpdateCurrentLineData(string(p.CommandLines[p.CurrentCommandLineIndex].Data))
			}
			p.Node.uiRender()
		}
		return
	}

	p.Editor.Write(keyStr)
	p.Node.uiRender()
}

func (p *NodeTerminal) PrepareNewCommand() {
	p.Editor.WriteNewLine(p.CommandPrefix)
}

func (p *NodeTerminal) PopNewCommand() (ret []byte) {
	if nil == p.NewCommand {
		return
	}

	ret = p.NewCommand.Data
	p.NewCommand = nil
	if len(p.CommandPrefix) > 0 {
		return ret[len(p.CommandPrefix):]
	} else {
		return ret
	}
}

func (p *NodeTerminal) WriteNewLine(line string) {
	p.Editor.WriteNewLine(line)
	p.Editor.CurrentLine = p.InitNewLine()
}

func (p *NodeTerminal) ClearLines() {
	p.NewCommand = nil
	p.CommandLines = make([]*editor.Line, 0)
	p.Editor.ClearLines()
}

func (p *NodeTerminal) NodeDataAfterRenderHandle() {
	p.Editor.AfterRenderHandle()
}

func (p *NodeTerminal) NodeDataFocusMode() {
	if false == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = true
		p.Node.tmpFocusModeBorder = p.Node.UIBlock.Border
		p.Node.tmpFocusModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.Border = true
		p.Node.UIBlock.BorderFg = COLOR_FOCUS_MODE_BORDERFG
		p.Node.uiRender()
	}
}

func (p *NodeTerminal) NodeDataUnFocusMode() {
	if true == p.Node.isCalledFocusMode {
		p.Node.isCalledFocusMode = false
		p.Node.UIBlock.Border = p.Node.tmpFocusModeBorder
		p.Node.UIBlock.BorderFg = p.Node.tmpFocusModeBorderFg
		p.Node.uiRender()
	}
}

func (p *NodeTerminal) NodeDataActiveMode() {
	if false == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = true
		p.Node.tmpActiveModeBorderFg = p.Node.UIBlock.BorderFg
		p.Node.UIBlock.BorderFg = COLOR_ACTIVE_MODE_BORDERFG
	}
	p.Editor.ActiveMode()
	p.Node.uiRender()
}

func (p *NodeTerminal) NodeDataUnActiveMode() {
	if true == p.Node.isCalledActiveMode {
		p.Node.isCalledActiveMode = false
		p.Node.UIBlock.BorderFg = p.Node.tmpActiveModeBorderFg
		p.Editor.UnActiveMode()
		p.Node.uiRender()
	}
}
