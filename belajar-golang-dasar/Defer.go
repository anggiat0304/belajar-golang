package main

import "fmt"

func logging() {
	fmt.Println("Fungsi sudah selesai dijankan")
}
func rubApplication() {
	defer logging()
	fmt.Println("Fungsi dijankan")
}
func main() {
	rubApplication()
}
