package main

import (
	"os"
	"net"
	"fmt"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("Incoming Connection!\n")

	// For now, assume a file transaction
	for {
		buf := make([]byte, 4096)
		bytesRead, err := conn.Read(buf)
		if err != nil {
			return
		}

		fileData := buf[0:bytesRead]
		outputFile, err := os.Create("storage.ytc")
		if err != nil {
			fmt.Printf("Error creating storage.ytc")
			return
		}
		defer outputFile.Close()

		_, err = outputFile.Write(buf[0:bytesRead])
		if err != nil {
			fmt.Printf("Had problems writing the file!\n")
		}
				
		fmt.Printf("%s\n", string(fileData))
		fmt.Printf("File command recognized\n")
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
