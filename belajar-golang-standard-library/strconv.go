package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Boolean:", result)
	}

	intResult, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Integer:", intResult)
	}

	floatResult, err := strconv.ParseFloat("3.14159", 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Float:", floatResult)
	}

	str := strconv.Itoa(67890)
	fmt.Println("Integer to String:", str)

	boolStr := strconv.FormatBool(false)
	fmt.Println("Boolean to String:", boolStr)

	floatStr := strconv.FormatFloat(2.71828, 'f', 2, 64)
	fmt.Println("Float to String:", floatStr)

	binary := strconv.FormatInt(255, 2)
	fmt.Println("Integer to Binary String:", binary)
}
