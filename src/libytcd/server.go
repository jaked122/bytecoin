package libytcd

import (
	"io"
	"net/http"
	"time"
)

const (
	root            = "/"
	postTransaction = "/postTransaction"
)

type ytcServer struct {
	b *BlockChain
	d *http.ServeMux
}

func NewYtcd() (y *ytcServer) {
	y = new(ytcServer)
	y.b = NewBlockChain()

	y.d = http.NewServeMux()
	y.d.HandleFunc(root, y.loadHomepage)
	y.d.HandleFunc(postTransaction, y.handleTransaction)

	return
}

// Will eventually load html explaining how to get started?
func (y *ytcServer) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (y *ytcServer) handleTransaction(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "learning how to parse http requests")

}

func (y *ytcServer) Listen(addr string) (err error) {
	s := &http.Server{
		Addr:           addr,
		Handler:        y.d,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
