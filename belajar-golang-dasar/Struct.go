package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello,", customer.Name)
}
func main() {
	var anggiat Customer
	anggiat.Name = "Anggiat"
	anggiat.Address = "Jakarta"
	anggiat.Age = 12
	fmt.Println(anggiat)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     90,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Bogor", 12}
	budi.sayHello()
	fmt.Println(budi)
}
