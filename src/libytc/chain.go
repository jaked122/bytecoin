package libytc

type Update interface {
	Verify(interface{}) (err error)
	Apply(interface{})
	Chain() string
	ChainID() string
	String() string
}

type Block interface {
	Revision() uint64
	Updates() []Update
	Chain() string
	ChainID() string
}

type Encoder interface {
	EncodeUpdate(Update) []byte
	DecodeUpdate([]byte) Update
	EncodeBlock(Block) []byte
	DecodeBlock([]byte) Block
}
