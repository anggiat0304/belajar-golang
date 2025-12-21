package main

import "fmt"

func sayHello(name string) string {
	return "hello " + name
}
func main() {
	greeting := sayHello
	fmt.Println(greeting("Anggiat"))
}
