package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("error while setting up the connection: ", err)
		return
	}

	defer conn.Close()

	_, err = fmt.Fprintf(conn, "hello server this is a tcp client message/r/n/r/n")

	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}
	fmt.Println("Response status:", status)

}
