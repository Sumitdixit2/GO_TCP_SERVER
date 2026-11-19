package main

import (
	"fmt"
	"log"
	"net"
)

type newconn struct {
	RemoteAddress net.Addr
	LocalAddress  net.Addr
	Received      string
	Bytes         int
}

func (c *newconn) printres() {
	fmt.Println("New Connection!")
	fmt.Println("Remote Address: ", c.RemoteAddress)
	fmt.Println("Local Address: ", c.LocalAddress)
	fmt.Println("Received: ", c.Received)
	fmt.Println("Bytes Received: ", c.Bytes)
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("not connecting sad: %v", err)
			continue
		}

		go handleconnection(conn)
	}
}

func handleconnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("error while reading from connection: ", err)
		return
	}

	message := string(buffer[:n])

	newconn := newconn{RemoteAddress: conn.RemoteAddr(), LocalAddress: conn.LocalAddr(), Received: message, Bytes: n}

	newconn.printres()

	conn.Write([]byte("hello client this is the server nd i got your message right here"))
}
