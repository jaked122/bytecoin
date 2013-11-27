package libytcd

import (
	"libGFC"
	"libytc"
	"testing"
)

func TestNetworkPortSimple(t *testing.T) {
	t.Log("Make stuff")
	a := NewNetworkPort(NewServer(nil))
	b := NewNetworkPort(NewServer(nil))
	t.Log("Done making")

	a.s.event = make(chan bool)

	c := make(chan error)
	go func() {
		c <- nil
		err := a.ListenNetwork("127.0.0.1:1777")
		if err != nil {
			t.Fatal(err)
		}
	}()

	t.Log("Before Block")
	<-c

	t.Log("Connecting")
	err := b.ConnectAddress("127.0.0.1:1777")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Connected")

	<-a.s.event

	t.Log("Recieved event")

	if len(a.s.ports) != 1 {
		t.Log(a.s.ports)
		t.Fatal("Failure to connect")
	}

	b.s.event = a.s.event
	a.s.event = nil

	priv, _ := libytc.DeterministicKey(0)
	_, h := libGFC.NewHost("destination")
	record := libGFC.NewHostUpdate(h)

	a.s.TransactionChannel <- TransactionError{record, nil, c}
	err = <-c
	if err != nil {
		t.Fatal(err)
	}
	<-b.s.event

	u := libGFC.NewTransferUpdate("Origin", "Origin", 1)
	u.Sign(priv)

	a.s.TransactionChannel <- TransactionError{u, nil, c}
	err = <-c
	if err != nil {
		t.Fatal(err)
	}
	<-b.s.event

	if len(b.s.SeenTransactions) != 2 {
		t.Fatal("Length is not one, length is %d", len(b.s.SeenTransactions))
	}

}
