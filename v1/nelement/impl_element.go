package nelement

import "github.com/fengdotdev/golibs-nativedrive/interfaces"

var _ interfaces.Element = (*NElement)(nil)

// GetChildren implements interfaces.Element.
func (n *NElement) GetChildren() []interfaces.Element {
	panic("unimplemented")
}

// GetName implements interfaces.Element.
func (n *NElement) GetName() string {
	panic("unimplemented")
}

// GetParent implements interfaces.Element.
func (n *NElement) GetParent() interfaces.Element {
	panic("unimplemented")
}

// GetPath implements interfaces.Element.
func (n *NElement) GetPath() string {
	panic("unimplemented")
}

// GetSize implements interfaces.Element.
func (n *NElement) GetSize() int64 {
	panic("unimplemented")
}

// IsFile implements interfaces.Element.
func (n *NElement) IsFile() bool {
	panic("unimplemented")
}

// IsFolder implements interfaces.Element.
func (n *NElement) IsFolder() bool {
	panic("unimplemented")
}

// Type implements interfaces.Element.
func (n *NElement) Type() interfaces.ElementType {
	panic("unimplemented")
}

// IsRoot implements interfaces.Element.
func (n *NElement) IsRoot() bool {
	panic("unimplemented")
}
