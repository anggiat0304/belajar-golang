package main

import "fmt"

func main() {
	var name = "anggiat"
	var name2 = "anggiat"

	var result bool = name == name2
	fmt.Println(result)

	result = name <= name2
	fmt.Println(result)

	result = name >= name2
	fmt.Println(result)

	result = name < name2
	fmt.Println(result)

	result = name > name2
	fmt.Println(result)

	result = name != name2
	fmt.Println(result)

}
