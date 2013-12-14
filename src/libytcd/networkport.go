package libytcd

import (
	"encoding/json"
	"libGFC"
	"libytc"
	"log"
	"net"
)

type MessageFormat struct {
	Type    string
	Chain   string
	Payload []byte
}

type NetworkConnection struct {
	outbound *json.Encoder
	inbound  *json.Decoder
	s        *Server
	sync     chan bool
}

func NewNetworkConnection(c net.Conn) (n *NetworkConnection) {
	n = new(NetworkConnection)
	n.outbound = json.NewEncoder(c)
	n.inbound = json.NewDecoder(c)
	n.sync = make(chan bool)

	go n.HandleNetworkConnection()
	return
}

func (n *NetworkConnection) AddServer(s *Server) {
	n.s = s
}

func (n *NetworkConnection) AddBlock(block libytc.Block) {
	msg := new(MessageFormat)
	msg.Type = "Block"
	msg.Payload = n.s.EncodeBlock(block)
	msg.Chain = block.Chain()
	n.outbound.Encode(msg)
}

func (n *NetworkConnection) AddTransaction(transaction libytc.Update) {
	msg := new(MessageFormat)
	msg.Type = "Transaction"
	msg.Payload = n.s.EncodeUpdate(transaction)
	msg.Chain = transaction.Chain()
	n.outbound.Encode(msg)
}

func (n *NetworkConnection) Reconnect(chain string) {
	msg := new(MessageFormat)
	msg.Type = "Reconnect"
	msg.Payload, _ = json.Marshal(n.s.state[chain].Revision)
	msg.Chain = chain
	n.outbound.Encode(msg)
	<-n.sync
}

func (n *NetworkConnection) HandleNetworkConnection() {
	for {
		v := new(MessageFormat)
		n.inbound.Decode(v)
		switch v.Type {
		case "Transaction":
			t := n.s.DecodeUpdate(v.Payload, v.Chain)
			c := make(chan error)
			n.s.TransactionChannel <- TransactionError{t, n, c}
			_ = <-c
		case "Block":
			b := n.s.DecodeBlock(v.Payload, v.Chain)
			c := make(chan error)
			n.s.BlockChannel <- BlockError{b, n, c}
			_ = <-c
		case "Reconnect":
			b := new(uint64)
			json.Unmarshal(v.Payload, b)
			if *b < n.s.state[v.Chain].Revision {
				o := new(MessageFormat)
				o.Type = "State"
				o.Payload, _ = json.Marshal(n.s.state[v.Chain])
				n.outbound.Encode(o)
			}
		case "State":
			b := new(libGFC.GFCChain)
			json.Unmarshal(v.Payload, b)
			if b != nil && b.Revision > n.s.state["GFC"].Revision {
				n.s.state["GFC"] = b
				n.sync <- true
			} else {
				log.Fatal(b)
			}
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
		n.s.AddPort(NewNetworkConnection(c))
	}
}

func (n *NetworkPort) ConnectAddress(addr string) (port *NetworkConnection, err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	port = NewNetworkConnection(c)
	n.s.AddPort(port)
	return
}
