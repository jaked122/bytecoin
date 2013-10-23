package libytcd

type Account string

type YTC uint64

type Transaction struct {
	Source      Account
	Destination Account
	Amount      YTC
}
