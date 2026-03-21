package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Anggiat"
	names[1] = "Coki"
	names[2] = "young lex"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		10,
		92,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))

	values[2] = 1
	fmt.Println(values)

	var values2 = [...]int{
		90,
		10,
		92,
		98,
		87,
		89,
	}
	fmt.Println(values2)
	fmt.Println(values2[0])
	fmt.Println(values2[1])
	fmt.Println(values2[2])
	fmt.Println(len(values2))

}
