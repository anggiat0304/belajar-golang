package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	var a = 10
	var b = 20
	var c = 30

	var sum = a + b + c
	var min = a - b - c
	var multiply = a * b * c
	var divide = b / a
	var modules = c % b
	fmt.Println(sum)
	fmt.Println(min)
	fmt.Println(multiply)
	fmt.Println(divide)
	fmt.Println(modules)

	penjumlahan := helper.Penjumlahan(2, 1)
	fmt.Println(penjumlahan)
}
