package libytcd

import (
	"testing"
)

func SimpleTest(t *testing.T) {
	s := NewState()

	h := NewHostUpdate()
	h.Key = "New"
	h.Signature = "Signature"

	if !h.Verify(s) {
		t.Fatal("Failed to verify HostUpdate")
	}

	h.Apply(s)

	r := NewTransferUpdate()
	r.Source = "hard"
	r.Destination = "New"
	r.Amount = 1

	if !r.Verify(s) {
		t.Fatal("Failed to verify transaction")
	}

	r.Apply(s)

	if s.Hosts["New"].Balance <= 0 {
		t.Fatal("New coin balance is to low")
	}

}
