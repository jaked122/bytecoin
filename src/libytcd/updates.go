package libytcd

type Update interface {
	Verify(s *State) bool
	Apply(s *State)
}

type TransferUpdate struct {
	Type        string
	Source      HostKey
	Destination HostKey
	Amount      YTCAmount
	Signature   Signature
	//Public Key? Check bitcoin
}

func (t *TransferUpdate) Verify(s *State) bool {

	_, found := s.Hosts[t.Source]
	if !found {
		return false
	}

	_, found = s.Hosts[t.Destination]
	if !found {
		return false
	}

	if t.Amount > s.Hosts[t.Source].Balance {
		return false
	}

	//Verify Signature
	return true
}

func (t *TransferUpdate) Apply(s *State) {
	source := s.Hosts[t.Source]
	source.Balance -= t.Amount
	dest := s.Hosts[t.Destination]
	dest.Balance += t.Amount
	return
}

func NewTransferUpdate() (t *TransferUpdate) {
	t = new(TransferUpdate)
	t.Type = "TransferUpdate"
	return
}

type HostUpdate struct {
	Type      string
	Key       HostKey
	Signature Signature
}

func NewHostUpdate() (h *HostUpdate) {
	h = new(HostUpdate)
	h.Type = "HostUpdate"
	return
}

func (t *HostUpdate) Verify(s *State) bool {
	// Verify signature from old key
	o, found := s.Hosts[t.Key]
	if found && o.Key != t.Key {
		return false
	}

	//verify Signature
	return true
}

func (t *HostUpdate) Apply(s *State) {
	host := s.Hosts[t.Key]
	host.Key = t.Key
	return
}
