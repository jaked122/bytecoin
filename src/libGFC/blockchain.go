package libGFC

import (
	"crypto/ecdsa"
	"libytc"
	"sort"
)

type FileChainRecord struct {
	Id          string
	Balance     uint64
	Location    []string
	FreeSpace   uint64 ///bytes
	TakenSpace  uint64
	RentedSpace uint64
	KeyList     []libytc.HostKey
}

func NewHost(location string) (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = RandomIdString()
	host.Location = []string{location}
	private, public := libytc.RandomKey()
	host.KeyList = append(host.KeyList, public)
	return
}

func NewFile(filesize uint64) (file *FileChainRecord) {
	file = new(FileChainRecord)
	file.Id = RandomIdString()
	file.RentedSpace = filesize
	return
}

func OriginHostRecord() (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = "Origin"
	host.Location = []string{"127.0.0.1"}
	host.Balance = 1e8
	private, public := libytc.DeterministicKey(0)
	host.KeyList = append(host.KeyList, public)
	return
}

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
