package libytcd

import (
	"crypto/ecdsa"
	"libytc"
)

type TransactionError struct {
	BlockMessage libytc.Update
	Source       Port
	error        chan error
}

type BlockError struct {
	BlockMessage libytc.Block
	Source       Port
	error        chan error
}

type KeyError struct {
	Id    string
	Key   *ecdsa.PrivateKey
	error chan error
}

type Port interface {
	AddServer(s *Server)
	AddTransaction(transaction libytc.Update)
	AddBlock(block libytc.Block)
}
