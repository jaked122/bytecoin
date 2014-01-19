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

func (g *GFCBlock) Apply(i interface{}) (err error) {
	s := i.(*GFCChain)
	for _, v := range g.updates {
		err = v.Verify(s)
		if err != nil {
			return
		}
		v.Apply(s)
	}
	return
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
