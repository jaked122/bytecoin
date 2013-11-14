package libytcd

import (
	"encoding/json"
	"libGFC"
	"net"
)

type NetworkPort struct {
	neighbors []*json.Encoder
}

func NewNetworkPort() (n *NetworkPort) {
	n = new(NetworkPort)
	return
}

func (n *NetworkPort) HandleNetworkConnection(c net.Conn) {
	je := json.NewEncoder(c)
	n.neighbors = append(n.neighbors, je)

	j := json.NewDecoder(c)
	for {
		v := make([]libGFC.BlockMessage, 1)
		j.Decode(&v)
	}
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

		go n.HandleNetworkConnection(c)
	}
}

func (n *NetworkPort) ConnectAddress(addr string) (err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	go n.HandleNetworkConnection(c)
	return
}
