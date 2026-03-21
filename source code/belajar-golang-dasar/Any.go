package main

import "fmt"

func ups() any {
	return "ups"
}
func oops() interface {
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
	var coba = oops()
	fmt.Println(coba)
}
