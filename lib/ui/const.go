package ui

import (
	"fin/ui/utils"

	"github.com/gizak/termui"
)

const (
	TABLE_TR_COLS = 12
)

var (
	COLOR_BODY_DEFAULT_COLORFG                   = "blue"
	COLOR_SELECTED_OPTION_COLORFG                = "gray"
	COLOR_SELECTED_OPTION_COLORBG                = "blue"
	COLOR_FOCUS_MODE_BORDERFG                    = utils.COLOR_RED
	COLOR_ACTIVE_MODE_BORDERFG                   = utils.COLOR_BLUE
	COLOR_ACTIVE_MODE_BORDERBG                   = utils.COLOR_BLUE
	COLOR_DEFAULT_TEXT_COLOR_FG                  = utils.COLOR_BLACK
	COLOR_DEFAULT_BORDER_LABEL_FG                = utils.COLOR_BLUE
	COLOR_DEFAULT_BORDER_FG                      = utils.COLOR_GRAY
	COLOR_DEFAULT_GAUGE_BARCOLOR                 = utils.COLOR_BLUE
	COLOR_DEFAULT_GAUGE_PERCENTCOLOR             = utils.COLOR_BLUE
	COLOR_DEFAULT_GAUGE_PERCENTCOLOR_HIGHLIGHTED = utils.COLOR_GREEN
	COLOR_DEFAULT_TABPANE_FG                     = utils.COLOR_DEFAULT
	COLOR_DEFAULT_TABPANE_BG                     = utils.COLOR_BLACK
	COLOR_DEFAULT_TAB_FG                         = utils.COLOR_DEFAULT
	COLOR_DEFAULT_TAB_BG                         = utils.COLOR_DEFAULT
	COLOR_DEFAULT_ACTIVE_TAB_FG                  = utils.COLOR_DEFAULT
	COLOR_DEFAULT_ACTIVE_TAB_BG                  = utils.COLOR_DEFAULT
	COLOR_DEFAULT_LINE_CHART_AXES                = utils.COLOR_WHITE
	COLOR_DEFAULT_LINE_CHART_LINE                = utils.COLOR_BLUE | termui.AttrBold
)
