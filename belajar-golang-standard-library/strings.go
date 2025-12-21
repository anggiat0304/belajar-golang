package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Anggiat Pangaribuan", "Anggiat"))
	fmt.Println(strings.Split("Anggiat Pangaribuan", " "))
	fmt.Println(strings.ToLower("Anggiat Pangaribuan"))
	fmt.Println(strings.ToUpper("Anggiat Pangaribuan"))
	fmt.Println(strings.Trim("    ANggiat Pangaribuan   ", " "))
	fmt.Println(strings.ReplaceAll("anggiat anggiat anggiat", "anggiat", "ganteng"))
}
