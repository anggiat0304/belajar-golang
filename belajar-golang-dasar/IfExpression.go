package main

import "fmt"

func main() {
	name := "anggiat"
	if name == "anggiat" {
		fmt.Println("anggiat")
	} else if name == "coki" {
		fmt.Println("coki")
	} else {
		fmt.Println("false")
	}

	if length := len(name); length > 2 {
		fmt.Println("lebih besar dari dua")
	} else {
		fmt.Println(len(name))
	}
}
