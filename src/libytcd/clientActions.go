package libytcd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostTransaction(src Account, dest Account, amount YTCVolume) {
	var t Transaction
	t.Source = src
	t.Destination = dest
	t.Amount = amount
	t.Signature[0] = 1

	// all the printfs will eventually be moved to a test function
	fmt.Printf("%v\n", t)

	b, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	} else {
		fmt.Printf("%s\n", b)
	}

	buf := bytes.NewBuffer(b)
	resp, err := http.Post("http://127.0.0.1:800/postTransaction", "application/json", buf) // not sure what the second field should be
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v\n", resp)
	}
}
