package ui

type NodeModal struct {
	page *Page
}

func (p *Page) prepareModals() {
	p.Modals = make(map[string]*NodeModal)
}

func (p *Node) InitNodeModal(content string) error {
	nodeModal := new(NodeModal)
	p.Data = nodeModal

	page, err := Parse(content)
	if nil != err {
		return err
	}
	nodeModal.page = page

	err = page.Render()
	if nil != err {
		return err
	}

	nodeModal.page.MainPage = p.page

	return nil
}
