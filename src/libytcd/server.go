package libytcd

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}

	t := new(Transaction)
	err = json.Unmarshal(b, t)
	if err != nil {
		log.Print(err)
	}

	// currently n is returning as '0', which suggests that the body is empty?
	// will debug later

	// error checking on request
	// t := {http request body}
	// error checking on t
	// verifying signature on t
	y.b.AddTransaction(*t)
	// send non error response?
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
