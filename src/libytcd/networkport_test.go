package libytcd

import (
	"log"
	"testing"
)

func TestNetworkPortSimple(t *testing.T) {
	t.Log("Make stuff")
	a := NewNetworkPort(NewServer(nil))
	b := NewNetworkPort(NewServer(nil))
	t.Log("Done making")

	a.s.event = make(chan bool)

	c := make(chan bool)
	go func() {
		c <- true
		err := a.ListenNetwork("127.0.0.1:1777")
		if err != nil {
			log.Fatal(err)
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
}
