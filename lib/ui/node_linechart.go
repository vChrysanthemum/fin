package ui

import (
	uiutils "fin/ui/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gizak/termui"
)

type NodeLineChart struct {
	*Node
}

func (p *Node) InitNodeLineChart() {
	nodeLineChart := new(NodeLineChart)
	nodeLineChart.Node = p

	p.Data = nodeLineChart

	uiBuffer := termui.NewLineChart()
	p.uiBuffer = uiBuffer
	p.UIBlock = &uiBuffer.Block
	p.Display = &p.UIBlock.Display

	p.UIBlock.Border = true
	p.UIBlock.Height = 10
	p.UIBlock.Width = 10

	uiBuffer.Mode = "braille"
	uiBuffer.AxesColor = uiutils.COLOR_WHITE
	uiBuffer.LineColor = uiutils.COLOR_BLUE | termui.AttrBold

	return
}

func (p *NodeLineChart) NodeDataSetValue(content string) {
	uiBuffer := p.Node.uiBuffer.(*termui.LineChart)
	var (
		_arr []string
		arr  []float64
		_f64 float64
		err  error
	)
	_arr = strings.Split(content, ",")
	for _, v := range _arr {
		_f64, err = strconv.ParseFloat(strings.Trim(v, " "), 64)
		if nil == err {
			arr = append(arr, _f64)
		}
	}
	uiBuffer.Data = arr

	p.Node.uiRender()
	return
}

func (p *NodeLineChart) NodeDataGetValue() (string, bool) {
	return fmt.Sprintf("%v", p.Node.uiBuffer.(*termui.LineChart).Data), true
}
