package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	var (
		address = "localhost"
		port    = "8080"
	)

	log.Println("Starting client")

	conn, err := net.Dial("tcp", address+":"+port)
	defer func() {
		conn.Close()
		log.Println("Closed the connection")
	}()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to", conn.RemoteAddr())

	var reader = bufio.NewReader(os.Stdin)

	for {
		buffer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if _, err = conn.Write([]byte(buffer)); err != nil {
			log.Fatal(err)
		}
	}
}
