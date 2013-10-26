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
	ID             HostHash
}

type FileRecord struct {
	Balance  YTCAmount
	Proofs   []Proofs
	Hosts    []HostHash
	ID       FileHash
	Metadata struct{}
}

type State struct {
	Hosts map[HostHash]HostRecord
	Files map[FileHash]FileRecord
}

func NewState() (s *State) {
	s = new(State)
	s.Hosts = make(map[HostHash]HostRecord)
	s.Files = make(map[FileHash]FileRecord)
	return
}

type Updates interface {
	Verify() bool
	Apply(s *State)
}
