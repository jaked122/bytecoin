package libytcd

import (
	"libGFC"
	"testing"
	"time"
)

func TestBlockGeneration(t *testing.T) {
	t.Log("Create server")
	s := NewServer(nil)
	s.event = make(chan bool)

	e := make(chan time.Time)
	s.calculateBlock = (<-chan time.Time)(e)

	_, h := libGFC.NewHost("destination")
	record := libGFC.NewHostUpdate(h)

	c := make(chan error)

	s.TransactionChannel <- TransactionError{record, nil, c}
	<-c
	<-s.event

	if _, found := s.Keys[s.state.NextHost().Id]; !found {
		t.Log(s.state.NextHost())
		t.Log(s.state.NextHost().Id)
		t.Fatal("Next host is not us?")
	}

	e <- time.Now()
	<-s.event
	<-s.event

	if len(s.SeenTransactions) != 0 {
		t.Fatal("transaction still in queue", s.SeenTransactions)
	}

	if s.state.Revision != 1 {
		t.Fatal("Wrong revision number")
	}
}
