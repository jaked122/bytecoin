package libytcd

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetAddress(addr string) string {
	resp, err := http.Get("http://" + addr + "/newWallet")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}
