package libytcd

import (
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	y1 := NewYtcd()
	c := make(chan struct{})
	go func() {
		c <- struct{}{}
		err := y1.ListenNetwork(":1337")
		t.Fatal(err)
	}()
	_ = <-c

	y2 := NewYtcd()
	go func() {
		c <- struct{}{}
		err := y2.ListenNetwork(":1338")
		t.Fatal(err)
	}()
	_ = <-c

	err := y2.ConnectAddress("127.0.0.1:1337")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	if len(y1.neighbors) != 1 {
		t.Fatal("y1 didn't connect")
	}

	time.Sleep(1)

	if len(y2.neighbors) != 1 {
		t.Log(y1)
		t.Log(y2)
		t.Fatal("y2 didn't connect")
	}

	y2.Send(&HostUpdate{"HostUpdate", "foo", "foo"})
	time.Sleep(1 * time.Second)
	y2.Send(&TransferUpdate{"TransferUpdate", "hard", "foo", 1, "sig"})

	time.Sleep(1 * time.Second)

	if y1.s.Hosts["foo"].Balance != 1 {
		t.Log(y1.s.Hosts)
		t.Fatal("Balance is not 1")
	}
}
