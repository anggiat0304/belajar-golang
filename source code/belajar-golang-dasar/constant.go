package main

import "fmt"

func main() {
	const firstname string = "anggiat"
	const lastname = "pangaribuan"
	fmt.Println(firstname)
	fmt.Println(lastname)

	const (
		kelas string = "satu"
		umur         = 18
	)
	fmt.Println(kelas)
	fmt.Println(umur)
}
