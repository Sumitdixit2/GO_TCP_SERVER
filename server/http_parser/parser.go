package httpparser

import (
	"bytes"
	"fmt"
)

type HTTPHeaders struct {
	Host          string
	UserAgent     string
	Accept        string
	ContentType   string
	ContentLength int
}

type client struct {
	Method  string
	Path    string
	Headers HTTPHeaders
	Body    []byte
}

func Parser(data []byte) {
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	message := string(data[:headerEnd])

	fmt.Println("message is", message)
}
