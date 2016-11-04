package ui

import "golang.org/x/net/html"

func (p *Page) parseHtmlScript(parentNode *Node, htmlNode *html.Node) (ret *Node, isFallthrough bool) {
	ret = nil
	isFallthrough = false

	if nil == htmlNode.FirstChild {
		return
	}

	doc := htmlNode.FirstChild.Data
	docType := ""
	for _, attr := range htmlNode.Attr {
		if "type" == attr.Key {
			docType = attr.Val
		}
	}

	p.AppendScript(doc, docType)

	return
}
