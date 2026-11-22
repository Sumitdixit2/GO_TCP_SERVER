package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type httpheader struct {
	Host       string
	User_Agent string
	Accept     string
}

type reqheaders struct {
	Content_Type   string
	Content_Length int
}

type httpreq struct {
	Method          string
	Path            string
	HttpHeaders     httpheader
	ResponseHeaders reqheaders
	data            json.Delim
}

func (r *httpreq) sendreq(conn net.Conn) (int, error) {
	message := r.Method + "\r\n" + r.Path + "\r\n" + string(r.HttpHeaders.Host) + string(r.HttpHeaders.User_Agent) + string(r.HttpHeaders.Accept) + "\r\n" + string(r.ResponseHeaders.Content_Length) + string(r.ResponseHeaders.Content_Type) + "\r\n\r\n" + string(r.data)
	n, err := fmt.Fprintf(conn, message)

	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	buffer := make([]byte, 1024)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("error while setting up the connection: ", err)
		return
	}

	defer conn.Close()

	req := httpreq{
		Method:          "GET",
		Path:            "/user",
		HttpHeaders:     httpheader{Host: "example.com", User_Agent: "mozilla", Accept: "*/*"},
		ResponseHeaders: reqheaders{Content_Type: "application/json", Content_Length: 45},
	}

	_, err = req.sendreq(conn)

	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	n, err := conn.Read(buffer)

	message := string(buffer[:n])

	fmt.Println("message from the server is: ", message)

	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

}
