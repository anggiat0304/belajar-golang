package main

import (
	"fmt"
)

func endApplication() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Terjadi kesalahan", message)
}
func runApplication(err bool) {
	defer endApplication()
	if err {
		panic("error")
	}
}
func main() {
	runApplication(false)
	fmt.Println()
	// runApplication(true)
	fmt.Println("Last Bander")
}
