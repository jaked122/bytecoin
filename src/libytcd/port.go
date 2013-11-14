package libytcd

import (
	"libGFC"
)

type Port interface {
	addTransactionChannel(transaction chan BlockMessage)
	addBlockChannel(block chan []BlockMessage)
}
