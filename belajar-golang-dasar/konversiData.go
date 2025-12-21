package main

import "fmt"

func main() {
	var nilai32 int32 = 7879
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Anggiat Pangaribuan"
	var a = name[0]
	var aString = string(a)

	fmt.Println(aString)
}
