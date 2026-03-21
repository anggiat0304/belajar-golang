package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}
func runApp(error bool) {
	defer endApp()
	fmt.Println("Aplikasi mulai berjalan......")
	if error {
		panic("Aplikasi mengalami kerusakan")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	fmt.Println("")
	runApp(true)
}
