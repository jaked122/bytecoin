package libytcd

import (
    "testing"
)

func TestSimple(t * testing.T) {
    y := NewYtcd()
    go func() {
         err := y.Listen(":1337")
         t.Fatal(err)
    }()
}
