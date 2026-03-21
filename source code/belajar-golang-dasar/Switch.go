package main

import (
	"fmt"
)

func main() {
	name := "Anggiat"
	length := len(name)
	switch name {
	case "LALA":
		fmt.Println("Hello LALA")
	case "Anggiat":
		fmt.Println("Hello Anggiat")
	default:
		fmt.Println("Hi , bang")
	}

	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Kepanjangann bang nnamanya")
	case false:
		fmt.Println("Nama oke lah")
	}

	switch {
	case name == "Anggiat":
		fmt.Println("Anggiat")
	case length <= 100:
		fmt.Printf("Aman")
	}

}
