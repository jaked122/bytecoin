package libytcd

import (
	"libGFC"
)

type Server struct {
	ports       []Port
	transaction chan TransactionError
	blocks      chan BlockError
	state       *libGFC.GFCChain
}

func NewServer(ports []Port) (s *Server) {
	s = new(Server)
	s.state = libGFC.NewGFCChain()

	s.blocks = make(chan BlockError)
	s.transaction = make(chan TransactionError)

	go s.handleChannels()

	for _, p := range ports {
		s.AddPort(p)
	}

	return
}

func (s *Server) AddPort(port Port) {
	s.ports = append(s.ports, port)
	port.AddTransactionChannel(s.transaction)
	port.AddBlockChannel(s.blocks)
}

func (s *Server) handleChannels() {
	for {
		select {
		case c := <-s.transaction:
			update := c.BlockMessage
			err := update.Verify(s.state)
			if err != nil {
				c.error <- err
			} else {
				c.error <- nil
				for _, p := range s.ports {
					if p != c.Source {
						p.AddTransaction(update)
					}
				}
			}
		case block := <-s.blocks:
			var err error = nil
			for _, v := range block.BlockMessage {
				err = v.Verify(s.state)
				if err != nil {
					break
				}
				v.Apply(s.state)
			}
			block.error <- err

			for _, p := range s.ports {
				if p != block.Source {
					p.AddBlock(block.BlockMessage)
				}
			}
		}
	}
}
