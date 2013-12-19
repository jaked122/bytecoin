package libGFC

import (
	"crypto/ecdsa"
	"libytc"
	"testing"
)

func barChain() (priv *ecdsa.PrivateKey, record *FileChainRecord, s *GFCChain) {
	s = NewGFCChain()
	priv, record = NewHost("bar")
	update := NewHostUpdate(record)
	update.Apply(s)
	return
}

func LocationUpdateTest(t *testing.T) {
	priv, record, s := barChain()

	l := NewLocationUpdate(record.Id, []string{"foo"})
	l.Sign(priv)

	if err := l.Verify(s); err != nil {
		t.Fatal(err)
	}

	l.Apply(s)

	if s.State[record.Id].Location[0] != "foo" {
		t.Fatal("Location != foo")
	}
}

func SimpleTest(t *testing.T) {
	priv, record, s := barChain()

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

func TestTranferEncoding(t *testing.T) {

	priv, record, s := barChain()

	h := NewTransferUpdate(record.Id, record.Id, 0)
	h.Sign(priv)

	b := GFCEncoder{}.EncodeUpdate(h)
	j := GFCEncoder{}.DecodeUpdate(b)
	if len(j.(*TransferUpdate).Signature.M) != 1 {
		t.Fatal(j.(*TransferUpdate).Signature.M)
	}

	if h.Verify(s) != nil {
		t.Fatal(h.Verify(s))
	}

	t.Log(string(b))

	if j.Verify(s) != nil {
		t.Fatal(j.Verify(s))
	}
}

func TestKeyUpdate(t *testing.T) {

	priv, record, s := barChain()

	_, npub := libytc.DeterministicKey(2)

	NewKeyList := make(map[string]float64)
	NewKeyList[npub.Hash()] = 1.0

	k := NewKeyUpdate(record.Id, NewKeyList)
	k.Sign(priv)

	b := GFCEncoder{}.EncodeUpdate(k)
	j := GFCEncoder{}.DecodeUpdate(b)

	err := k.Verify(s)
	if err != nil {
		t.Fatal(err)
	}

	err = j.Verify(s)
	if err != nil {
		t.Fatal(err)
	}

	j.Apply(s)

	if s.State[record.Id].KeyList[npub.Hash()] != 1.0 {
		t.Fatal(s.State[record.Id])
	}
}
