package ui

import (
	"fmt"
	"in/ui"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	title := "Hello World!"
	s := fmt.Sprintf(`
	<html>
	<head>
		<title> %s </title>
	</head>
	<body>
	</body>
	</html>
	`, title)
	page, err := ui.Parse(s)
	assert.NotNil(t, page)
	assert.Nil(t, err)
	assert.Equal(t, title, page.Title)
}

func TestParseHtmlBodySelect(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	BorderLabel := "测试"
	title := "Hello World!"
	s := fmt.Sprintf(`
	<html>
    <body>
        <select BorderLabel="%s" style="none">
			<option value="%s"> %s </option>
			<option value="%s"> %s </option>
		</select>
	</body>
	</html>
	`, BorderLabel, title, title, title, title)
	page, err := ui.Parse(s)
	assert.NotNil(t, page)
	assert.Nil(t, err)

	sum := 0
	var check func(node *ui.Node)
	check = func(node *ui.Node) {
		sum++
		if sum > 10 {
			t.Errorf("dead loop")
			return
		}
		if ul, ok := node.Data.(*ui.NodeSelect); true == ok {
			assert.Equal(t, BorderLabel, node.UIBlock.BorderLabel)

			for _, v := range ul.Children {
				assert.Equal(t, title, v.Data)
				assert.Equal(t, title, v.Value)
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			check(child)
		}
	}
	check(page.FirstChildNode)
}

func TestParseId(t *testing.T) {
	id := "menu"
	title := "Hello World!"
	s := fmt.Sprintf(`
	<html>
    <body>
        <select id="%s">
			<option>%s</option>
			<option></option>
		</select>
	</body>
	</html>
	`, id, title)

	page, err := ui.Parse(s)
	assert.NotNil(t, page)
	assert.Nil(t, err)

	assert.NotNil(t, page.IdToNodeMap[id])

	ul := page.IdToNodeMap[id].Data.(*ui.NodeSelect)
	assert.Equal(t, ul.Children[0].Data, title)
}
