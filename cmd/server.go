package main

import (
	"log"
	"net"
)

const client_msg_buffer = 1024

func main() {
	var (
		address = "localhost"
		port    = "8080"
	)

	ln, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server ready for connections, address:", ln.Addr())
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting a connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	address := conn.RemoteAddr()

	log.Println("Got a connection from:", conn.RemoteAddr())
	defer func() {
		conn.Close()
		log.Println("Closed connection to", address)
	}()

	buffer := make([]byte, client_msg_buffer)

	for {
		len, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from", address, "Reason:", err)
			return
		}
		msg := buffer[:len]
		log.Println("Message from", address, len, "bytes:", string(msg))
	}
}
