// Package main contains the server implementation
package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

const clientMsgBuffer = 1024

func main() {
	const (
		address = "localhost"
		port    = "8080"
	)

	ln, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server ready for connections, address:", ln.Addr())
	defer ln.Close()

	var (
		interruptCh = make(chan os.Signal, 1)
		quitCh      = make(chan struct{})
	)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	go func() {
		<-interruptCh
		close(quitCh)
		ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error, stopping listener:", err)
			break
		}

		wg.Add(1)
		go handleConnection(conn, &wg, quitCh)
	}

	wg.Wait()
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup, quitCh <-chan struct{}) {
	address := conn.RemoteAddr()
	defer func() {
		conn.Close()
		log.Println("Closed connection to", address)
		wg.Done()
	}()

	log.Println("Got a connection from:", conn.RemoteAddr())

	buffer := make([]byte, clientMsgBuffer)

	go func() {
		for {
			_, ok := <-quitCh
			if !ok {
				log.Println("Closing connection to", address)
				conn.Close()
				break
			}
		}
	}()

	for {
		// First read the client's message
		length, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from", address, "Reason:", err)
			return
		}
		msg := buffer[:length]
		log.Println("Message from", address, length, "bytes:", string(msg))

		// After the client sent the message respond with some data
		response := []byte("Received " + strconv.Itoa(length) + " bytes!")

		if _, err := conn.Write(response); err != nil {
			log.Println("Error sending reponse to", address, "Reason", err)
		} else {
			log.Println("Sent responce to the client")
		}
	}
}
