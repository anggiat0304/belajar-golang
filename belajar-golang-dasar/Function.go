package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hi", name)
}
func penjumlahan(a int, b int) int {
	return a + b
}
func pengaliandanpenjumlahann(a int, b int) (int, int) {
	return a * b, a + b
}
func main() {
	fmt.Println("Function tanpa return")
	sayHello("Anggiat")
	sayHello("HArto")

	fmt.Println("\nFunction dengan return")
	fmt.Println("Function penjumlahan", penjumlahan(1, 2))

	fmt.Println("\nFuntion multiple return")
	pengalian, penjumlahan := pengaliandanpenjumlahann(1, 3)
	fmt.Println("pengalian", pengalian, "penjumlahan", penjumlahan)
}
