package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func main() {
	address1 := Address{"Jakarta Barat", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Medan"
	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = &address1
	fmt.Println(address3)

	address3.City = "Padang"
	fmt.Println(address1)
}
