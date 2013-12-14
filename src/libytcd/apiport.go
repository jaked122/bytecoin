package libytcd

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"libGFC"
	"libytc"
	"net/http"
	"strconv"
	"time"
)

const (
	apiroot            = "/"
	apipostTransaction = "/postTransaction"
	apinewWallet       = "/newWallet"
	apisendMoney       = "/sendMoney"
)

type ApiPort struct {
	d *http.ServeMux
	s *Server
}

func NewApiPort() (a *ApiPort) {
	a = new(ApiPort)

	a.d = http.NewServeMux()
	a.d.HandleFunc(apiroot, a.loadHomepage)
	a.d.HandleFunc(apinewWallet, a.newWallet)
	a.d.HandleFunc(apisendMoney, a.sendMoney)

	return
}

func (a *ApiPort) AddServer(s *Server) {
	a.s = s
}

func (a *ApiPort) AddTransaction(transaction libytc.Update) {

}

func (a *ApiPort) AddBlock(block libytc.Block) {

}

// Will eventually load html explaining how to get started?
func (a *ApiPort) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (a *ApiPort) newWallet(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, 8)
	rand.Read(b)

	priv, host := libGFC.NewHost("foo")
	h := libGFC.NewHostUpdate(host)
	h.Sign(priv)

	c := make(chan error)

	a.s.KeyChannel <- KeyError{host.Id, priv, c}
	err := <-c
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, err.Error())
		return
	}

	a.s.TransactionChannel <- TransactionError{h, a, c}

	err = <-c
	if err == nil {
		io.WriteString(w, host.Id)
	} else {
		w.WriteHeader(400)
		io.WriteString(w, err.Error())
	}
}

func (a *ApiPort) sendMoney(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	v := make(map[string]string)
	err := json.Unmarshal(b, &v)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	source := v["Source"]
	destination := v["Destination"]
	amount, _ := strconv.ParseUint(v["Amount"], 10, 64)

	t := libGFC.NewTransferUpdate(source, destination, amount)

	c := make(chan error)

	a.s.TransactionChannel <- TransactionError{t, a, c}

	err = <-c
	if err == nil {
		io.WriteString(w, "true")
	} else {
		io.WriteString(w, err.Error())
	}

	return
}

func (a *ApiPort) Listen(addr string) (err error) {
	s := &http.Server{
		Addr:           addr,
		Handler:        a.d,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
