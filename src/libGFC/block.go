package libGFC

import (
	"libytc"
)

type GFCBlock struct {
	updates  []libytc.Update
	revision uint64
}

func NewGFCBlock(revision uint64, updates []libytc.Update) (g *GFCBlock) {
	return &GFCBlock{updates, revision}
}

func (g *GFCBlock) Updates() []libytc.Update {
	return g.updates
}

func (g *GFCBlock) Revision() uint64 {
	return g.revision
}

func (g *GFCBlock) Chain() string {
	return "GFC"
}

func (g *GFCBlock) ChainID() string {
	return "GFC"
}
