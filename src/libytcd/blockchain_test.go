package libytcd

import (
	"testing"
)

func BlockchainTest(t *testing.T) {
	b := NewBlockChain()
	b.AddTransaction(Transaction{"source", "dest", 10, "sig"})

}
