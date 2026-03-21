package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return "..."
	}
}
func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Kontol", spamFilter)
}
