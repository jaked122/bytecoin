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

type FileID string

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

type StorageRent struct {
	Source     Account
	Size       StoreSize
	Servers    uint64
	Redundancy float64
	ID         FileID
	Expiration Time
	Signature  Signature
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

func (b *BlockChain) VerifyAnnouceStorage(a StorageAnnounce) (err error) {

	// check that source can pay deductable

	// check that destination exists

	// check that DHT loc is online (make a request through the DHT)

	// check that the signature is valid

	return nil
}

func (b *BlockChain) AnnounceStorage(a StorageAnnounce) (err error) {
	err = b.VerifyAnnouceStorage(a)
	if err != nil {
		return err
	}

	// add free space to host binSortTree

	return nil
}

func (b *BlockChain) VerifyRentStorage(r StorageRent) (err error) {

	// check that source can afford to host the file

	// check that there is enough space on the network for the file + redundancy (??) { files may be atomic in size - 1 MB each or something, and large files just need to be split up

	// check that there are enough servers on the network

	// check that the ID is not being used by the wallet already

	// read expiration

	// check that the signature is valid

	return nil
}

func (b *BlockChain) RentStorage(r StorageRent) (err error) {
	err = b.VerifyRentStorage(r)
	if err != nil {
		return err
	}

	// file doesn't go into the block chain right away, we need to know the block
	// after this block before we can determine where the file belongs

	// so instead this trade is going to go into a queue that will be processed later
	// when the block after next is being processed

	// put rent into queue

	return
}
