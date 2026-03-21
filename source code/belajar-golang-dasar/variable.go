package main

import "fmt"

func main() {
	var nama string
	nama = "Anggiat Pangaribuan"
	fmt.Println(nama)
	nama = "Anggiat Pasaribu"
	fmt.Println(nama)

	// deklarasi tanpa var
	name := "Anggiat Pangaribuan"
	fmt.Println(name)
	name = "Anggiat Pasaribu"
	fmt.Println(name)

	// deklarasi bannyak variable
	var (
		firstname = "Anggiat"
		lastname  = "Pangaribuan"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}
