package libFileSwarm

import (
	"libytc"
)

type Block struct {
	BlockNumber uint64
	BlockHash   string
	ChainId     string

	entropyhash   map[string]string
	entropystring map[string]string

	storagehash   map[string]string
	storagestring map[string]string

	incomingsignals []*Signal
	outgoinsignals  []*Signal

	hostsignatures map[string]*libytc.SignatureMap
	indictments    []*Indictment

	Transactionproofs []libytc.Update
}

func (b *Block) Revision() uint64 {
	return b.BlockNumber
}

func (b *Block) Chain() string {
	return "FileSwarm"
}

func (b *Block) ChainID() string {
	return b.ChainId
}
func (b *Block) Apply(i interface{}) (err error) {
	s := i.(*State)
	s = s
	return
}
