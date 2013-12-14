package libGFC

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"libytc"
	"log"
)

type TransferUpdate struct {
	Source      string
	Destination string
	Amount      uint64
	Signature   libytc.Signature
	//Public Key? Check bitcoin
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

	//Verify Signature
	if !ecdsa.Verify(&h.KeyList[0].PublicKey, []byte(t.String()), t.Signature.R, t.Signature.S) {
		return errors.New("Invalid Signature")
	}

	return
}

func (t *TransferUpdate) Sign(key *ecdsa.PrivateKey) {
	h := []byte(t.String())
	s, r, err := ecdsa.Sign(rand.Reader, key, h)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = libytc.Signature{s, r}
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

type HostUpdate struct {
	Record    *FileChainRecord
	Signature libytc.Signature
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
	t.Signature = libytc.Signature{s, r}
	return
}

func (t *HostUpdate) Verify(i interface{}) (err error) {
	s := i.(*GFCChain)
	// Verify signature from old key
	h, found := s.State[t.Record.Id]
	if !found {
		return
	}

	if !ecdsa.Verify(&h.KeyList[0].PublicKey, []byte(t.String()), t.Signature.R, t.Signature.S) {
		return errors.New("Invalid Signature")
	}

	return
}

func (t *HostUpdate) Apply(i interface{}) {
	s := i.(*GFCChain)
	s.State[t.Record.Id] = t.Record
	return
}

func (t *HostUpdate) Chain() string {
	return "GFC"
}

func (t *HostUpdate) String() (str string) {
	str = "Hostupdate\n"
	str += "Record:" + fmt.Sprint(t.Record) + "\n"
	return
}

type UnknownUpdate struct {
	payload string
	chain   string
	Type    string
}

func NewUnknownUpdate(payload string, chain string, Type string) (t *UnknownUpdate) {
	t = new(UnknownUpdate)
	t.payload = payload
	t.chain = chain
	t.Type = Type
	return
}

func (t *UnknownUpdate) Verify(i interface{}) (err error) {
	log.Fatal("Cannot Verify Unknown Update")
	return
}

func (t *UnknownUpdate) Apply(i interface{}) {
	log.Fatal("Cannot Apply Unknown Update")
	return
}

func (t *UnknownUpdate) Chain() string {
	return t.chain
}

func (t *UnknownUpdate) String() string {
	return fmt.Sprint(t)
}

func (t *UnknownUpdate) MarshallJSON(b []byte, err error) {
	b = []byte(t.payload)
	return
}
