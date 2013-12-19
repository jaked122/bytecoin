package libytc

import (
	"encoding/json"
	"testing"
)

func TestHostKeyEncoding(t *testing.T) {
	_, h := DeterministicKey(0)

	b, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))

	n := HostKey{}
	err = json.Unmarshal(b, &n)
	if err != nil {
		t.Fatal(err)
	}

	if n.PublicKey == nil {
		t.Fatal(n)
	}
}
