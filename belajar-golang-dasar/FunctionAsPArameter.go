package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Anggiat", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println("=========================")
	filter := spamFilter
	sayHelloWithFilter("Bujang", filter)
}
