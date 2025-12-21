package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	fmt.Println(address1)
	address2.Country = "Amerika"
	fmt.Println(address1)

	// address2 = &Address{"Medan", "Sumatra Utara", "Indonesia"}
	*address2 = Address{"Medan", "Sumatra Utara", "Indonesia"}
	fmt.Println(address2)
	fmt.Println(address1)
}
