package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fileReader()
	stringReader()
	networkReader()
}

func fileReader() {
	f, err := os.Open("./go-stuff/reader/small.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 3)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}

func stringReader() {
	s := "a small piece of string"
	r := strings.NewReader(s)

	buf := make([]byte, 3)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}

func networkReader() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 10)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}