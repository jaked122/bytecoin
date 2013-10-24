package main

import (
	"libytcd"
)

func main() {
	var srcAccnt [64]byte
	var destAccnt [64]byte

	srcAccnt[0] = 's'
	destAccnt[0] = 'd'
	libytcd.PostTransaction(srcAccnt, destAccnt, 1234)
}
