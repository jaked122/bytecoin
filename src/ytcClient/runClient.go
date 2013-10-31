package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	flag.Parse()
	command := flag.Arg(0)
	switch command {
	case "address":
		resp, _ := http.Get("http://localhost:800/newWallet")
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	case "transfer":
		v := make(map[string]string)
		v["Source"] = flag.Arg(1)
		v["Destination"] = flag.Arg(2)
		v["Amount"] = flag.Arg(3)

		b, _ := json.Marshal(v)
		s := bytes.NewBuffer(b)
		resp, _ := http.Post("http://localhost:800/sendMoney", "application/json", s)
		b, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	default:
		fmt.Println(flag.Arg(0))
		fmt.Println("Command not recognized")
	}
}
