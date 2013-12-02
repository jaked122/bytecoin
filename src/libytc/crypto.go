package libytc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	mathrand "math/rand"
)

type HostKey struct {
	ecdsa.PublicKey
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func (h HostKey) String() (str string) {
	str = fmt.Sprint(h.PublicKey)
	return
}

func (h HostKey) Hash() (str string) {
	hash := sha512.New()
	str = hex.EncodeToString(hash.Sum([]byte(h.String())))
	return
}

func (h HostKey) MarshalText() (text []byte, err error) {
	a := struct {
		X, Y *big.Int
	}{h.PublicKey.X, h.PublicKey.Y}
	return json.Marshal(a)
}

func (h HostKey) UnmarshalText(text []byte) (err error) {
	a := struct {
		X, Y *big.Int
	}{}
	err = json.Unmarshal(text, &a)

	h.PublicKey.X = a.X
	h.PublicKey.Y = a.Y
	h.Curve = elliptic.P521()
	return
}

func OriginKey() (priv *ecdsa.PrivateKey, host HostKey) {
	return DeterministicKey(3021)
}

func DeterministicKey(i int64) (priv *ecdsa.PrivateKey, host HostKey) {
	source := MakeReader(mathrand.New(mathrand.NewSource(i)))
	priv, host = NewKey(source)
	return
}

func NewKey(source io.Reader) (priv *ecdsa.PrivateKey, host HostKey) {
	priv, err := ecdsa.GenerateKey(elliptic.P521(), source)
	if err != nil {
		log.Fatal(err.Error())
	}

	host.PublicKey = priv.PublicKey
	return
}

func RandomKey() (priv *ecdsa.PrivateKey, host HostKey) {
	source := cryptorand.Reader
	priv, host = NewKey(source)
	return
}
