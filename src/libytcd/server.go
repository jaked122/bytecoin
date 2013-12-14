package libytcd

import (
	"crypto/ecdsa"
	"libGFC"
	"libytc"
	"net"
	"time"
)

type Server struct {
	ports              []Port
	state              map[string]*libGFC.GFCChain
	encoder            map[string]libytc.Encoder
	TransactionChannel chan TransactionError
	BlockChannel       chan BlockError
	KeyChannel         chan KeyError
	event              chan bool
	calculateBlock     <-chan time.Time
	SeenTransactions   map[string]libytc.Update
	Keys               map[string]*ecdsa.PrivateKey
}

func NewServer(ports []Port) (s *Server) {
	s = new(Server)
	s.state = make(map[string]*libGFC.GFCChain)
	s.state["GFC"] = libGFC.NewGFCChain()

	s.encoder = make(map[string]libytc.Encoder)
	s.encoder["GFC"] = libGFC.GFCEncoder{}

	s.BlockChannel = make(chan BlockError)
	s.TransactionChannel = make(chan TransactionError)
	s.KeyChannel = make(chan KeyError)
	s.SeenTransactions = make(map[string]libytc.Update)
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
			err := update.Verify(s.state[update.Chain()])
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
			chain := "GFC"
			for _, v := range block.BlockMessage.Updates() {
				err = v.Verify(s.state[v.Chain()])
				if err != nil {
					break
				}
				v.Apply(s.state[v.Chain()])
			}

			block.error <- err

			for _, p := range s.ports {
				if p != block.Source {
					p.AddBlock(block.BlockMessage)
				}
			}

			s.state[chain].Revision += 1
			s.SeenTransactions = make(map[string]libytc.Update)

		case key := <-s.KeyChannel:
			s.Keys[key.Id] = key.Key
			key.error <- nil

		case _ = <-s.calculateBlock:

			if _, found := s.Keys[s.state["GFC"].NextHost().Id]; found {

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

func (s *Server) produceBlock() (block libytc.Block) {

	//Find our entry
	var location *libGFC.FileChainRecord = nil
	for _, l := range s.state["GFC"].State {
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
	if s.state["GFC"].Revision == 0 {
		key, r := libGFC.OriginHostRecord()
		t := libGFC.NewTransferUpdate(r.Id, location.Id, r.Balance)
		t.Sign(key)
		s.SeenTransactions[t.String()] = t
	}

	update := make([]libytc.Update, len(s.SeenTransactions))

	i := uint(0)
	for _, v := range s.SeenTransactions {
		update[i] = v
		i++
	}

	block = libGFC.NewGFCBlock(s.state["GFC"].Revision+1, update)
	return

}

func (s *Server) EncodeUpdate(transaction libytc.Update) []byte {
	return s.encoder[transaction.Chain()].EncodeUpdate(transaction)
}

func (s *Server) EncodeBlock(block libytc.Block) []byte {
	return s.encoder[block.Chain()].EncodeBlock(block)
}

func (s *Server) DecodeUpdate(b []byte, chain string) libytc.Update {
	return s.encoder[chain].DecodeUpdate(b)
}

func (s *Server) DecodeBlock(b []byte, chain string) libytc.Block {
	return s.encoder[chain].DecodeBlock(b)
}
