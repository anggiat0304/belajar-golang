package main

import "fmt"

func main() {
	person := map[string]string{
		"Nama":  "Anggiat",
		"Marga": "Pangaribuan",
	}
	fmt.Println(person)
	fmt.Println(person["Nama"])
	fmt.Println(person["Marga"])
	fmt.Println(person["Alamat"])

	book := make(map[string]string)
	book["Title"] = "Malin Kudang"
	book["Tahun"] = "2005"
	book["Opps"] = "Opps"
	fmt.Println(book)
	fmt.Println(book["Title"])
	fmt.Println(book["Tahun"])

	delete(book, "Opps")
	fmt.Println(book)

}
