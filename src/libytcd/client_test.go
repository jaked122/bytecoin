package libytcd

import (
	"log"
	"testing"
)

func TestClientApiSimple(t *testing.T) {
	api := NewApiPort()
	go func() {
		err := api.Listen(":1337")
		if err != nil {
			log.Fatal(err)
		}
	}()

	b := make([]Port, 1)
	b[0] = api
	_ = NewServer(b)

	s := GetAddress("localhost:1337")

	if len(s) <= 0 {
		t.Fatal(s)
	}
}
