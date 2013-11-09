package libGFC

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"libytcd"
	"log"
)

type Update interface {
	Verify(s *GFCChain) (err error)
	Apply(s *GFCChain)
}

type TransferUpdate struct {
	Source      string
	Destination string
	Amount      uint64
	Signature   libytcd.Signature
	//Public Key? Check bitcoin
}

func (t *TransferUpdate) Verify(s *GFCChain) (err error) {

	_, found := s.State[t.Source]
	if !found {
		return errors.New("No Such Source")
	}

	_, found = s.State[t.Destination]
	if !found {
		return errors.New("No such destination")
	}

	if t.Amount > s.State[t.Source].Balance {
		return errors.New("Not enough money in source account")
	}

	//Verify Signature
	return
}

func (t *TransferUpdate) Sign(key *ecdsa.PrivateKey) {
	h := []byte(t.String())
	s, r, err := ecdsa.Sign(rand.Reader, key, h)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = libytcd.Signature{s, r}
	return
}

func (t *TransferUpdate) Apply(s *GFCChain) {
	source := s.State[t.Source]
	source.Balance -= t.Amount
	s.State[t.Source] = source
	dest := s.State[t.Destination]
	dest.Balance += t.Amount
	s.State[t.Destination] = dest
	return
}

func (t *TransferUpdate) String() (str string) {
	str = "TransferUpdate\n"
	str += "Source:" + string(t.Source) + "\n"
	str += "Destination:" + string(t.Destination) + "\n"
	str += "Amount:" + string(t.Amount) + "\n"
	log.Print(str)
	return
}

func NewTransferUpdate(source string, destination string, amount uint64) (t *TransferUpdate) {
	t = new(TransferUpdate)
	t.Source = source
	t.Destination = destination
	t.Amount = amount
	return
}

type HostUpdate struct {
	Record    *FileChainRecord
	Signature libytcd.Signature
}

func NewHostUpdate(f *FileChainRecord) (h *HostUpdate) {
	h = new(HostUpdate)
	h.Record = f
	return
}

func (t *HostUpdate) Sign(key *ecdsa.PrivateKey) {
	h := []byte(t.String())
	s, r, err := ecdsa.Sign(rand.Reader, key, h)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = libytcd.Signature{s, r}
	return
}

func (t *HostUpdate) Verify(s *GFCChain) (err error) {
	// Verify signature from old key
	_, found := s.State[t.Record.Id]
	if found {
		//verify Signature
	}
	return
}

func (t *HostUpdate) String() (str string) {
	str = "Hostupdate\n"
	str += "Record:" + fmt.Sprint(t.Record) + "\n"
	return
}

func (t *HostUpdate) Apply(s *GFCChain) {
	s.State[t.Record.Id] = t.Record
	return
}
