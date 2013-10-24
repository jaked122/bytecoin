package libytcd

type Account [64]byte // 512 bits
type YTCVolume uint64
type SignedTransaction [64]byte // 512 bits

type Transaction struct {
	Source      Account
	Destination Account
	Amount      YTCVolume
	Signature   SignedTransaction
}
