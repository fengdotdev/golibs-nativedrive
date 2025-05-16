package ndrive

import (
	"context"

	"github.com/fengdotdev/golibs-nativedrive/interfaces"
)

var _ interfaces.CRUDWithCTX = (*NDrive)(nil)

// All implements trait.CRUDWithCTX.
func (n *NDrive) All(ctx context.Context) map[string]interfaces.Element {
	panic("unimplemented")
}

// Clean implements trait.CRUDWithCTX.
func (n *NDrive) Clean(ctx context.Context) {
	panic("unimplemented")
}

// Create implements trait.CRUDWithCTX.
func (n *NDrive) Create(ctx context.Context, id string, item interfaces.Element) error {
	panic("unimplemented")
}

// Delete implements trait.CRUDWithCTX.
func (n *NDrive) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Exists implements trait.CRUDWithCTX.
func (n *NDrive) Exists(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// Iterate implements trait.CRUDWithCTX.
func (n *NDrive) Iterate(ctx context.Context, fn func(string, interfaces.Element) (stop bool, err error)) error {
	panic("unimplemented")
}

// Keys implements trait.CRUDWithCTX.
func (n *NDrive) Keys(ctx context.Context) []string {
	panic("unimplemented")
}

// Len implements trait.CRUDWithCTX.
func (n *NDrive) Len(ctx context.Context) int {
	panic("unimplemented")
}

// Populate implements trait.CRUDWithCTX.
func (n *NDrive) Populate(ctx context.Context, items map[string]interfaces.Element) {
	panic("unimplemented")
}

// Read implements trait.CRUDWithCTX.
func (n *NDrive) Read(ctx context.Context, id string) (interfaces.Element, error) {
	panic("unimplemented")
}

// Update implements trait.CRUDWithCTX.
func (n *NDrive) Update(ctx context.Context, id string, item interfaces.Element) error {
	panic("unimplemented")
}

// Values implements trait.CRUDWithCTX.
func (n *NDrive) Values(ctx context.Context) []interfaces.Element {
	panic("unimplemented")
}
