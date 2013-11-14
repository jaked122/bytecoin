package libytc

import (
	"io"
	"math/rand"
)

type dummy struct {
	*rand.Rand
}

func (d *dummy) Read(p []byte) (n int, err error) {
	u := d.Rand.Uint32()
	for n = 0; n < len(p) && n < 4; n++ {
		p[n] = byte(u >> (8 * uint(n)))
	}

	return
}

func MakeReader(r *rand.Rand) (i io.Reader) {
	i = &dummy{r}
	return

}
