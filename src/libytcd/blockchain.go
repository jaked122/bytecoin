package libytcd

import (
	"errors"
	"time"
)

type Account string
type Volume uint64
type Signature string

type StoreSize uint64
type DHTLoc string
type Time time.Time

type Transaction struct {
	Source      Account
	Destination Account
	Amount      Volume
	Expiration  Time
	Signature   Signature
}

type StorageAnnounce struct {
	Source      Account
	Destination Account
	Size        StoreSize
	Location    DHTLoc
	Signature   Signature
}

type BlockChain struct {
	state        map[Account]Volume
	transactions []Transaction
}

func NewBlockChain() (b *BlockChain) {
	b = new(BlockChain)
	b.transactions = make([]Transaction, 100)
	b.state = make(map[Account]Volume)
	return
}

func (b *BlockChain) VerifyTransaction(t Transaction) (err error) {

	// need to verify that both src and dest exist

	if b.state[t.Source] < t.Amount {
		return errors.New("Insufficient Funds")
	}

	// need to verify that transaction hasn't expired

	// need to verify that the signature is valid

	return nil
}

func (b *BlockChain) AddTransaction(t Transaction) (err error) {
	err = b.VerifyTransaction(t)
	if err != nil {
		return err
	}

	b.transactions = append(b.transactions, t)
	b.state[t.Source] -= t.Amount
	b.state[t.Destination] += t.Amount

	return nil
}

func (b *BlockChain) VerifyAnnouceStorage(s StorageAnnounce) (err error) {

	// check that source exists
	// check that source can pay deductable

	// check that destination exists

	// check that DHT loc is online (make a request through the DHT)

	// check that the signature is valid

	return nil
}

func (b *BlockChain) AnnounceStorage(s StorageAnnounce) (err error) {
	err = b.VerifyAnnouceStorage(s)
	if err != nil {
		return err
	}

	return nil
}
