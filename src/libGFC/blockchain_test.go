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

	if host.Location != "foo" {
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
