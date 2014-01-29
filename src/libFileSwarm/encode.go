package libFileSwarm

import (
	"encoding/json"
	"libytc"
	"log"
)

type Encoder struct {
}

func (e Encoder) EncodeBlock(b libytc.Block) []byte {
	block := b.(*Block)

	encodedBlock, err := json.Marshal(block)
	if err != nil {
		log.Fatal(err)
	}

	return encodedBlock
}

func (e Encoder) DecodeBlock(b []byte) libytc.Block {
	block := new(Block)
	json.Unmarshal(b, block)
	return block
}
