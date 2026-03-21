package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, v := range number {
		total = total + v
	}
	return total
}
func main() {
	total := sumAll(10, 20, 20, 20, 30, 200000000)
	fmt.Println(total)

	number := []int{10, 30, 30, 30}
	fmt.Println("Konversi slice to variadic fucntion", sumAll(number...))
}
