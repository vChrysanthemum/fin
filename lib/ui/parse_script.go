package ui

import "golang.org/x/net/html"

func (p *Page) parseHTMLScript(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = nil
	isFallthrough = false

	docType := ""
	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "type":
			docType = attr.Val
		}
	}
	if "text/lua" != docType {
		return
	}

	isReadFromFile := false
	for _, attr := range htmlNode.Attr {
		switch attr.Key {
		case "src":
			isReadFromFile = true
			p.AppendScript(ScriptDoc{
				DataType: "file",
				Data:     attr.Val,
			})
		}
	}

	if false == isReadFromFile && nil != htmlNode.FirstChild {
		p.AppendScript(ScriptDoc{
			DataType: "string",
			Data:     htmlNode.FirstChild.Data,
		})
	}

	return
}
