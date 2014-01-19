package libFileSwarm

import (
	"crypto/ecdsa"
	"errors"
	"libytc"
)

type HeartbeatUpdate struct {
	HostId         string
	SwarmId        string
	EntropyHash    string
	EntropyString  string
	BlockSignature *libytc.SignatureMap
	Signature      *libytc.SignatureMap
}

func (h *HeartbeatUpdate) Verify(i interface{}) (err error) {
	s := i.(*State)

	_, found := s.piecemapping[h.HostId]
	if !found {
		return errors.New("No such host in file swarm")
	}

	host, found := s.swarmtracker.State[h.HostId]
	if !found {
		return errors.New("Host missing from swarmtracker")
	}

	hash := s.currentblock.entropyhash[h.HostId]
	if libytc.StringHash(h.EntropyString) != hash {
		return errors.New("Host entropy string does not match hash")
	}

	return host.Verify(h.String(), h.Signature)
}

func (h *HeartbeatUpdate) Sign(key *ecdsa.PrivateKey) {
	h.Signature = libytc.SignMap(nil, h.String(), key)
}

func (h *HeartbeatUpdate) Apply(i interface{}) {
	s := i.(*State)
	s = s

}

func (h *HeartbeatUpdate) Chain() string {
	return "FileSwarm"
}

func (h *HeartbeatUpdate) ChainId() string {
	return h.SwarmId
}

func (h *HeartbeatUpdate) String() (str string) {
	str = "HeartbeatUpdate\n"
	str += "HostId:" + h.HostId + "\n"
	str += "SwarmId:" + h.SwarmId + "\n"
	str += "EntropyHash:" + h.EntropyHash + "\n"
	str += "EntropyString:" + h.EntropyString + "\n"
	str += "BlockSignature:" + h.BlockSignature.String() + "\n"
	return str
}
