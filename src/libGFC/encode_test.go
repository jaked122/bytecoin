package libGFC

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	g := GFCEncoder{}
	priv, record := NewHost("foo")

	update := NewHostUpdate(record)
	update.Sign(priv)
	v := make([]Update, 0)
	v = append(v, update)

	arr := g.EncodeUpdates(v)
	v = g.DecodeUpdates(arr)
	if v[0].(*HostUpdate).Record.Location != "foo" {
		t.Fatal("Location is not foo")
	}
}
