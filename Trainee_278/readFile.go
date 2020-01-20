package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b) // n is no. of bytes, err is error
		if err == io.EOF {
			break
		}

		fmt.Printf("n: %v err: %v b:%v\n", n, err, b)
		fmt.Printf("b[:n]: %q\n", b[:n]) // from starting to end of bytes

	}
}
