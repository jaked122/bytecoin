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

	vs := v.(*TransferUpdate)
	if *vs != *r {
		t.Log(vs.Source == r.Source)
		t.Log(vs.Destination == r.Destination)
		t.Log(vs.Amount == r.Amount)
		t.Log(vs.Signature == r.Signature)
		t.Log(vs.Signature)
		t.Log(r.Signature)
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
