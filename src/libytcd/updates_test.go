package libytcd

import (
	"testing"
)

func SimpleTest(t *testing.T) {
	s := NewState()

	priv, pub := OriginKey()
	npriv, npub := DeterministicKey(1)

	h := NewHostUpdate()
	h.Key = npub
	h.Sign(npriv)

	if h.Verify(s) != nil {
		t.Fatal("Failed to verify HostUpdate")
	}

	h.Apply(s)

	r := NewTransferUpdate()
	r.Source = pub.Hash()
	r.Destination = npub.Hash()
	r.Amount = 1
	r.Sign(priv)

	if r.Verify(s) != nil {
		t.Fatal("Failed to verify transaction")
	}

	r.Apply(s)

	if s.Hosts[npub.Hash()].Balance != 1 {
		t.Fatal("New coin balance is to low")
	}

}
