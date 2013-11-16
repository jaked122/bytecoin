package libytcd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"libGFC"
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
	d           *http.ServeMux
	key         map[string]*ecdsa.PrivateKey
	transaction chan TransactionError
	block       chan BlockError
}

func NewApiPort() (a *ApiPort) {
	a = new(ApiPort)

	a.key = make(map[string]*ecdsa.PrivateKey)

	a.d = http.NewServeMux()
	a.d.HandleFunc(apiroot, a.loadHomepage)
	a.d.HandleFunc(apinewWallet, a.newWallet)
	a.d.HandleFunc(apisendMoney, a.sendMoney)

	return
}

func (a *ApiPort) AddTransactionChannel(transaction chan TransactionError) {
	a.transaction = transaction
}

func (a *ApiPort) AddBlockChannel(block chan BlockError) {
	a.block = block
}

func (a *ApiPort) AddTransaction(transaction libGFC.Update) {

}

func (a *ApiPort) AddBlock(block []libGFC.Update) {

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

	a.key[host.Id] = priv

	c := make(chan error)

	a.transaction <- TransactionError{h, a, c}

	err := <-c
	if err == nil {
		io.WriteString(w, host.Id)
	} else {
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

	a.transaction <- TransactionError{t, a, c}

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
