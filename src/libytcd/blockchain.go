package libytcd

import ()

type FileSpace uint64
type FileHash string
type HostKey string
type YTCAmount uint64
type Location string
type Proofs string
type Address string
type Signature string

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
	Hosts map[HostKey]HostRecord
	Files map[FileHash]FileRecord
}

func NewState() (s *State) {
	s = new(State)
	s.Hosts = make(map[HostKey]HostRecord)
	s.Hosts["hard"] = HostRecord{0, 0, nil, 10, "", "hard"}
	s.Files = make(map[FileHash]FileRecord)
	return
}
