package libytcd

import (
	"testing"
	"time"
)

func BlockchainTest(t *testing.T) {
	b := NewBlockChain()

	// not sure why the Time() typecast is needed, but `make test` fails without it
	b.AddTransaction(Transaction{"source", "dest", 10, Time(time.Now()), "sig"})
}
