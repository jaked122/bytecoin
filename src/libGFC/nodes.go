package libGFC

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"libytc"
)

type FileChainRecord struct {
	Id          string
	Balance     uint64
	Location    []string
	FreeSpace   uint64 ///bytes
	TakenSpace  uint64
	RentedSpace uint64
	KeyList     map[string]float64
}

func NewHost(location string) (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = RandomIdString()
	host.Location = []string{location}
	private, public := libytc.RandomKey()
	host.KeyList = make(map[string]float64)
	host.KeyList[public.Hash()] = 1.0
	return
}

func NewFile(filesize uint64) (file *FileChainRecord) {
	file = new(FileChainRecord)
	file.Id = RandomIdString()
	file.RentedSpace = filesize
	file.KeyList = make(map[string]float64)
	return
}

func OriginHostRecord() (private *ecdsa.PrivateKey, host *FileChainRecord) {
	host = new(FileChainRecord)
	host.Id = "Origin"
	host.Location = []string{"127.0.0.1"}
	host.Balance = 1e8
	private, public := libytc.DeterministicKey(0)
	host.KeyList = make(map[string]float64)
	host.KeyList[public.Hash()] = 1.0
	return
}

func (f FileChainRecord) Verify(content string, signatures *libytc.SignatureMap) (err error) {
	i := float64(0)

	for key, signature := range signatures.M {
		if err = libytc.Verify(content, key, signature); err == nil {
			v, ok := f.KeyList[key.Hash()]
			if !ok {
				err = errors.New("key != key?")
				return
			}
			i += v
		} else {
			return
		}
	}

	if i < 1.0 {
		e := fmt.Sprintf("Consensus not reached i=%f\n", i)
		e += fmt.Sprint(*signatures)
		e += "\n"
		e += fmt.Sprint(f.KeyList)
		err = errors.New(e)
	}
	return
}
