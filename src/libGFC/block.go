package libGFC

import (
	"encoding/json"
	"log"
)

type BlockMessage struct {
	Type    string
	Message []byte
}

type Block struct {
	Blocks []BlockMessage
}

func (b *Block) UpdateName(u Update) (out string) {
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

func (b *Block) MakeType(Type string) (u Update) {
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

func EncodeUpdates(up []Update) (out []byte) {
	b := new(Block)
	for _, v := range up {
		name := b.UpdateName(v)
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		message := BlockMessage{name, bytes}
		b.Blocks = append(b.Blocks, message)
	}

	out, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DecodeUpdates(in []byte) (up []Update) {
	b := new(Block)
	err := json.Unmarshal(in, b)
	if err != nil {
		log.Fatal(err)
	}

	for _, mess := range b.Blocks {
		v := b.MakeType(mess.Type)
		err := json.Unmarshal(mess.Message, v)
		if err != nil {
			log.Fatal(err)
		}

		up = append(up, v)
	}
	return
}
