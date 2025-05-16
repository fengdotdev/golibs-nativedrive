package interfaces

import "github.com/fengdotdev/golibs-traits/trait"

type DirWorker = trait.DirWorker

type CRUDWithCTX = trait.CRUDWithCTX[string, Element]
