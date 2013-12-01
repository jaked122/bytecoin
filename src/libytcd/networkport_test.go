package libytcd

import (
	"libGFC"
	"libytc"
	"testing"
	"time"
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
	_, err := b.ConnectAddress("127.0.0.1:1777")
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

func TestRemoteUpdates(t *testing.T) {
	//Create a server with a block

	s := NewServer(nil)
	s.event = make(chan bool)

	e := make(chan time.Time)
	s.calculateBlock = (<-chan time.Time)(e)

	e <- time.Now()
	<-s.event
	<-s.event

	if s.state.Revision != 1 {
		t.Fatal("Wrong revision number")
	}

	p1 := NewNetworkPort(s)

	c := make(chan error)
	go func() {
		c <- nil
		err := p1.ListenNetwork("127.0.0.1:1338")
		if err != nil {
			t.Fatal(err)
		}
	}()
	<-c

	s2 := NewServer(nil)
	s2.event = make(chan bool)

	p2 := NewNetworkPort(s2)
	pc := make(chan *NetworkConnection)
	go func() {
		port, err := p2.ConnectAddress("127.0.0.1:1338")
		if err != nil {
			t.Fatal(err)
		}
		pc <- port
	}()

	<-s2.event
	<-s.event
	p2c := <-pc
	p2c.Reconnect()

	if s2.state.Revision != 1 {
		t.Fatal(s2.state.Revision)
	}
}
