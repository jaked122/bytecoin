package libGFC

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func RandomIdString() (id string) {
	b := make([]byte, 4)
	n, err := rand.Read(b)
	if err != nil || n != 4 {
		log.Fatal(err)
	}
	id = hex.EncodeToString(b)
	return
}
