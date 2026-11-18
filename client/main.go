package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	buffer := make([]byte, 1024)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("error while setting up the connection: ", err)
		return
	}

	defer conn.Close()

	_, err = fmt.Fprintf(conn, "hello server this is a tcp client message")

	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	fmt.Println("connection is: ", conn)

	n, err := conn.Read(buffer)

	message := string(buffer[:n])

	fmt.Println("message from the server is: ", message)

	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

}
