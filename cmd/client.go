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

	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to", conn.RemoteAddr())

	defer fmt.Println("Closed the connection")
	defer conn.Close()
}
