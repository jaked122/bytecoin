package libytcd

import (
    "net/http"
    "io"
    "time"
)

type ytcServer struct {
    state map[Account]YTC
    transactions [] Transaction
}

func NewYtcd() (y * ytcServer) {
    y = new(ytcServer)
    return
}

func (y * ytcServer) AddTransaction(t Transaction) {
    y.transactions = append(y.transactions, t)
    y.state[t.Source] -= t.Amount
    y.state[t.Destination] += t.Amount
}

func (y * ytcServer) Listen(addr string) (err error) {

    d := http.NewServeMux()
    d.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
        io.WriteString(w, "hello world")
    })

    s := &http.Server {
        Addr: addr,
        Handler: d,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    return s.ListenAndServe()
}

