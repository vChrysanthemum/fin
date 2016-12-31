package ui

import (
	"bufio"
	"fin/ui/utils"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gizak/termui"
)

type ClearScreenBuffer struct {
	Buf termui.Buffer
}

func NewClearScreenBuffer() *ClearScreenBuffer {
	buf := termui.NewBuffer()
	buf.SetArea(image.Rectangle{
		image.Point{0, 0},
		image.Point{termui.TermWidth(), termui.TermHeight()},
	})
	buf.Fill(' ', utils.COLOR_DEFAULT, utils.COLOR_DEFAULT)
	return &ClearScreenBuffer{
		Buf: buf,
	}
}

func (p *ClearScreenBuffer) Buffer() termui.Buffer {
	return p.Buf
}

func (p *ClearScreenBuffer) RefreshArea() {
	p.Buf.SetArea(image.Rectangle{
		image.Point{0, 0},
		image.Point{termui.TermWidth() - 1, termui.TermHeight() - 1},
	})
}

func (p *Page) dumpNodesHtmlData(node *Node) {
	log.Println(node.HtmlData)
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		p.dumpNodesHtmlData(childNode)
	}
}

func (p *Page) DumpNodesHtmlData() {
	p.dumpNodesHtmlData(p.FirstChildNode)
}

func (p *Page) RemoveNode(node *Node) {
	if nodeDataOnRemover, ok := node.Data.(NodeDataOnRemover); true == ok {
		nodeDataOnRemover.NodeDataOnRemove()
	}

	delete(p.IdToNodeMap, node.Id)

	if nil != node.PrevSibling {
		node.PrevSibling.NextSibling = node.NextSibling
	}

	if nil != node.NextSibling {
		node.NextSibling.PrevSibling = node.PrevSibling
	}

	if nil != node.Parent {
		node.Parent.ChildrenCount -= 1
		if node.Parent.FirstChild == node {
			node.Parent.FirstChild = node.NextSibling
		}
		if node.Parent.LastChild == node {
			node.Parent.LastChild = node.PrevSibling
		}
	}

	p.ReRender()
}

func GetFileContent(path string) ([]byte, error) {
	path = filepath.Join(GlobalOption.ProjectPath, path)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(bufio.NewReader(file))
}

func IsVimKeyPressUp(keyStr string) bool {
	if "k" == keyStr || "<up>" == keyStr {
		return true
	} else {
		return false
	}
}

func IsVimKeyPressDown(keyStr string) bool {
	if "j" == keyStr || "<down>" == keyStr {
		return true
	} else {
		return false
	}
}

func IsVimKeyPressLeft(keyStr string) bool {
	if "h" == keyStr || "<left>" == keyStr {
		return true
	} else {
		return false
	}
}

func IsVimKeyPressRight(keyStr string) bool {
	if "l" == keyStr || "<right>" == keyStr {
		return true
	} else {
		return false
	}
}
