package main

import "fmt"

type Address struct {
	City, Provincy, Country string
}

func ChangeCountry(address *Address, to string) {
	address.Country = to
}
func main() {
	alamat := Address{"Subang", "Jawa Barat", ""}
	fmt.Println(alamat)
	ChangeCountry(&alamat, "Indonesia")
	fmt.Println(alamat)
}
