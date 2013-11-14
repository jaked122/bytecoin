package libytcd

import (
	"libGFC"
)

type TransactionError struct {
	BlockMessage libGFC.Update
	error        chan error
}

type BlockError struct {
	BlockMessage []libGFC.Update
	error        chan error
}

type Port interface {
	AddTransactionChannel(transaction chan TransactionError)
	AddBlockChannel(block chan BlockError)
}
