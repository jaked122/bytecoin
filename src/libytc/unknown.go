package libytc

import (
	"fmt"
	"log"
)

type UnknownUpdate struct {
	payload string
	chain   string
	Type    string
}

func NewUnknownUpdate(payload string, chain string, Type string) (t *UnknownUpdate) {
	t = new(UnknownUpdate)
	t.payload = payload
	t.chain = chain
	t.Type = Type
	return
}

func (t *UnknownUpdate) Verify(i interface{}) (err error) {
	log.Fatal("Cannot Verify Unknown Update")
	return
}

func (t *UnknownUpdate) Apply(i interface{}) {
	log.Fatal("Cannot Apply Unknown Update")
	return
}

func (t *UnknownUpdate) Chain() string {
	return t.chain
}

func (t *UnknownUpdate) String() string {
	return fmt.Sprint(t)
}

func (t *UnknownUpdate) MarshallJSON(b []byte, err error) {
	b = []byte(t.payload)
	return
}
