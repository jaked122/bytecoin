package libGFC

import (
	"libytc"
	"testing"
)

func TestEncoding(t *testing.T) {
	g := GFCEncoder{}
	priv, record := NewHost("foo")

	update := NewHostUpdate(record)
	update.Sign(priv)
	v := make([]libytc.Update, 0)
	v = append(v, update)

	arr := g.EncodeBlock(NewGFCBlock(0, v))
	v = g.DecodeBlock(arr).Updates()
	if v[0].(*HostUpdate).Record.Location[0] != "foo" {
		t.Fatal("Location is not foo")
	}
}
