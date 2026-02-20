package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Jhon", "Anggiat", "Paul", "Bapakmu"}
	values := []int{100, 20, 30, 10}
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Anggiat"))
	fmt.Println(slices.Index(names, "Bapakmu"))
}
