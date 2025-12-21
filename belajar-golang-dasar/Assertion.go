package main

import (
	"fmt"
)

func random() any {
	return "OK"
}
func handlingException() {
	message := recover()
	fmt.Println("Terjadi kesalahan", message)
}
func main() {
	defer handlingException()
	var result any = random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
	// var resultInt = result.(int)
	// message := recover()
	// fmt.Println("terjadi kesalahan", message)
	// fmt.Println(resultInt)

}
