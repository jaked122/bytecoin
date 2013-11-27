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

	_ = <-a.s.event

	t.Log("Recieved event")

	if len(a.s.ports) != 1 {
		t.Log(a.s.ports)
		t.Fatal("Failure to connect")
	}

	priv, _ := libytc.DeterministicKey(0)
	_, h := libGFC.NewHost("destination")
	record := libGFC.NewHostUpdate(h)

	a.s.transaction <- TransactionError{record, nil, c}
	<-c

	u := libGFC.NewTransferUpdate("Origin", "destination", 1)
	u.Sign(priv)
	a.s.transaction <- TransactionError{u, nil, c}

	<-c

	if len(b.s.SeenTransactions) != 1 {
		t.Fatal("Length is not one, length is %d", len(b.s.SeenTransactions))
	}

}
