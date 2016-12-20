package ui

import (
	"fin/ui"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHtmlScript(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	ui.Init(ui.Option{ResBaseDir: os.Args[1]})
	script := strings.Trim(`
	function test()
		print(Node("head"):GetHtmlData())
		print('call lua success.')
	end
	test()
	--[[
	print(Node("head"):Data())
	]]
	`, " \r\n\t")

	s := fmt.Sprintf(`
	<html>
    <body>
	<div>
		<div id="head">
			<select>
				<option>
				测试1(fg-white,bg-green)
				</option>
				<option>
				测试2(fg-blue)
				</option>
			</select>
		</div>
		<script type="text/lua">%s</script>
	</div>
	</body>
	</html>
	`, script)

	page, err := ui.Parse(s)
	assert.NotNil(t, page)
	assert.Nil(t, err)
	assert.Equal(t, page.GetLuaDocs(0).Data, script)
}
