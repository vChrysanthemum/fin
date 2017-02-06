package ui

import (
	"fin/ui"
	"fmt"
	"log"
	"testing"

	"github.com/gizak/termui"
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

func TestParseID(t *testing.T) {
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

	assert.NotNil(t, page.IDToNodeMap[id])

	ul := page.IDToNodeMap[id].Data.(*ui.NodeSelect)
	assert.Equal(t, ul.Children[0].Data, title)
}

func TestParseNode(t *testing.T) {
	id2 := "child"
	title := "Hello World!"

	s2 := fmt.Sprintf(`
	<par id="%s">%s</par>
	`, id2, title)

	node, err := ParseNode(s2)
	assert.Nil(t, err)
	assert.Equal(t, node.HTMLData, "par")
	assert.Equal(t, node.ID, id2)
}

func TestAppendNode(t *testing.T) {
	id1 := "menu"
	id2 := "child"
	title := "Hello World!"

	s1 := fmt.Sprintf(`
	<html>
    <body>
        <div id="%s">
		</div>
	</body>
	</html>
	`, id1)
	page, err := ui.Parse(s1)
	assert.NotNil(t, page)
	assert.Nil(t, err)

	assert.NotNil(t, page.IDToNodeMap[id1])

	s2 := fmt.Sprintf(`
	<par id="%s">%s</par>
	`, id2, title)

	err = page.AppendNode(page.IDToNodeMap[id1], s2)
	assert.Nil(t, err)
	assert.Equal(t, page.IDToNodeMap[id2].HTMLData, "par")
	assert.Equal(t, page.IDToNodeMap[id2].UIBuffer.(*termui.Par).Text, title)
}
