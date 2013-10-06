package main

import (
	"net"
	"os"
	"fmt"
)

func storeFile(filename string) {
	inputFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: could not open file\n")
		os.Exit(1)
	}
	defer inputFile.Close()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Failed to connect to the server!\n")
		os.Exit(0);
	}

	buf := make([]byte, 4096)
	for {
		bytesRead, err := inputFile.Read(buf)
		fmt.Printf("Printing some bytes: %u\n", bytesRead)
		if err != nil {
			return
		}
		
		conn.Write(buf[:bytesRead])
		//fmt.Fprintf(conn, "%x", buf[:bytesRead])
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("must include term 'store'\n")
		os.Exit(1)
	}

	if args[1] == "store" {
		if len(args) < 3 {
			fmt.Printf("Must include a file to store!\n")
			os.Exit(1)
		}

		storeFile(args[2])
	}
}
