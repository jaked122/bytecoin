package libytcd

import (
	"errors"
)

type Update interface {
	Verify(s *State) (err error)
	Apply(s *State)
}

type TransferUpdate struct {
	Type        string
	Source      HostKey
	Destination HostKey
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

func (t *TransferUpdate) Apply(s *State) {
	source := s.Hosts[t.Source]
	source.Balance -= t.Amount
	s.Hosts[t.Source] = source
	dest := s.Hosts[t.Destination]
	dest.Balance += t.Amount
	s.Hosts[t.Destination] = dest
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

func (t *HostUpdate) Verify(s *State) (err error) {
	// Verify signature from old key
	o, found := s.Hosts[t.Key]
	if found && o.Key != t.Key {
		return errors.New("Host collision")
	}

	//verify Signature
	return
}

func (t *HostUpdate) Apply(s *State) {
	host := s.Hosts[t.Key]
	host.Key = t.Key
	s.Hosts[t.Key] = host
	return
}
