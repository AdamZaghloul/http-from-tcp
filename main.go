package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	offset := 0
	b := make([]byte, 8)
	line := ""

	for err == nil {
		_, err = file.ReadAt(b, int64(offset))
		parts := strings.Split(string(b), "\n")
		line += parts[0]
		if len(parts) > 1 {
			fmt.Printf("read: %s\n", line)
			line = parts[1]
		}
		offset += 8
	}
}
