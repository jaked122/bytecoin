package libytcd

import (
	"libGFC"
)

type Server struct {
	ports            []Port
	transaction      chan TransactionError
	blocks           chan BlockError
	state            *libGFC.GFCChain
	event            chan bool
	SeenTransactions map[string]bool
}

func NewServer(ports []Port) (s *Server) {
	s = new(Server)
	s.state = libGFC.NewGFCChain()

	s.blocks = make(chan BlockError)
	s.transaction = make(chan TransactionError)
	s.SeenTransactions = make(map[string]bool)

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
	if s.event != nil {
		s.event <- true
	}
}

func (s *Server) handleChannels() {
	for {
		select {
		case c := <-s.transaction:
			if s.event != nil {
				s.event <- true
			}
			update := c.BlockMessage
			_, found := s.SeenTransactions[update.String()]
			if found {
				c.error <- nil
				continue
			}
			err := update.Verify(s.state)
			if err != nil {
				c.error <- err
			} else {
				s.SeenTransactions[update.String()] = true
				c.error <- nil
				for _, p := range s.ports {
					if p != c.Source {
						p.AddTransaction(update)
					}
				}
			}
		case block := <-s.blocks:
			if s.event != nil {
				s.event <- true
			}
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
