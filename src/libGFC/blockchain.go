package libGFC

import (
	"crypto/ecdsa"
	"libytc"
)

type FileChainRecord struct {
	Id          string
	Balance     uint64
	Location    string
	FreeSpace   uint64 ///bytes
	TakenSpace  uint64
	RentedSpace uint64
	keyList     []libytc.HostKey
}

func NewHost(location string) (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = RandomIdString()
	host.Location = location
	private, public := libytc.RandomKey()
	host.keyList = append(host.keyList, public)
	return
}

func NewFile(filesize uint64) (file *FileChainRecord) {
	file = new(FileChainRecord)
	file.Id = RandomIdString()
	file.RentedSpace = filesize
	return
}

type GFCChain struct {
	State map[string]*FileChainRecord
}

func OriginHostRecord() (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = "Origin"
	host.Location = "127.0.0.1"
	host.Balance = 1e8
	private, public := libytc.DeterministicKey(0)
	host.keyList = append(host.keyList, public)
	return
}

func NewGFCChain() (g *GFCChain) {
	g = new(GFCChain)
	g.State = make(map[string]*FileChainRecord)

	_, s := OriginHostRecord()
	update := NewHostUpdate(s)
	update.Apply(g)
	return
}
