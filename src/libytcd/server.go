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
	state        map[Account]YTCVolume
	transactions []Transaction
}

func NewYtcd() (y *ytcServer) {
	y = new(ytcServer)
	return
}

// Will eventually load html explaining how to get started?
func loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "learning how to parse http requests")
	// error checking on request
	// t := {http request body}
	// error checking on t
	// verifying signature on t
	// AddTransaction(t)
	// send non error response?
}

func (y *ytcServer) AddTransaction(t Transaction) {
	y.transactions = append(y.transactions, t)
	y.state[t.Source] -= t.Amount
	y.state[t.Destination] += t.Amount
}

func (y *ytcServer) Listen(addr string) (err error) {

	d := http.NewServeMux()
	d.HandleFunc(root, loadHomepage)
	d.HandleFunc(postTransaction, handleTransaction)

	s := &http.Server{
		Addr:           addr,
		Handler:        d,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
