package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

const (
	TableTrCols = 12

	ColorBodyDefaultColorFg                  = "blue"
	ColorSelectedOptionColorFg               = "gray"
	ColorSelectedOptionColorbg               = "blue"
	ColorFocusModeBorderFg                   = utils.ColorRed
	ColorActiveModeBorderFg                  = utils.ColorBlue
	ColorActiveModeBorderbg                  = utils.ColorBlue
	ColorDefaultTextColorFg                  = utils.ColorBlack
	ColorDefaultBorderLabelFg                = utils.ColorBlue
	ColorDefaultBorderFg                     = utils.ColorGray
	ColorDefaultGaugeBarcolor                = utils.ColorBlue
	ColorDefaultGaugePercentcolor            = utils.ColorBlue
	ColorDefaultGaugePercentColorHighlighted = utils.ColorGreen
	ColorDefaultTabpaneFg                    = utils.ColorDefault
	ColorDefaultTabpaneBg                    = utils.ColorBlack
	ColorDefaultTabFg                        = utils.ColorDefault
	ColorDefaultTabBg                        = utils.ColorDefault
	ColorDefaultActiveTabFg                  = utils.ColorDefault
	ColorDefaultActiveTabBg                  = utils.ColorDefault
	ColorDefaultLineChartAxes                = utils.ColorWhite
	ColorDefaultLineChartLine                = utils.ColorBlue | termui.AttrBold
)
