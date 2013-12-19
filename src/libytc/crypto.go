package libytc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	mathrand "math/rand"
)

type HostKey struct {
	PublicKey *ecdsa.PublicKey
}

type Signature struct {
	R *big.Int
	S *big.Int
}

func Sign(str string, priv *ecdsa.PrivateKey) Signature {
	r, s, err := ecdsa.Sign(cryptorand.Reader, priv, []byte(str))
	if err != nil {
		log.Fatal(err)
	}
	return Signature{r, s}
}

func SignMap(m *SignatureMap, str string, priv *ecdsa.PrivateKey) *SignatureMap {
	if m == nil {
		m = NewSignatureMap(1)
	}

	(m.M)[&HostKey{&priv.PublicKey}] = Sign(str, priv)
	return m
}

func Verify(str string, pub *HostKey, s Signature) (err error) {
	if !ecdsa.Verify(pub.PublicKey, []byte(str), s.R, s.S) {
		err = errors.New("Invalid Signature")
	}
	return
}

func (h *HostKey) String() (str string) {
	str = fmt.Sprint(h.PublicKey)
	return
}

func (h *HostKey) Hash() (str string) {
	hash := sha512.New()
	str = hex.EncodeToString(hash.Sum([]byte(h.String())))
	return
}

func (h *HostKey) MarshalJSON() (text []byte, err error) {
	a := struct {
		X, Y *big.Int
	}{h.PublicKey.X, h.PublicKey.Y}
	return json.Marshal(a)
}

func (h *HostKey) UnmarshalJSON(text []byte) (err error) {
	a := &struct {
		X, Y *big.Int
	}{}
	err = json.Unmarshal(text, a)
	if err != nil {
		return
	}

	h.PublicKey = &ecdsa.PublicKey{elliptic.P521(), a.X, a.Y}

	return
}

func OriginKey() (priv *ecdsa.PrivateKey, host *HostKey) {
	return DeterministicKey(3021)
}

func DeterministicKey(i int64) (priv *ecdsa.PrivateKey, host *HostKey) {
	source := MakeReader(mathrand.New(mathrand.NewSource(i)))
	priv, host = NewKey(source)
	return
}

func NewKey(source io.Reader) (priv *ecdsa.PrivateKey, host *HostKey) {
	priv, err := ecdsa.GenerateKey(elliptic.P521(), source)
	if err != nil {
		log.Fatal(err.Error())
	}

	host = new(HostKey)
	host.PublicKey = &priv.PublicKey
	return
}

func RandomKey() (priv *ecdsa.PrivateKey, host *HostKey) {
	source := cryptorand.Reader
	priv, host = NewKey(source)
	return
}
