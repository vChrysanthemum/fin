package ui

import (
	"fin/ui/utils"
	"strconv"

	"github.com/gizak/termui"
	"github.com/gizak/termui/extra"

	"golang.org/x/net/html"
)

func (p *Node) ParseAttribute(attr []html.Attribute) (isUIChange, isNeedReRenderPage bool) {
	isUIChange = false
	isNeedReRenderPage = false

	if nil != p.UIBlock {
		p.UIBlock.BorderLabelFg = ColorDefaultBorderLabelFg
		p.UIBlock.BorderFg = ColorDefaultBorderFg

		for _, v := range attr {
			p.HTMLAttribute[v.Key] = v
			switch v.Key {
			case "position":
				if "absolute" == v.Val {
					p.Position = v.Val
				}

			case "top":
				tmp, err := strconv.Atoi(v.Val)
				if nil == err {
					p.UIBlock.Y = tmp
					p.isSettedPositionY = true
				}
			case "left":
				tmp, err := strconv.Atoi(v.Val)
				if nil == err {
					p.UIBlock.X = tmp
					p.isSettedPositionX = true
				}

			case "colorfg":
				p.ColorFg = v.Val

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
				p.UIBlock.BorderLabelFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultBorderLabelFg)

			case "borderlabel":
				isUIChange = true
				p.UIBlock.BorderLabel = v.Val

			case "borderfg":
				isUIChange = true
				p.UIBlock.BorderFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultBorderFg)

			case "border":
				isUIChange = true
				p.UIBlock.Border = utils.StringToBool(v.Val, p.UIBlock.Border)

			case "borderleft":
				isUIChange = true
				p.UIBlock.BorderLeft = utils.StringToBool(v.Val, p.UIBlock.BorderLeft)

			case "borderright":
				isUIChange = true
				p.UIBlock.BorderRight = utils.StringToBool(v.Val, p.UIBlock.BorderRight)

			case "bordertop":
				isUIChange = true
				p.UIBlock.BorderTop = utils.StringToBool(v.Val, p.UIBlock.BorderTop)

			case "borderbottom":
				isUIChange = true
				p.UIBlock.BorderBottom = utils.StringToBool(v.Val, p.UIBlock.BorderBottom)

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
				if nil != p.UIBuffer {
					if uiBuffer, ok := p.UIBuffer.(*extra.Tabpane); true == ok {
						isUIChange = true
						uiBuffer.TabFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabFg)
					}
				}

			case "tabbg":
				if nil != p.UIBuffer {
					if uiBuffer, ok := p.UIBuffer.(*extra.Tabpane); true == ok {
						isUIChange = true
						uiBuffer.TabBg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultTabBg)
					}
				}

			case "activetabfg":
				if nil != p.UIBuffer {
					if uiBuffer, ok := p.UIBuffer.(*extra.Tabpane); true == ok {
						isUIChange = true
						uiBuffer.ActiveTabFg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultActiveTabFg)
					}
				}

			case "activetabbg":
				if nil != p.UIBuffer {
					if uiBuffer, ok := p.UIBuffer.(*extra.Tabpane); true == ok {
						isUIChange = true
						uiBuffer.ActiveTabBg = utils.ColorToTermuiAttribute(v.Val, ColorDefaultActiveTabBg)
					}
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
