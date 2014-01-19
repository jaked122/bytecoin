package libytc

import (
	"encoding/json"
	"sort"
)

type SignatureMap struct {
	M map[*HostKey]Signature
}

func NewSignatureMap(l int) (s *SignatureMap) {
	s = new(SignatureMap)
	s.M = make(map[*HostKey]Signature, l)
	return
}

type sigpair struct {
	Key       *HostKey
	Signature Signature
}

func (s *SignatureMap) MarshalJSON() ([]byte, error) {
	v := make([]sigpair, 0, len(s.M))
	for key, sig := range s.M {
		v = append(v, sigpair{key, sig})
	}

	return json.Marshal(v)
}

func (s *SignatureMap) UnmarshalJSON(b []byte) (err error) {
	var p *[]sigpair
	v := make([]sigpair, 0)
	p = &v

	err = json.Unmarshal(b, p)
	if err != nil {
		return
	}
	if s.M == nil {
		s.M = make(map[*HostKey]Signature)
	}

	for _, p := range *p {
		s.M[p.Key] = p.Signature
	}

	return
}

func (s *SignatureMap) String() (str string) {
	str = "SignatureMap\n"

	v := make([]*HostKey, 0, len(s.M))
	for key, _ := range s.M {
		v = append(v, key)
	}

	sort.Sort(HostAlphabetically(v))

	for _, key := range v {
		str += "Signature:" + s.M[key].String()
	}
	return
}
