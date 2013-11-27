package libytcd

import (
	"crypto/ecdsa"
	"libGFC"
	"time"
)

type Server struct {
	ports              []Port
	state              *libGFC.GFCChain
	TransactionChannel chan TransactionError
	BlockChannel       chan BlockError
	KeyChannel         chan KeyError
	event              chan bool
	calculateBlock     <-chan time.Time
	SeenTransactions   map[string]libGFC.Update
	Keys               map[string]*ecdsa.PrivateKey
}

func NewServer(ports []Port) (s *Server) {
	s = new(Server)
	s.state = libGFC.NewGFCChain()

	s.BlockChannel = make(chan BlockError)
	s.TransactionChannel = make(chan TransactionError)
	s.KeyChannel = make(chan KeyError)
	s.SeenTransactions = make(map[string]libGFC.Update)
	s.calculateBlock = time.Tick(1 * time.Minute)

	s.Keys = make(map[string]*ecdsa.PrivateKey)
	key, _ := libGFC.OriginHostRecord()
	s.Keys["127.0.0.1"] = key

	go s.handleChannels()

	for _, p := range ports {
		s.AddPort(p)
	}

	return
}

func (s *Server) AddPort(port Port) {
	s.ports = append(s.ports, port)
	port.AddServer(s)
	if s.event != nil {
		s.event <- true
	}
}

func (s *Server) handleChannels() {
	for {
		select {
		case c := <-s.TransactionChannel:
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
				s.SeenTransactions[update.String()] = update
				c.error <- nil
				for _, p := range s.ports {
					if p != c.Source {
						p.AddTransaction(update)
					}
				}
			}
		case block := <-s.BlockChannel:
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
		case key := <-s.KeyChannel:
			s.Keys[key.Id] = key.Key
			key.error <- nil
		case _ = <-s.calculateBlock:
			if _, found := s.Keys[s.state.NextHost().Id]; found {
				block := make([]libGFC.Update, len(s.SeenTransactions))
				i := uint(0)
				for _, v := range s.SeenTransactions {
					block[i] = v
					i++
				}

				c := make(chan error)
				s.BlockChannel <- BlockError{block, nil, c}
				<-c
			}
		}
	}
}
