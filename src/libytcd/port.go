package libytcd

import (
	"crypto/ecdsa"
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

type KeyError struct {
	Id    string
	Key   *ecdsa.PrivateKey
	error chan error
}

type Port interface {
	AddServer(s *Server)
	AddTransaction(transaction libGFC.Update)
	AddBlock(block []libGFC.Update)
}
