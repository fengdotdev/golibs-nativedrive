package nelement

func RootElement() *NElement {
	return &NElement{
		children: nil,
		parent:   nil,
		name:     "root",
		root:     nil,
		isRoot:   true,
	}
}

func NewFileElement(name string, parent *NElement) *NElement {
	return &NElement{
		children: nil,
		parent:   parent,
		name:     name,
		root:     parent.root,
		isRoot:   false,
	}
}
