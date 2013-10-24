package libytcd

type Account string
type YTCVolume uint64
type SignedTransaction string

type Transaction struct {
	Source      Account
	Destination Account
	Amount      YTCVolume
	Signature   SignedTransaction
}

type BlockChain struct {
	state        map[Account]YTCVolume
	transactions []Transaction
}

func NewBlockChain() (b *BlockChain) {
	b = new(BlockChain)
	b.transactions = make([]Transaction, 100)
	b.state = make(map[Account]YTCVolume)
	return
}

func (b *BlockChain) AddTransaction(t Transaction) {
	b.transactions = append(b.transactions, t)
	b.state[t.Source] -= t.Amount
	b.state[t.Destination] += t.Amount
}
