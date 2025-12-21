package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	name string
}

type Person2 struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"100"`
	Email   string `required:"true" max:"50" `
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		fmt.Println("Type Name ", valueType.Name())
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "With type", structField.Type)
		fmt.Println("isRequired?", structField.Tag.Get("required"))
		fmt.Println("maximum: ", structField.Tag.Get("max"))
	}
}
func isValid(data interface{}) {
	fmt.Println("Check isValid")
	t := reflect.TypeOf(data)
	fmt.Println("NumField: ", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("Field Name :", field.Name)
		if field.Tag.Get("required") == "true" {
			isValid := reflect.ValueOf(data).Field(i).Interface() != ""
			fmt.Println(isValid)
		}
	}
}
func main() {
	readField(Sample{"Anggiat"})
	readField(Person2{"Coki", "Jakarta Barat", "anggiatpangaribuan12@gmail.com"})
	person := Person2{
		Name:    "Anggiat Pangaribuan",
		Address: "Sitoluana",
	}
	isValid(person)
}
