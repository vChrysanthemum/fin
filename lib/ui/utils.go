package ui

import (
	"bufio"
	"fin/ui/utils"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
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
	buf.Fill(' ', utils.ColorDefault, utils.ColorDefault)
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

func (p *Page) dumpNodesHTMLData(node *Node) {
	log.Println(node.HTMLData)
	for childNode := node.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		p.dumpNodesHTMLData(childNode)
	}
}

func (p *Page) DumpNodesHTMLData() {
	p.dumpNodesHTMLData(p.FirstChildNode)
}

func (p *Page) AppendNode(node *Node, content string) error {
	childNode, err := ParseNode(content)
	if nil != err {
		return err
	}

	idToNodeMap := childNode.ExtractChildsMapIDNodes()
	for k, v := range idToNodeMap {
		p.IDToNodeMap[k] = v
	}

	node.addChild(childNode)

	p.ReRender()

	return nil
}

func (p *Page) RemoveNode(node *Node) {
	if nodeDataOnRemover, ok := node.Data.(NodeDataOnRemover); true == ok {
		nodeDataOnRemover.NodeDataOnRemove()
	}

	delete(p.IDToNodeMap, node.ID)

	if nil != node.PrevSibling {
		node.PrevSibling.NextSibling = node.NextSibling
	}

	if nil != node.NextSibling {
		node.NextSibling.PrevSibling = node.PrevSibling
	}

	if nil != node.Parent {
		node.Parent.ChildrenCount--
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
	}
	return false
}

func IsVimKeyPressDown(keyStr string) bool {
	if "j" == keyStr || "<down>" == keyStr {
		return true
	}
	return false
}

func IsVimKeyPressLeft(keyStr string) bool {
	if "h" == keyStr || "<left>" == keyStr {
		return true
	}
	return false
}

func IsVimKeyPressRight(keyStr string) bool {
	if "l" == keyStr || "<right>" == keyStr {
		return true
	}
	return false
}

func evt2KeyStr(e termbox.Event) string {
	k := string(e.Ch)
	pre := ""
	mod := ""

	if e.Mod == termbox.ModAlt {
		mod = "M-"
	}
	if e.Ch == 0 {
		if e.Key > 0xFFFF-12 {
			k = "<f" + strconv.Itoa(0xFFFF-int(e.Key)+1) + ">"
		} else if e.Key > 0xFFFF-25 {
			ks := []string{"<insert>", "<delete>", "<home>", "<end>", "<previous>", "<next>", "<up>", "<down>", "<left>", "<right>"}
			k = ks[0xFFFF-int(e.Key)-12]
		}

		if e.Key <= 0x7F {
			pre = "C-"
			k = string('a' - 1 + int(e.Key))
			kmap := map[termbox.Key][2]string{
				termbox.KeyCtrlSpace:     {"C-", "<space>"},
				termbox.KeyBackspace:     {"", "<backspace>"},
				termbox.KeyTab:           {"", "<tab>"},
				termbox.KeyEnter:         {"", "<enter>"},
				termbox.KeyEsc:           {"", "<escape>"},
				termbox.KeyCtrlBackslash: {"C-", "\\"},
				termbox.KeyCtrlSlash:     {"C-", "/"},
				termbox.KeySpace:         {"", "<space>"},
				termbox.KeyCtrl8:         {"C-", "8"},
			}
			if sk, ok := kmap[e.Key]; ok {
				pre = sk[0]
				k = sk[1]
			}
		}
	}

	return pre + mod + k
}
