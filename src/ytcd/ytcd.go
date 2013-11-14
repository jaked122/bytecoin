package main

import (
	"libytcd"
)

func main() {
	api := libytcd.NewApiPort()
	api.Listen(":800")
}
