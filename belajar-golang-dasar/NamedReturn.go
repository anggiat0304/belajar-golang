package main

import "fmt"

func getCompleteName() (firstName, middlename, lastname string) {
	firstName = "anggiat"
	middlename = "marajahan"
	lastname = "pangaribuan"
	return firstName, middlename, lastname
}
func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
