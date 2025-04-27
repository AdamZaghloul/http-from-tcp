package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Connection accepted.")

		ch := getLinesChannel(conn)

		for line := range ch {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection closed.")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {

	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)
		currentLineContents := ""
		for {
			b := make([]byte, 8, 8)
			n, err := f.Read(b)
			if err != nil {
				if currentLineContents != "" {
					ch <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				ch <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()

	return ch
}
