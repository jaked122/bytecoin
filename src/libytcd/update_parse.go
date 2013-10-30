package libytcd

import (
	"encoding/json"
	"errors"
	"log"
)

func ParseUpdate(raw []byte) (u Update, err error) {
	type TypePeek struct {
		Type string
	}

	v := &TypePeek{}

	err = json.Unmarshal(raw, v)
	if err != nil {
		return
	}

	switch v.Type {
	case "TransferUpdate":
		t := NewTransferUpdate()
		err = json.Unmarshal(raw, t)
		u = t
	case "HostUpdate":
		h := NewHostUpdate()
		err = json.Unmarshal(raw, h)
		u = h
	default:
		log.Print(string(raw))
		err = errors.New("Invalid Type: " + v.Type)
	}
	return
}
