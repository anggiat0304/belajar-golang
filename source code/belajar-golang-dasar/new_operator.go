package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func main() {
	var address *Address = new(Address)
	var address2 *Address = address

	address2.Country = "Indonesia"
	fmt.Println(address)
	fmt.Println(address2)
}
