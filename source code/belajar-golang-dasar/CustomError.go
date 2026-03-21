package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func cariOrang(id int, name string) (string, error) {
	if id == 0 {
		return "", &validationError{Message: "Id tidak bisa 0 bang"}
	} else if name == "bajingan" {
		return "", &notFoundError{Message: "Yang benar dikit lah kau bajiangan mana ada nama bajingan disini pukimak"}
	}
	return name + " ditemukan", nil
}

func main() {
	people, err := cariOrang(2, "bajingan")
	if err == nil {
		fmt.Println(people)
	} else {
		fmt.Println(err)
	}
}
