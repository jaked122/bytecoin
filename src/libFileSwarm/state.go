package libFileSwarm

type State struct {
	swarmid            string
	corehostcount      uint64
	corehostredundancy uint64
	growthhostcount    uint64
	totalspace         uint64

	corehosthashes   map[string][]string
	growthhosthashes map[string][]string

	previousblocks []*Block
	currentblock   *Block
}

type Block struct {
	blockNumber uint64
	blockHash   string

	entropyhash   map[string]string
	entropystring map[string]string

	storagehash   map[string]string
	storagestring map[string]string

	incomingsignals []*Signal
	outgoinsignals  []*Signal

	hostsignatures map[string]string
	indictments    []*Indictment
}

type Signal struct {
}

type Indictment struct {
}
