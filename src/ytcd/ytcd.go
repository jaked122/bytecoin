package main

import (
	"libytcd"
	"log"
)

func main() {
	y := libytcd.NewYtcd()
	log.Print("Created")
	err := y.ListenCtl(":800")
	if err != nil {
		log.Fatal(err)
	}
}
