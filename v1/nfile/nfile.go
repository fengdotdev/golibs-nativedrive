package nfile

import "github.com/fengdotdev/golibs-nativedrive/ntrait"

var _ ntrait.File = (*NFile)(nil)

type NFile struct {
}

// GetContent implements trait.File.
func (n *NFile) GetContent() ([]byte, error) {
	panic("unimplemented")
}

// GetName implements trait.File.
func (n *NFile) GetName() string {
	panic("unimplemented")
}

// GetPath implements trait.File.
func (n *NFile) GetPath() string {
	panic("unimplemented")
}

// GetSize implements trait.File.
func (n *NFile) GetSize() int64 {
	panic("unimplemented")
}
