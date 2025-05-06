package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Could not resolve %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Could not dial %s", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Could not parse input %s", err.Error())
			continue
		}

		_, err = conn.Write([]byte(str))
		if err != nil {
			log.Fatal("Could not write input %s", err.Error())
			continue
		}
	}
}
