package main

import (
	"errors"
	"fmt"
)

func pembagian(dibagi int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh NOL ")
	} else {
		return dibagi / pembagi, nil
	}
}
func main() {
	fmt.Println(pembagian(4, 23))
	fmt.Println(pembagian(1, 2))
	fmt.Println(pembagian(6, 2))
	fmt.Println(pembagian(4, 0))
	hasil, err := pembagian(3, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
