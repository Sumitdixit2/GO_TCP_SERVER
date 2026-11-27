package httpparser

import (
	"bytes"
	"fmt"
)

func Parser(data []byte) {
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	message := string(headerEnd)

	fmt.Println("message is", message)
}
