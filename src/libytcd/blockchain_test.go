package libytcd

import (
	"testing"
)

func BlockchainTest(t *testing.T) {
	s := NewState()
	if s.Hosts["hard"].Balance <= 0 {
		t.Error("Balance not present")
	}
}
