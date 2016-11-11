package ui

import (
	"in/ui/editor"

	"github.com/gizak/termui"
)

type NodeTerminal struct {
	*Node
	*editor.Editor
	ActiveModeBorderColor  termui.Attribute
	CommandPrefix          string
	NewCommand             *editor.Line
	CommandLines           []*editor.Line
	WaitKeyPressEnterChans []chan bool
}

func (p *Node) InitNodeTerminal() *NodeTerminal {
	nodeTerminal := new(NodeTerminal)
	nodeTerminal.Node = p
	nodeTerminal.Editor = editor.NewEditor()
	nodeTerminal.ActiveModeBorderColor = COLOR_ACTIVE_MODE_BORDERFG
	nodeTerminal.CommandPrefix = "> "
	nodeTerminal.PrepareNewCommand()

	p.Data = nodeTerminal
	p.Border = false
	p.KeyPress = nodeTerminal.KeyPress
	p.OnKeyPressEnter = nodeTerminal.OnKeyPressEnter
	p.FocusMode = nodeTerminal.FocusMode
	p.UnFocusMode = nodeTerminal.UnFocusMode
	p.ActiveMode = nodeTerminal.ActiveMode
	p.UnActiveMode = nodeTerminal.UnActiveMode

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
		}

		if len(p.WaitKeyPressEnterChans) > 0 {
			for _, c := range p.WaitKeyPressEnterChans {
				c <- true
				close(c)
			}
			p.WaitKeyPressEnterChans = make([]chan bool, 0)
		}

		p.Editor.WriteNewLine("")
		return
	}

	// 禁止删除一行
	if "C-8" == keyStr && (nil == p.CurrentLine || len(p.CurrentLine.Data) <= len(p.CommandPrefix)) {
		Beep()
		p.Editor.ResetCursor()
		return
	}

	p.Editor.Write(keyStr)
	uiRender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeTerminal) OnKeyPressEnter() {
	c := make(chan bool, 0)
	p.WaitKeyPressEnterChans = append(p.WaitKeyPressEnterChans, c)
	<-c
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

func (p *NodeTerminal) FocusMode() {
	p.Node.uiBuffer.(*editor.Editor).Border = true
	p.Node.uiBuffer.(*editor.Editor).BorderFg = COLOR_FOCUS_MODE_BORDERFG
	uiRender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeTerminal) UnFocusMode() {
	p.Node.uiBuffer.(*editor.Editor).Border = p.Node.Border
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	uiRender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeTerminal) ActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.ActiveModeBorderColor
	p.Node.uiBuffer.(*editor.Editor).ActiveMode()
	uiRender(p.Node.uiBuffer.(termui.Bufferer))
}

func (p *NodeTerminal) UnActiveMode() {
	p.Node.uiBuffer.(*editor.Editor).BorderFg = p.Node.BorderFg
	p.Node.uiBuffer.(*editor.Editor).UnActiveMode()
	uiRender(p.Node.uiBuffer.(termui.Bufferer))
}
