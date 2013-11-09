package libytcd

import (
	"io"
	"math/rand"
)

type dummy struct {
	*rand.Rand
}

func (d *dummy) Read(p []byte) (n int, err error) {
	n = 1
	p[0] = byte(d.Rand.Uint32())
	return
}

func MakeReader(r *rand.Rand) (i io.Reader) {
	i = &dummy{r}
	return

}
