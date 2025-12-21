package main

import "fmt"

func factorialLoop(value int) int64 {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return int64(result)
}

func main() {
	fmt.Println(factorialLoop(10))
}
