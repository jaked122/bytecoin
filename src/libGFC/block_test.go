package libGFC

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	priv, record := NewHost("foo")

	update := NewHostUpdate(record)
	update.Sign(priv)
	v := make([]Update, 0)
	v = append(v, update)

	arr := EncodeUpdates(v)
	v = DecodeUpdates(arr)
	if v[0].(*HostUpdate).Record.Location != "foo" {
		t.Fatal("Location is not foo")
	}
}
