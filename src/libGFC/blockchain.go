package libGFC

import (
	"crypto/ecdsa"
	"crypto/rand"
	"libytcd"
)

type FileChainRecord struct {
	Id          string
	Balance     uint64
	Location    string
	FreeSpace   uint64 ///bytes
	TakenSpace  uint64
	RentedSpace uint64
	keyList     []libytcd.HostKey
}

func NewHost(location string) (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = RandomIdString()
	host.Location = location
	private, public := libytcd.NewKey(rand.Reader)
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

func NewGFCChain() (g *GFCChain) {
	g = new(GFCChain)
	return
}
