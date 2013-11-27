package libytcd

import (
	"encoding/json"
	"libGFC"
	"log"
	"net"
)

type MessageFormat struct {
	Type    string
	Payload []byte
}

type NetworkConnection struct {
	outbound *json.Encoder
	inbound  *json.Decoder
	s        *Server
}

func NewNetworkConnection(c net.Conn) (n *NetworkConnection) {
	n = new(NetworkConnection)
	n.outbound = json.NewEncoder(c)
	n.inbound = json.NewDecoder(c)

	go n.HandleNetworkConnection()
	return
}

func (n *NetworkConnection) AddServer(s *Server) {
	n.s = s
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
			n.s.TransactionChannel <- TransactionError{t, n, c}
			_ = <-c
		case "Block":
			b := libGFC.DecodeUpdates(v.Payload)
			c := make(chan error)
			n.s.BlockChannel <- BlockError{b, n, c}
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
		log.Print("Listened")
		if err != nil {
			return err
		}
		n.s.AddPort(NewNetworkConnection(c))
	}
}

func (n *NetworkPort) ConnectAddress(addr string) (err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	n.s.AddPort(NewNetworkConnection(c))
	return
}
