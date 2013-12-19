package libGFC

import (
	"sort"
)

type GFCChain struct {
	State    map[string]*FileChainRecord
	Revision uint64
}

func NewGFCChain() (g *GFCChain) {
	g = new(GFCChain)
	g.State = make(map[string]*FileChainRecord)
	g.Revision = 0

	_, s := OriginHostRecord()
	update := NewHostUpdate(s)
	update.Apply(g)
	return
}

func (g *GFCChain) NextHost() *FileChainRecord {
	i := uint64(0)
	hosts := make([]*FileChainRecord, len(g.State))
	for _, host := range g.State {
		hosts[i] = host
		i++
	}

	sort.Sort(Hosts(hosts))

	i = 0
	for {
		for _, host := range hosts {
			if host.Balance == 0 {
				continue
			}
			if i >= g.Revision {
				return host
			}
			i++
		}
	}
}

type Hosts []*FileChainRecord

func (s Hosts) Len() int {
	return len(s)
}

func (s Hosts) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func (s Hosts) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}
