package libytcd

type Account string
type YTCVolume uint64
type SignedTransaction string

type Transaction struct {
    Source Account
    Destination Account
    Amount YTCVolume
    Signature SignedTransaction
}
