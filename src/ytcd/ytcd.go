package main

import (
	"libytcd"
)

func main() {
	api := libytcd.NewApiPort()
	api.Listen(":800")
	b := make([]libytcd.Port, 1)
	b[1] = api
	_ = libytcd.NewServer(b)
}
