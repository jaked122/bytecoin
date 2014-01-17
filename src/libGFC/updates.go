package libGFC

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"libytc"
)

type TransferUpdate struct {
	Source      string
	Destination string
	Amount      uint64
	Signature   *libytc.SignatureMap
}

func (t *TransferUpdate) Verify(i interface{}) (err error) {

	s := i.(*GFCChain)

	h, found := s.State[t.Source]
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

	return h.Verify(t.String(), t.Signature)
}

func (t *TransferUpdate) Sign(key *ecdsa.PrivateKey) {
	t.Signature = libytc.SignMap(nil, t.String(), key)
	return
}

func (t *TransferUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	source := s.State[t.Source]
	source.Balance -= t.Amount
	s.State[t.Source] = source
	dest := s.State[t.Destination]
	dest.Balance += t.Amount
	s.State[t.Destination] = dest
	return
}

func (t *TransferUpdate) Chain() string {
	return "GFC"
}

func (t *TransferUpdate) ChainID() string {
	return "GFC"
}

func (t *TransferUpdate) String() (str string) {
	str = "TransferUpdate\n"
	str += "Source:" + string(t.Source) + "\n"
	str += "Destination:" + string(t.Destination) + "\n"
	str += "Amount:" + string(t.Amount) + "\n"
	return
}

func NewTransferUpdate(source string, destination string, amount uint64) (t *TransferUpdate) {
	t = new(TransferUpdate)
	t.Source = source
	t.Destination = destination
	t.Amount = amount
	return
}

type LocationUpdate struct {
	Id        string
	Location  []string
	Signature *libytc.SignatureMap
}

func NewLocationUpdate(Id string, Location []string) (l *LocationUpdate) {
	l = new(LocationUpdate)
	l.Id = Id
	l.Location = Location
	return
}

func (l *LocationUpdate) Verify(i interface{}) (err error) {
	s := i.(*GFCChain)

	h, found := s.State[l.Id]
	if !found {
		return errors.New("Id does not exist")
	}

	return h.Verify(l.String(), l.Signature)
}

func (l *LocationUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	s.State[l.Id].Location = l.Location
	return
}

func (l *LocationUpdate) String() (s string) {
	s = "LocationUpdate\n"
	s += fmt.Sprint("Id: %s\n", l.Id)
	s += fmt.Sprint("Location: %s\n", fmt.Sprint(l.Location))
	return
}

func (l *LocationUpdate) Sign(key *ecdsa.PrivateKey) {
	l.Signature = libytc.SignMap(nil, l.String(), key)
	return
}

func (l *LocationUpdate) Chain() string {
	return "GFC"
}

func (l *LocationUpdate) ChainID() string {
	return "GFC"
}

type HostUpdate struct {
	Record    *FileChainRecord
	Signature *libytc.SignatureMap
}

func NewHostUpdate(f *FileChainRecord) (h *HostUpdate) {
	h = new(HostUpdate)
	h.Record = f
	return
}

func (t *HostUpdate) Sign(key *ecdsa.PrivateKey) {
	t.Signature = libytc.SignMap(nil, t.String(), key)
	return
}

func (t *HostUpdate) Verify(i interface{}) (err error) {
	s := i.(*GFCChain)
	// Verify signature from old key
	h, found := s.State[t.Record.Id]
	if !found {
		return
	}

	return h.Verify(t.String(), t.Signature)
}

func (t *HostUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	s.State[t.Record.Id] = t.Record
	return
}

func (t *HostUpdate) Chain() string {
	return "GFC"
}

func (t *HostUpdate) ChainID() string {
	return "GFC"
}

func (t *HostUpdate) String() (str string) {
	str = "Hostupdate\n"
	str += "Record:" + fmt.Sprint(t.Record) + "\n"
	return
}

type KeyUpdate struct {
	Id        string
	KeyList   map[string]float64
	Signature *libytc.SignatureMap
}

func NewKeyUpdate(Id string, KeyList map[string]float64) (k *KeyUpdate) {
	k = new(KeyUpdate)
	k.Id = Id
	k.KeyList = KeyList
	return
}

func (k *KeyUpdate) Verify(i interface{}) (err error) {
	s := i.(*GFCChain)

	h, found := s.State[k.Id]
	if !found {
		return errors.New("Id does not exist")
	}

	return h.Verify(k.String(), k.Signature)
}

func (k *KeyUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	s.State[k.Id].KeyList = k.KeyList
	return
}

func (k *KeyUpdate) String() (s string) {
	s = "KeyUpdate\n"
	s += fmt.Sprint("Id: %s\n", k.Id)
	s += fmt.Sprint("KeyList: %s\n", fmt.Sprint(k.KeyList))
	return
}

func (k *KeyUpdate) Sign(key *ecdsa.PrivateKey) {
	k.Signature = libytc.SignMap(k.Signature, k.String(), key)
	return
}

func (k *KeyUpdate) Chain() string {
	return "GFC"
}

func (k *KeyUpdate) ChainID() string {
	return "GFC"
}

type SpaceUpdate struct {
	Id             string
	AvailableSpace uint64
	Signature      *libytc.SignatureMap
}

func NewSpaceUpdate(Id string, AvailableSpace uint64) (k *SpaceUpdate) {
	k = new(SpaceUpdate)
	k.Id = Id
	k.AvailableSpace = AvailableSpace
	return
}

func (k *SpaceUpdate) Verify(i interface{}) (err error) {
	s := i.(*GFCChain)

	h, found := s.State[k.Id]
	if !found {
		return errors.New("Id does not exist")
	}

	return h.Verify(k.String(), k.Signature)
}

func (k *SpaceUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	s.State[k.Id].AvailableSpace = k.AvailableSpace
	return
}

func (k *SpaceUpdate) String() (s string) {
	s = "SpaceUpdate\n"
	s += fmt.Sprint("Id: %s\n", k.Id)
	s += fmt.Sprint("AvailableSpace: %s\n", fmt.Sprint(k.AvailableSpace))
	return
}

func (k *SpaceUpdate) Sign(key *ecdsa.PrivateKey) {
	k.Signature = libytc.SignMap(k.Signature, k.String(), key)
	return
}

func (k *SpaceUpdate) Chain() string {
	return "GFC"
}

func (k *SpaceUpdate) ChainID() string {
	return "GFC"
}
