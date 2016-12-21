package ui

import (
	uiutils "fin/ui/utils"
	"strconv"

	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"

	"golang.org/x/net/html"
)

func (p *Node) ParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	if nil == p.UIBlock {
		return
	}
	p.UIBlock.BorderLabelFg = COLOR_DEFAULT_BORDER_LABEL_FG
	p.UIBlock.BorderFg = COLOR_DEFAULT_BORDER_FG

	for _, v := range attr {
		p.HtmlAttribute[v.Key] = v
		switch v.Key {
		case "float":
			isUIChange = true
			switch v.Val {
			case "left":
				p.UIBlock.Float = termui.AlignLeft
			case "right":
				p.UIBlock.Float = termui.AlignRight
			case "top":
				p.UIBlock.Float = termui.AlignTop
			case "bottom":
				p.UIBlock.Float = termui.AlignBottom
			case "centervertical":
				p.UIBlock.Float = termui.AlignCenterVertical
			case "centerhorizontal":
				p.UIBlock.Float = termui.AlignCenterHorizontal
			case "center":
				p.UIBlock.Float = termui.AlignCenter
			}

		case "display":
			isUIChange = true
			if "none" == v.Val {
				*p.Display = false
			} else {
				*p.Display = true
			}

		case "paddingtop":
			isUIChange = true
			p.UIBlock.PaddingTop, _ = strconv.Atoi(v.Val)

		case "paddingbottom":
			isUIChange = true
			p.UIBlock.PaddingBottom, _ = strconv.Atoi(v.Val)

		case "paddingleft":
			isUIChange = true
			p.UIBlock.PaddingLeft, _ = strconv.Atoi(v.Val)

		case "paddingright":
			isUIChange = true
			p.UIBlock.PaddingRight, _ = strconv.Atoi(v.Val)

		case "borderlabelfg":
			isUIChange = true
			p.UIBlock.BorderLabelFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_BORDER_LABEL_FG)

		case "borderlabel":
			isUIChange = true
			p.UIBlock.BorderLabel = v.Val

		case "borderfg":
			isUIChange = true
			p.UIBlock.BorderFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_BORDER_FG)

		case "border":
			isUIChange = true
			p.UIBlock.Border = uiutils.StringToBool(v.Val, p.UIBlock.Border)

		case "borderleft":
			isUIChange = true
			p.UIBlock.BorderLeft = uiutils.StringToBool(v.Val, p.UIBlock.BorderLeft)

		case "borderright":
			isUIChange = true
			p.UIBlock.BorderRight = uiutils.StringToBool(v.Val, p.UIBlock.BorderRight)

		case "bordertop":
			isUIChange = true
			p.UIBlock.BorderTop = uiutils.StringToBool(v.Val, p.UIBlock.BorderTop)

		case "borderbottom":
			isUIChange = true
			p.UIBlock.BorderBottom = uiutils.StringToBool(v.Val, p.UIBlock.BorderBottom)

		case "height":
			isUIChange = true
			isNeedReRenderPage = true
			p.UIBlock.Height, _ = strconv.Atoi(v.Val)
			if p.UIBlock.Height < 0 {
				p.UIBlock.Height = 0
			}
			p.isShouldCalculateHeight = false

		case "width":
			isUIChange = true
			isNeedReRenderPage = true
			p.UIBlock.Width, _ = strconv.Atoi(v.Val)
			if p.UIBlock.Width < 0 {
				p.UIBlock.Width = 0
			}
			p.isShouldCalculateWidth = false

		case "tabfg":
			if nil != p.uiBuffer {
				if uiBuffer, ok := p.uiBuffer.(*extra.Tabpane); true == ok {
					isUIChange = true
					uiBuffer.TabFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TAB_FG)
				}
			}

		case "tabbg":
			if nil != p.uiBuffer {
				if uiBuffer, ok := p.uiBuffer.(*extra.Tabpane); true == ok {
					isUIChange = true
					uiBuffer.TabBg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_TAB_BG)
				}
			}

		case "activetabfg":
			if nil != p.uiBuffer {
				if uiBuffer, ok := p.uiBuffer.(*extra.Tabpane); true == ok {
					isUIChange = true
					uiBuffer.ActiveTabFg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_ACTIVE_TAB_FG)
				}
			}

		case "activetabbg":
			if nil != p.uiBuffer {
				if uiBuffer, ok := p.uiBuffer.(*extra.Tabpane); true == ok {
					isUIChange = true
					uiBuffer.ActiveTabBg = uiutils.ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_ACTIVE_TAB_BG)
				}
			}
		}

	}

	if nodeDataParseAttributer, ok := p.Data.(NodeDataParseAttributer); true == ok {
		_isUIChange, _isNeedReRenderPage := nodeDataParseAttributer.NodeDataParseAttribute(attr)
		if true == (isUIChange || _isUIChange) {
			isUIChange = true
		}
		if true == (isNeedReRenderPage || _isNeedReRenderPage) {
			isNeedReRenderPage = true
		}
	}

	return
}
