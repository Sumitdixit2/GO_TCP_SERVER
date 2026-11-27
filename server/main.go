package main

import (
	"fmt"
	"log"
	"net"
	httpparser "tcp-server/server/http_parser"
)

// goal is to umm... ig like take our res from client , parse the binary data into proper format
// and then likee take the response then create an api which takes the req and then if the req is cool then we send a response
//
//what i need is the clientreq type and the client res type like what to send to the client...

type res struct {
	Version    string
	StatusCode int
	Headers    httpparser.HTTPHeaders
	Body       []byte
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

	fmt.Println("data from client is: ", message)

	httpparser.Parser(buffer[:n])

	conn.Write([]byte("hello client this is the server nd i got your message right here"))
}
