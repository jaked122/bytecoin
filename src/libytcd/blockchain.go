package libytcd

import (
	"time"
)

type FileSpace uint64
type FileHash string
type HostHash string
type YTCAmount uint64
type Location string
type Proofs string
type Address string

type HostRecord struct {
	SpaceAvailable FileSpace
	StoredFiles    []FileHash
	Balance        YTCAmount
	Location       Address
	ID             HostHash
}

type FileRecord struct {
	Balance  YTCAmount
	Proofs   []Proofs
	Hosts    []HostHash
	ID       FileHash
	Metadata struct{}
}

type Account string
type Volume uint64
type Signature string

type StoreSize uint64
type DHTLoc string
type Time time.Time

type Block struct {
	Hosts map[HostHash]HostRecord
	Files map[FileHash]FileRecord
}

func NewBlock() (b *Block) {
	b = new(Block)
	b.Hosts = make(map[HostHash]HostRecord)
	b.Files = make(map[FileHash]FileRecord)
	return
}

type BlockChain struct {
	Blocks   []*Block
	NewBlock *Block
}

func NewBlockChain() (b *BlockChain) {
	b = new(BlockChain)
	b.Blocks = make([]*Block, 100)
	b.NewBlock = NewBlock()
	return
}
