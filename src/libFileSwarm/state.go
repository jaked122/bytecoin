package libFileSwarm

import (
	"libGFC"
)

type State struct {
	swarmtracker libGFC.GFCChain

	swarmid        string
	hostcount      uint64
	hostredundancy uint64
	totalspace     uint64

	piecemapping map[string][]string

	previousblocks []*Block
	currentblock   *Block
}

type Signal struct {
}

type Indictment struct {
}
