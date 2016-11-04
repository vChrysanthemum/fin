package ui

import (
	"fmt"
	"inn/ui"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHtmlScript(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	ui.Init(ui.Option{LuaResBaseDir: filepath.Join(os.Args[1], "lua")})
	script := strings.Trim(`
	function test()
		print(Node("head"):HtmlData())
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
	assert.Equal(t, page.GetLuaDocs(0), script)
}
