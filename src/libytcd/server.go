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
	s.ports = ports
	s.state = libGFC.NewGFCChain()

	s.blocks = make(chan BlockError)
	s.transaction = make(chan TransactionError)

	go s.handleChannels()

	for _, p := range ports {
		p.AddTransactionChannel(s.transaction)
		p.AddBlockChannel(s.blocks)
	}

	return
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
				for _, p := range s.ports {
					if p != c.Source {
						p.AddTransaction(update)
					}
				}
			}
		case block := <-s.blocks:
			for _, v := range block.BlockMessage {
				err := v.Verify(s.state)
				if err != nil {
					block.error <- err
					break
				}
				v.Apply(s.state)
			}

			for _, p := range s.ports {
				if p != block.Source {
					p.AddBlock(block.BlockMessage)
				}
			}
		}
	}
}
