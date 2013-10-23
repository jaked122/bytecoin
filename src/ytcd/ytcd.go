package main

import (
	"libytcd"
	"log"
)

func main() {
	y := libytcd.NewYtcd()
	err := y.Listen(":800")
	if err != nil {
		log.Fatal(err)
	}
}
