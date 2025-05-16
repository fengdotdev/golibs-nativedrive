package ndrive

import (
	"sync"

	"github.com/fengdotdev/golibs-nativedrive/interfaces"
)

type NDrive struct {
	workingDir string
	index      map[string]*interfaces.Element
	self       interfaces.Element
	mu         sync.RWMutex
}
