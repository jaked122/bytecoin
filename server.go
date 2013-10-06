package main

import (
	"os"
	"net"
	"fmt"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("Incoming Connection!\n")

	outputFile, err := os.Create("storage.ytc")
	if err != nil {
		fmt.Printf("Error creating storage.ytc")
		return
	}
	defer outputFile.Close()

	// For now, assume a file transaction
	for {
		buf := make([]byte, 4096)
		bytesRead, err := conn.Read(buf)
		if err != nil {
			return
		}

		n, err := outputFile.Write(buf[:bytesRead])
		fmt.Printf("Bytes Written: %u\n", n);
		if err != nil {
			fmt.Printf("Had problems writing the file!\n")
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Server failed to open\n")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Client connection failed\n")
			continue
		}
		
		// Should this be restful?
		go handleConnection(conn)
	}
}
