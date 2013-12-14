package libGFC

import (
	"encoding/json"
	"log"
)

type Encoder interface {
	EncodeUpdate(Update) []byte
	EncodeUpdates([]Update) []byte
	DecodeUpdate([]byte) Update
	DecodeUpdates([]byte) []Update
}

type BlockMessage struct {
	Type    string
	Message []byte
	Chain   string
}

type Block struct {
	Blocks [][]byte
}

func UpdateName(u Update) (out string) {
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

func MakeType(Type string, Chain string) (u Update) {
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

func (g GFCEncoder) EncodeUpdate(up Update) (out []byte) {
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

func (g GFCEncoder) DecodeUpdate(in []byte) (up Update) {
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

func (g GFCEncoder) EncodeUpdates(up []Update) (out []byte) {
	b := new(Block)
	for _, v := range up {
		b.Blocks = append(b.Blocks, g.EncodeUpdate(v))
	}

	out, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (g GFCEncoder) DecodeUpdates(in []byte) (up []Update) {
	b := new(Block)
	err := json.Unmarshal(in, b)
	if err != nil {
		log.Fatal(err)
	}

	for _, mess := range b.Blocks {
		v := g.DecodeUpdate(mess)
		up = append(up, v)
	}
	return
}
