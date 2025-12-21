package main

import "fmt"

type HasName interface {
	GetNName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetNName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetNName() string {
	return animal.Name
}
func (person Person) GetNName() string {
	return person.Name
}
func main() {
	person := Person{Name: "Anggiat"}
	animal := Animal{Name: "BBB"}
	SayHello(person)
	SayHello(animal)

}
