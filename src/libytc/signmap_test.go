package libytc

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestSigMap(t *testing.T) {

	var v *SignatureMap
	v = NewSignatureMap(1)
	_, key := DeterministicKey(0)
	v.M[key] = Signature{big.NewInt(1), big.NewInt(2)}

	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))

	q := NewSignatureMap(0)
	err = json.Unmarshal(b, &q)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(q == nil)

	if len(v.M) != len(q.M) {
		t.Fatal(*v, *q)
	}
}
