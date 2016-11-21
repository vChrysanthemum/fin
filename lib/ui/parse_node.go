package ui

import (
	uiutils "in/ui/utils"
	"strconv"

	"golang.org/x/net/html"
)

func (p *Node) ParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool) {
	isUIChange = false
	isNeedRerenderPage = false

	if nil == p.UIBlock {
		return
	}
	p.UIBlock.BorderLabelFg = COLOR_DEFAULT_BORDER_LABEL_FG
	p.UIBlock.BorderFg = COLOR_DEFAULT_BORDER_FG

	for _, v := range attr {
		p.HtmlAttribute[v.Key] = v
		switch v.Key {
		case "ishide":
			isUIChange = true
			if "true" == v.Val {
				if false == p.isShouldHide {
					isNeedRerenderPage = true
					p.isShouldHide = true
				}
				//} else if "false" == v.Val {
			} else {
				if true == p.isShouldHide {
					isNeedRerenderPage = true
					p.isShouldHide = false
				}
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
			isNeedRerenderPage = true
			p.UIBlock.Height, _ = strconv.Atoi(v.Val)
			if p.UIBlock.Height < 0 {
				p.UIBlock.Height = 0
			}
			p.isShouldCalculateHeight = false

		case "width":
			isUIChange = true
			isNeedRerenderPage = true
			p.UIBlock.Width, _ = strconv.Atoi(v.Val)
			if p.UIBlock.Width < 0 {
				p.UIBlock.Width = 0
			}
			p.isShouldCalculateWidth = false
		}
	}

	if nodeDataParseAttributer, ok := p.Data.(NodeDataParseAttributer); true == ok {
		_isUIChange, _isNeedRerenderPage := nodeDataParseAttributer.NodeDataParseAttribute(attr)
		if true == (isUIChange || _isUIChange) {
			isUIChange = true
		}
		if true == (isNeedRerenderPage || _isNeedRerenderPage) {
			isNeedRerenderPage = true
		}
	}

	return
}
