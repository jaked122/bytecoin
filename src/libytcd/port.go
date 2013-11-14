package libytcd

import (
	"libGFC"
)

type MessageError struct {
	BlockMessage libGFC.BlockMessage
	error        chan bool
}

type BlockError struct {
	BlockMessage libGFC.BlockMessage
	error        chan bool
}

type Port interface {
	AddTransactionChannel(transaction chan MessageError)
	AddBlockChannel(block chan BlockError)
}
