package libytcd

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	r := NewTransferUpdate()
	priv, pub := OriginKey()
	r.Source = pub.Hash()
	r.Destination = pub.Hash()
	r.Amount = 1
	r.Sign(priv)

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
	h.Key = pub
	h.Sign(priv)

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
