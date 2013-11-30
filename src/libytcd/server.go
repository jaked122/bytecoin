package libytcd

import (
	"crypto/ecdsa"
	"libGFC"
	"net"
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
	key, host := libGFC.OriginHostRecord()
	s.Keys[host.Id] = key

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

			s.state.Revision += 1
			s.SeenTransactions = make(map[string]libGFC.Update)

		case key := <-s.KeyChannel:
			s.Keys[key.Id] = key.Key
			key.error <- nil

		case _ = <-s.calculateBlock:

			if _, found := s.Keys[s.state.NextHost().Id]; found {

				block := s.produceBlock()

				c := make(chan error)
				go func() {
					s.BlockChannel <- BlockError{block, nil, c}
					<-c
				}()
			}
		}

		if s.event != nil {
			s.event <- true
		}
	}
}

func (s *Server) GetLocation() (location string) {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		ip, _, _ := net.ParseCIDR(a.String())
		if ip.IsLoopback() {
			continue
		}
		//Ignore IPv6 for now
		if ip.To4() == nil {
			continue
		}
		location = ip.String()
	}
	return
}

func (s *Server) produceBlock() (block []libGFC.Update) {

	//Find our entry
	var location *libGFC.FileChainRecord = nil
	for _, l := range s.state.State {
		if l.Location == s.GetLocation() {
			location = l
			break
		}
	}

	//If we aren't in the map, add us
	if location == nil {
		key, r := libGFC.NewHost(s.GetLocation())
		t := libGFC.NewHostUpdate(r)
		t.Sign(key)
		location = r
		s.Keys[r.Id] = key
		s.SeenTransactions[t.String()] = t
	}

	//If we are bootstrapping, destroy the default entry
	if s.state.Revision == 0 {
		key, r := libGFC.OriginHostRecord()
		t := libGFC.NewTransferUpdate(r.Id, location.Id, r.Balance)
		t.Sign(key)
		s.SeenTransactions[t.String()] = t
	}

	block = make([]libGFC.Update, len(s.SeenTransactions))

	i := uint(0)
	for _, v := range s.SeenTransactions {
		block[i] = v
		i++
	}

	return

}
