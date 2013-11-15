package libytcd

import (
	"libGFC"
)

type TransactionError struct {
	BlockMessage libGFC.Update
	Source       Port
	error        chan error
}

type BlockError struct {
	BlockMessage []libGFC.Update
	Source       Port
	error        chan error
}

type Port interface {
	AddTransactionChannel(transaction chan TransactionError)
	AddBlockChannel(block chan BlockError)
	AddTransaction(transaction libGFC.Update)
	AddBlock(block []libGFC.Update)
}
