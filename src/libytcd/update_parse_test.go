package libytcd

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	r := NewTransferUpdate()
	r.Source = "Source"
	r.Destination = "Destination"
	r.Amount = 1
	r.Signature = "Signature"

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	v, err := ParseUpdate(b)
	if err != nil {
		t.Fatal(err)
	}

	if *v.(*TransferUpdate) != *r {
		t.Fatal(v, "!=", r)
	}

	h := NewHostUpdate()
	h.Key = "Key"
	h.Signature = "Signature"

	b, err = json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}

	v, err = ParseUpdate(b)
	if err != nil {
		t.Fatal(err)
	}

	if *v.(*HostUpdate) != *h {
		t.Fatal(v, "!=", r)
	}
}
