package libGFC

import (
	"encoding/json"
	"libytc"
	"log"
)

type BlockMessage struct {
	Type    string
	Message []byte
	Chain   string
}

type EncodedGFCBlock struct {
	Blocks   [][]byte
	Revision uint64
}

func UpdateName(u libytc.Update) (out string) {
	switch u.(type) {
	case *TransferUpdate:
		out = "TransferUpdate"
	case *HostUpdate:
		out = "HostUpdate"
	default:
		log.Print(u)
		log.Fatal("Unknown Update Type")
	}
	return
}

func MakeType(Type string, Chain string) (u libytc.Update) {
	if Chain != "GFC" {
		log.Fatal("Wrong chain %s", Chain)
	}
	switch Type {
	case "TransferUpdate":
		u = new(TransferUpdate)
	case "HostUpdate":
		u = new(HostUpdate)
	default:
		log.Fatal("Incorrect output")
	}
	return
}

type GFCEncoder struct{}

func (g GFCEncoder) EncodeUpdate(up libytc.Update) (out []byte) {
	var err error
	b := new(BlockMessage)
	b.Type = UpdateName(up)
	b.Message, err = json.Marshal(up)
	b.Chain = up.Chain()
	if err != nil {
		log.Fatal(err)
	}
	out, err = json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (g GFCEncoder) DecodeUpdate(in []byte) (up libytc.Update) {
	b := new(BlockMessage)
	err := json.Unmarshal(in, b)
	if err != nil {
		log.Fatal(err)
	}
	up = MakeType(b.Type, b.Chain)
	err = json.Unmarshal(b.Message, up)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (g GFCEncoder) EncodeBlock(block libytc.Block) (out []byte) {
	b := new(EncodedGFCBlock)
	b.Revision = block.Revision()
	for _, v := range block.Updates() {
		b.Blocks = append(b.Blocks, g.EncodeUpdate(v))
	}

	out, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (g GFCEncoder) DecodeBlock(in []byte) (block libytc.Block) {
	b := new(EncodedGFCBlock)
	err := json.Unmarshal(in, b)
	if err != nil {
		log.Fatal(err)
	}

	up := make([]libytc.Update, len(b.Blocks))

	for i, mess := range b.Blocks {
		v := g.DecodeUpdate(mess)
		up[i] = v
	}

	block = NewGFCBlock(b.Revision, up)

	return
}
