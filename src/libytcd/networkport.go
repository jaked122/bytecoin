package libytcd

import (
	"encoding/json"
	"libGFC"
	"net"
)

type MessageFormat struct {
	Type    string
	Payload []byte
}

type NetworkConnection struct {
	outbound    *json.Encoder
	inbound     *json.Decoder
	block       chan BlockError
	transaction chan TransactionError
}

func NewNetworkConnection(c net.Conn) (n *NetworkConnection) {
	n = new(NetworkConnection)
	n.outbound = json.NewEncoder(c)
	n.inbound = json.NewDecoder(c)

	go n.HandleNetworkConnection()
	return
}

func (n *NetworkConnection) AddTransactionChannel(transaction chan TransactionError) {
	n.transaction = transaction
}

func (n *NetworkConnection) AddBlockChannel(block chan BlockError) {
	n.block = block
}

func (n *NetworkConnection) AddBlock(block []libGFC.Update) {
	msg := new(MessageFormat)
	msg.Type = "Block"
	msg.Payload = libGFC.EncodeUpdates(block)
	n.outbound.Encode(msg)
}

func (n *NetworkConnection) AddTransaction(transaction libGFC.Update) {
	msg := new(MessageFormat)
	msg.Type = "Transaction"
	msg.Payload = libGFC.EncodeUpdate(transaction)
	n.outbound.Encode(msg)
}

func (n *NetworkConnection) HandleNetworkConnection() {
	for {
		v := new(MessageFormat)
		n.inbound.Decode(v)
		switch v.Type {
		case "Transaction":
			t := libGFC.DecodeUpdate(v.Payload)
			c := make(chan error)
			n.transaction <- TransactionError{t, n, c}
			_ = <-c
		case "Block":
			b := libGFC.DecodeUpdates(v.Payload)
			c := make(chan error)
			n.block <- BlockError{b, n, c}
			_ = <-c
		}
	}
}

type NetworkPort struct {
	s *Server
}

func NewNetworkPort(s *Server) (n *NetworkPort) {
	n = new(NetworkPort)
	n.s = s
	return
}

func (n *NetworkPort) ListenNetwork(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		_ = NewNetworkConnection(c)
	}
}

func (n *NetworkPort) ConnectAddress(addr string) (err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	_ = NewNetworkConnection(c)
	return
}
