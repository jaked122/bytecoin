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

func MakeType(Type string) (u Update) {
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

func EncodeUpdate(up Update) (out []byte) {
	var err error
	b := new(BlockMessage)
	b.Type = UpdateName(up)
	b.Message, err = json.Marshal(up)
	if err != nil {
		log.Fatal(err)
	}
	out, err = json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DecodeUpdate(in []byte) (up Update) {
	b := new(BlockMessage)
	err := json.Unmarshal(in, b)
	if err != nil {
		log.Fatal(err)
	}
	up = MakeType(b.Type)
	err = json.Unmarshal(b.Message, up)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func EncodeUpdates(up []Update) (out []byte) {
	b := new(Block)
	for _, v := range up {
		b.Blocks = append(b.Blocks, EncodeUpdate(v))
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
		v := DecodeUpdate(mess)
		up = append(up, v)
	}
	return
}
