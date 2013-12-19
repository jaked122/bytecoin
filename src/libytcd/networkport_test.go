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

	a.s.event = make(chan string)

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

	c = make(chan error)
	a.s.TransactionChannel <- TransactionError{record, nil, c}
	err = <-c
	if err != nil {
		t.Fatal(err)
	}
	o := <-b.s.event
	t.Log(o)

	u := libGFC.NewTransferUpdate("Origin", "Origin", 1)
	u.Sign(priv)

	if len(u.Signature.M) != 1 {
		t.Fatal("SignatureMap is not created")
	}

	c = make(chan error)
	a.s.TransactionChannel <- TransactionError{u, nil, c}
	err = <-c
	if err != nil {
		t.Fatal(err)
	}
	o = <-b.s.event
	t.Log(o)

	if len(b.s.SeenTransactions) != 2 {
		t.Fatal("Length is not 2, length is", len(b.s.SeenTransactions), b.s.SeenTransactions)
	}

}

func TestRemoteUpdates(t *testing.T) {
	//Create a server with a block

	s := NewServer(nil)
	s.event = make(chan string)

	e := make(chan time.Time)
	s.calculateBlock = (<-chan time.Time)(e)

	e <- time.Now()
	<-s.event
	<-s.event

	if s.state["GFC"].Revision != 1 {
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
	s2.event = make(chan string)

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
	p2c.Reconnect("GFC")

	if s2.state["GFC"].Revision != 1 {
		t.Fatal(s2.state["GFC"].Revision)
	}
}
