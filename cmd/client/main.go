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
		// Read the user input
		buffer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Send the input
		if _, err = conn.Write([]byte(buffer)); err != nil {
			log.Fatal(err)
		}

		// Read the responce and print it
		response := make([]byte, 1024)
		length, err := conn.Read(response)
		if err == nil {
			response = response[:length]
			log.Println("Responce:", string(response))
		} else {
			log.Println("Error reading responce:", err)
		}
	}
}
