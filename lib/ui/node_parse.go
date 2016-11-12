package ui

import (
	. "in/ui/utils"
	"strconv"

	"golang.org/x/net/html"
)

func (p *Node) ParseAttribute(attr []html.Attribute) (isUIChange, isNeedRerenderPage bool) {
	isUIChange = false
	isNeedRerenderPage = false

	if nil == p.uiBlock {
		return
	}
	p.uiBlock.BorderLabelFg = COLOR_DEFAULT_BORDER_LABEL_FG
	p.uiBlock.BorderFg = COLOR_DEFAULT_BORDER_FG

	for _, v := range attr {
		switch v.Key {
		case "borderlabelfg":
			isUIChange = true
			p.uiBlock.BorderLabelFg = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_BORDER_LABEL_FG)

		case "borderlabel":
			isUIChange = true
			p.uiBlock.BorderLabel = v.Val

		case "borderfg":
			isUIChange = true
			p.uiBlock.BorderFg = ColorToTermuiAttribute(v.Val, COLOR_DEFAULT_BORDER_FG)

		case "border":
			isUIChange = true
			p.uiBlock.Border = StringToBool(v.Val, p.uiBlock.Border)

		case "borderleft":
			isUIChange = true
			p.uiBlock.BorderLeft = StringToBool(v.Val, p.uiBlock.BorderLeft)

		case "borderright":
			isUIChange = true
			p.uiBlock.BorderRight = StringToBool(v.Val, p.uiBlock.BorderRight)

		case "bordertop":
			isUIChange = true
			p.uiBlock.BorderTop = StringToBool(v.Val, p.uiBlock.BorderTop)

		case "borderbottom":
			isUIChange = true
			p.uiBlock.BorderBottom = StringToBool(v.Val, p.uiBlock.BorderBottom)

		case "height":
			isUIChange = true
			isNeedRerenderPage = true
			p.uiBlock.Height, _ = strconv.Atoi(v.Val)
			if p.uiBlock.Height < 0 {
				p.uiBlock.Height = 0
			}
			p.isShouldCalculateHeight = false

		case "width":
			isUIChange = true
			isNeedRerenderPage = true
			p.uiBlock.Width, _ = strconv.Atoi(v.Val)
			if p.uiBlock.Width < 0 {
				p.uiBlock.Width = 0
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
