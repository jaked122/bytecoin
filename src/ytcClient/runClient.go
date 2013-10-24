package main

import (
	"libytcd"
)

func main() {
	srcAccnt := libytcd.Account("s")
	destAccnt := libytcd.Account("d")

	libytcd.PostTransaction(srcAccnt, destAccnt, 1234)
}
