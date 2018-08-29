package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader interface {
	Read(p []byte) (n int, err error)
}

type logWriter struct{}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
	}

	lw := logWriter{}

	io.Copy(lw, file)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many of bytes:", len(bs))
	return len(bs), nil
}
