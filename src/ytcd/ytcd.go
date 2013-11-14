package main

import (
	"libytcd"
	"log"
	"time"
)

func main() {
	api := libytcd.NewApiPort()
	go func() {
		err := api.Listen(":800")
		if err != nil {
			log.Fatal(err)
		}
	}()
	b := make([]libytcd.Port, 1)
	b[0] = api
	_ = libytcd.NewServer(b)
	for {
		time.Sleep(1 * time.Hour)
	}
}
