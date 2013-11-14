package libytcd

import (
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
	d *http.ServeMux
}

func NewApiPort() (y *ApiPort) {
	y = new(ApiPort)

	y.d = http.NewServeMux()
	y.d.HandleFunc(apiroot, y.loadHomepage)
	y.d.HandleFunc(apinewWallet, y.newWallet)
	y.d.HandleFunc(apisendMoney, y.sendMoney)

	return
}

// Will eventually load html explaining how to get started?
func (y *ApiPort) loadHomepage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func (y *ApiPort) newWallet(w http.ResponseWriter, r *http.Request) {

	b := make([]byte, 8)
	rand.Read(b)

	priv, host := libGFC.NewHost("foo")
	h := libGFC.NewHostUpdate(host)
	h.Sign(priv)
}

func (y *ApiPort) sendMoney(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	v := make(map[string]string)
	err := json.Unmarshal(b, &v)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	source := v["Source"]
	destination := v["Destination"]
	amount, _ := strconv.ParseUint(v["Amount"], 10, 64)

	_ = libGFC.NewTransferUpdate(source, destination, amount)

	io.WriteString(w, "true")

	return
}

func (y *ApiPort) Listen(addr string) (err error) {
	s := &http.Server{
		Addr:           addr,
		Handler:        y.d,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
}
