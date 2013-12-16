package libGFC

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	g := NewGFCChain()
	_ = g.State["hi"]
}

func TestNewHost(t *testing.T) {
	priv, host := NewHost("foo")
	if priv == nil {
		t.Fatal("priv is nil")
	}

	if host.Location[0] != "foo" {
		t.Fatal("Location is not foo")
	}
}

func TestNewFile(t *testing.T) {
	host := NewFile(100)
	if host.RentedSpace != 100 {
		t.Fatal("rented space is not 100")
	}

	if len(host.Id) <= 0 {
		t.Fatal("ID not created")
	}
}

func TestNextHost(t *testing.T) {
	chain := NewGFCChain()
	if chain.NextHost().Location[0] != "127.0.0.1" {
		t.Fatal("NextHost location is %s", chain.NextHost().Location)
	}
}
