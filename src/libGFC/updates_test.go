package libGFC

import (
	"testing"
)

func SimpleTest(t *testing.T) {
	s := NewGFCChain()

	priv, record := NewHost("bar")
	update := NewHostUpdate(record)
	update.Apply(s)

	c := s.State[record.Id]
	c.Balance = 1
	s.State[record.Id] = c

	npriv, nrecord := NewHost("foo")

	h := NewHostUpdate(record)
	h.Sign(npriv)

	if h.Verify(s) != nil {
		t.Fatal("Failed to verify HostUpdate")
	}

	h.Apply(s)

	r := NewTransferUpdate(record.Id, nrecord.Id, 1)
	r.Sign(priv)

	if r.Verify(s) != nil {
		t.Fatal("Failed to verify transaction")
	}

	r.Apply(s)

	if s.State[nrecord.Id].Balance != 1 {
		t.Fatal("New coin balance is to low")
	}

}
