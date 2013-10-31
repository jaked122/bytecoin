package libytcd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"log"
)

type Update interface {
	Verify(s *State) (err error)
	Apply(s *State)
}

type TransferUpdate struct {
	Type        string
	Source      HostHash
	Destination HostHash
	Amount      YTCAmount
	Signature   Signature
	//Public Key? Check bitcoin
}

func (t *TransferUpdate) Verify(s *State) (err error) {

	_, found := s.Hosts[t.Source]
	if !found {
		return errors.New("No Such Source")
	}

	_, found = s.Hosts[t.Destination]
	if !found {
		return errors.New("No such destination")
	}

	if t.Amount > s.Hosts[t.Source].Balance {
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
	t.Signature = Signature{s, r}
	return
}

func (t *TransferUpdate) Apply(s *State) {
	source := s.Hosts[t.Source]
	source.Balance -= t.Amount
	s.Hosts[t.Source] = source
	dest := s.Hosts[t.Destination]
	dest.Balance += t.Amount
	s.Hosts[t.Destination] = dest
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

func NewTransferUpdate() (t *TransferUpdate) {
	t = new(TransferUpdate)
	t.Type = "TransferUpdate"
	return
}

type HostUpdate struct {
	Type      string
	Key       HostKey
	Signature Signature
}

func NewHostUpdate() (h *HostUpdate) {
	h = new(HostUpdate)
	h.Type = "HostUpdate"
	return
}

func (t *HostUpdate) Sign(key *ecdsa.PrivateKey) {
	h := []byte(t.String())
	s, r, err := ecdsa.Sign(rand.Reader, key, h)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = Signature{s, r}
	return
}

func (t *HostUpdate) Verify(s *State) (err error) {
	// Verify signature from old key
	o, found := s.Hosts[t.Key.Hash()]
	if found && o.Key != t.Key {
		return errors.New("Host collision")
	}

	//verify Signature
	return
}

func (t *HostUpdate) String() (str string) {
	str = "Hostupdate\n"
	str += "Key:" + t.Key.String() + "\n"
	return
}

func (t *HostUpdate) Apply(s *State) {
	host := s.Hosts[t.Key.Hash()]
	host.Key = t.Key
	s.Hosts[t.Key.Hash()] = host
	return
}
