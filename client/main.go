package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type HTTPHeaders struct {
	Host          string
	UserAgent     string
	Accept        string
	ContentType   string
	ContentLength int
}

type HTTPRequest struct {
	Method  string
	Path    string
	Headers HTTPHeaders
	Body    []byte
}

func (r *HTTPRequest) SendReq(conn net.Conn) (int, error) {
	r.Headers.ContentLength = len(r.Body)

	message := fmt.Sprintf(
		"%s %s HTTP/1.1\r\n"+
			"Host: %s\r\n"+
			"User-Agent: %s\r\n"+
			"Accept: %s\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n",
		r.Method,
		r.Path,
		r.Headers.Host,
		r.Headers.UserAgent,
		r.Headers.Accept,
		r.Headers.ContentType,
		r.Headers.ContentLength,
	)

	n, err := io.WriteString(conn, message)
	if err != nil {
		return n, err
	}

	bodyN, err := conn.Write(r.Body)
	return n + bodyN, err
}

func main() {
	buffer := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("error while setting up the connection: ", err)
		return
	}

	data := map[string]any{
		"name": "Sumit",
		"age":  18,
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	req := HTTPRequest{
		Method: "POST",
		Path:   "/users",
		Headers: HTTPHeaders{
			Host:        "localhost:8080",
			UserAgent:   "MyGoClient/1.0",
			Accept:      "application/json",
			ContentType: "application/json",
		},
		Body: body,
	}

	_, err = req.SendReq(conn)

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
