package libytcd

import (
	//"net/http"
)

func PostTransaction(src Account, dest Account, amount YTCVolume) {
	var t Transaction
	t.Source = src
	t.Destination = dest
	t.Amount = amount
	t.Signature = "true"

	//We can do this if we write a 'Read' function into Transaction
	//http.Post("127.0.0.1:800/postTransaction", "ytc/transaction", t)
}
