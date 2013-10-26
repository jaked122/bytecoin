package libytcd

import (
	"testing"
)

func TestSimple(t *testing.T) {
	y := NewYtcd()
	go func() {
		err := y.ListenCtl(":1337")
		t.Fatal(err)
	}()
}
