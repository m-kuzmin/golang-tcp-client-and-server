package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var (
		address = "localhost"
		port    = "8080"
	)

	ln, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server ready for connections, address:", ln.Addr())
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting a connection:", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Got a connection from:", conn.RemoteAddr())
	defer fmt.Println("Closed the connection")
	defer conn.Close()
}
