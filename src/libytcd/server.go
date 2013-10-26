package libytcd

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	root            = "/"
	postTransaction = "/postTransaction"
	newWallet       = "/newWallet"
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

	return
}

// Will eventually load html explaining how to get started?
func (y *ytcServer) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (y *ytcServer) newWallet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "learning how to parse http requests")

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
