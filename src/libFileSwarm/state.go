package libFileSwarm

import (
	"libGFC"
	"libytc"
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

type Block struct {
	BlockNumber uint64
	BlockHash   string

	entropyhash   map[string]string
	entropystring map[string]string

	storagehash   map[string]string
	storagestring map[string]string

	incomingsignals []*Signal
	outgoinsignals  []*Signal

	hostsignatures map[string]*libytc.SignatureMap
	indictments    []*Indictment

	Transactionproofs []libytc.Updates
}

type Signal struct {
}

type Indictment struct {
}
