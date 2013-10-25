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
	announceStorage = "/announceStorage"
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
	y.d.HandleFunc(announceStorage, y.handleAnnounceStorage)

	return
}

// Will eventually load html explaining how to get started?
func (y *ytcServer) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (y *ytcServer) handleAnnounceStorage(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Print(err)
		return
	}

	s := new(StorageAnnounce)
	err = json.Unmarshal(b, s)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Print(err)
		return
	}

	err = y.b.AnnounceStorage(*s)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Print(err)
		return
	}

	// Eventually return details of announce
	io.WriteString(w, "AnnounceStorage Success")
}

func (y *ytcServer) handleTransaction(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error()) // is this secure?
		log.Print(err)
		return
	}

	t := new(Transaction)
	err = json.Unmarshal(b, t)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Print(err)
		return
	}

	err = y.b.AddTransaction(*t)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Print(err)
		return
	}

	// Eventually return details of transaction
	io.WriteString(w, "Transaction Success")
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
