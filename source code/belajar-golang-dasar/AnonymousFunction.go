package main

import "fmt"

type Blacklist func(string) bool

func registerUser(user string, blacklist Blacklist) {
	if blacklist(user) {
		fmt.Println("You are blocked", user)
	} else {
		fmt.Println("Welcome", user)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "kontol"
	}
	registerUser("Anggiat", blacklist)

	registerUser("anjing", func(s string) bool { return s == "anjing" })

}
