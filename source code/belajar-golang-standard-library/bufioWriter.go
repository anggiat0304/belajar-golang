package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Anggiat Pangaribuan\n")
	_, _ = writer.WriteString("Alaksa Simanjuntak\n")
	_, _ = writer.WriteString("Apapabik\n")
	writer.Flush()

}
