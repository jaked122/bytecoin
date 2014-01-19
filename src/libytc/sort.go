package libytc

type HostAlphabetically []*HostKey

func (h HostAlphabetically) Len() int {
	return len(h)
}

func (h HostAlphabetically) Less(i, j int) bool {
	return h[i].PublicKey.X.Cmp(h[j].PublicKey.X) < 0
}

func (h HostAlphabetically) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
