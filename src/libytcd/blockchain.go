package libytcd

import ()

type FileSpace uint64
type FileHash string
type HostHash string
type YTCAmount uint64
type Location string
type Proofs string
type Address string

type HostRecord struct {
	SpaceAvailable FileSpace
	SpaceUsed      FileSpace
	StoredFiles    []FileHash
	Balance        YTCAmount
	Location       Address
	Key            HostKey
}

type FileRecord struct {
	Balance  YTCAmount
	Proofs   []Proofs
	Hosts    []HostKey
	Key      FileHash
	Metadata struct{}
}

type State struct {
	Hosts map[HostHash]HostRecord
	Files map[FileHash]FileRecord
}

func NewState() (s *State) {
	s = new(State)
	s.Hosts = make(map[HostHash]HostRecord)
	_, key := OriginKey()
	name := key.Hash()
	s.Hosts[name] = HostRecord{0, 0, nil, 10, "", key}
	s.Files = make(map[FileHash]FileRecord)
	return
}
