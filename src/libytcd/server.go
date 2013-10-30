package libytcd

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
	s         *State
	d         *http.ServeMux
	neighbors []*json.Encoder
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

	y.Send(h)
}

func (y *ytcServer) Send(h Update) {
	for _, neighbor := range y.neighbors {
		err := neighbor.Encode(h)
		if err != nil {
			log.Print(err)
		}
	}
}

func (y *ytcServer) sendMoney(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	v := make(map[string]string)
	err := json.Unmarshal(b, &v)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	source := v["Source"]
	destination := v["Destination"]
	amount, _ := strconv.ParseUint(v["Amount"], 10, 64)

	t := NewTransferUpdate()
	t.Source = HostKey(source)
	t.Destination = HostKey(destination)
	t.Amount = YTCAmount(amount)
	t.Signature = Signature("fu")

	err = t.Verify(y.s)
	if err != nil {
		log.Print(y.s)
		io.WriteString(w, err.Error())
		return
	}
	t.Apply(y.s)
	io.WriteString(w, "true")

	y.Send(t)

	return
}

func (y *ytcServer) HandleNetworkConnection(c net.Conn) {
	je := json.NewEncoder(c)
	y.neighbors = append(y.neighbors, je)

	j := json.NewDecoder(c)
	for {
		v := make(map[string]interface{})
		j.Decode(&v)

		b, err := json.Marshal(v)
		if err != nil {
			log.Print(err)
			break
		}

		update, err := ParseUpdate(b)
		if err != nil {
			log.Print(err)
			break
		}

		err = update.Verify(y.s)
		if err != nil {
			log.Print(err)
		}

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

func (y *ytcServer) ConnectAddress(addr string) (err error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	go y.HandleNetworkConnection(c)
	return
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
