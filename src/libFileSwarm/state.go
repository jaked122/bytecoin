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

	indictments []*Indictment
}

type Block struct {
	blockNumber     uint64
	blockHash       string
	entropyhash     map[string]string
	entropystring   map[string]string
	incomingsignals []*Signal
	outgoinsignals  []*Signal
	hosthash        map[string]string
}

type Signal struct {
}

type Indictment struct {
}
