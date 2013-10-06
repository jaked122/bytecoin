package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Printf("Failed to connect to the server!\n")
		os.Exit(0);
	}

	fmt.Fprintf(conn, "filnononononono\n")
}
