package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	input := strings.NewReader("This is long string\nso take your time baby\n")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
