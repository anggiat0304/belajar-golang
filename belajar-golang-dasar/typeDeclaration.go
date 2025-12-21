package main

import "fmt"

func main() {
	type NoKTP string
	var ktpNumber NoKTP = "7987897987978"
	fmt.Println(ktpNumber)
	fmt.Println(NoKTP("98080808"))
}
