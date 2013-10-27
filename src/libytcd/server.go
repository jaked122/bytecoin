package libytcd

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	root            = "/"
	postTransaction = "/postTransaction"
	newWallet       = "/newWallet"
	sendMoney       = "/sendMoney"
)

type ytcServer struct {
	s *State
	d *http.ServeMux
}

func NewYtcd() (y *ytcServer) {
	y = new(ytcServer)
	y.s = NewState()

	y.d = http.NewServeMux()
	y.d.HandleFunc(root, y.loadHomepage)
	y.d.HandleFunc(newWallet, y.newWallet)
	y.d.HandleFunc(sendMoney, y.sendMoney)

	return
}

// Will eventually load html explaining how to get started?
func (y *ytcServer) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (y *ytcServer) newWallet(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, 8)
	rand.Read(b)

	h := NewHostUpdate()
	h.Key = HostKey(hex.EncodeToString(b))
	h.Signature = "Unimplemented"

	h.Verify(y.s)
	h.Apply(y.s)

	io.WriteString(w, string(h.Key))
}

func (y *ytcServer) sendMoney(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	v := make(map[string]string)
	_ = json.Unmarshal(b, v)

	source := v["Source"]
	destination := v["Destination"]
	amount, _ := strconv.ParseUint(v["Amount"], 10, 64)

	t := NewTransferUpdate()
	t.Source = HostKey(source)
	t.Destination = HostKey(destination)
	t.Amount = YTCAmount(amount)
	t.Signature = Signature("fu")

	ok := t.Verify(y.s)
	if !ok {
		io.WriteString(w, "false")
		return
	}
	t.Apply(y.s)
	io.WriteString(w, "true")
	return
}

func (y *ytcServer) HandleNetworkConnection(c net.Conn) {
	j := json.NewDecoder(c)
	for {
		var v interface{}
		j.Decode(v)

		b, err := json.Marshal(v)
		if err != nil {
			break
		}

		update, err := ParseUpdate(b)
		if err != nil {
			break
		}
		update.Verify(y.s)
		update.Apply(y.s)
	}
}

func (y *ytcServer) ListenNetwork(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go y.HandleNetworkConnection(c)
	}
}

func (y *ytcServer) ListenCtl(addr string) (err error) {
	s := &http.Server{
		Addr:           addr,
		Handler:        y.d,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
