package main

import (
	"os"
	"net"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func handleConnection(conn net.Conn, db *sql.DB) {
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
			break
		}

		_, err = outputFile.Write(buf[:bytesRead])
		if err != nil {
			fmt.Printf("Had problems writing the file!\n")
		}
	}

	// We need to hash the file before putting it in the db
	_, err = db.Exec("INSERT INTO hostsAndFiles(file, host) VALUES(25, 45)")
	if err != nil {
		fmt.Printf("Database query failed, most likely the file is already in the system\n")
	} else {
		fmt.Printf("Hopefully the db quere executed\n");
	}
}

func main() {
	// Try to get server online
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Server failed to open\n")
		os.Exit(1)
	}

	// Try to open the database
	db, err := sql.Open("sqlite3", "networkState.db")
	if err != nil {
		fmt.Printf("SQL Database failed to open!\n")
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS hostsAndFiles (file INT NOT NULL PRIMARY KEY, host INT)")
	if err != nil {
		fmt.Printf("Database Query Failed!\n")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Client connection failed\n")
			continue
		}
		
		// Should this be restful?
		go handleConnection(conn, db)
	}
}
