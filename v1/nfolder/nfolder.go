package nfolder

import (
	"github.com/fengdotdev/golibs-nativedrive/ntrait"
)

var _ ntrait.Folder = (*NFolder)(nil)




type NFolder struct{}

// GetList implements trait.Folder.
func (n *NFolder) GetList() ([]File, error) {
	panic("unimplemented")
}

// GetName implements trait.Folder.
func (n *NFolder) GetName() string {
	panic("unimplemented")
}

// GetPath implements trait.Folder.
func (n *NFolder) GetPath() string {
	panic("unimplemented")
}

// GetSize implements trait.Folder.
func (n *NFolder) GetSize() int64 {
	panic("unimplemented")
}
