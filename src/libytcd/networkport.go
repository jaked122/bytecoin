package libytcd

import (
	"encoding/json"
	"libGFC"
	"net"
)

type NetworkPort struct {
	neighbors []*json.Encoder
}

func NewNetworkPort() (y *NetworkPort) {
	y = new(NetworkPort)
	return
}

func (y *NetworkPort) HandleNetworkConnection(c net.Conn) {
	je := json.NewEncoder(c)
	y.neighbors = append(y.neighbors, je)

	j := json.NewDecoder(c)
	for {
		v := make([]libGFC.BlockMessage, 1)
		j.Decode(&v)
	}
}

func (y *NetworkPort) ListenNetwork(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go y.HandleNetworkConnection(c)
	}
}

func (y *NetworkPort) ConnectAddress(addr string) (err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	go y.HandleNetworkConnection(c)
	return
}
