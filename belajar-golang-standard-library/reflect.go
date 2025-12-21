package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	name string
}

type Person struct {
	Name string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name ", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "With type", valueField.Type)
	}
}
func main() {
	readField(Sample{"Anggiat"})
	readField(Person{"Coki"})
}
