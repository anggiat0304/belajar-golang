package main

import "fmt"

func main() {
	counter := 1
	fmt.Println("For statement terpisah")
	for counter <= 10 {
		fmt.Println("Counter", counter)
		counter++
	}
	fmt.Println("\nFor Statement")
	for i := 0; i <= 10; i++ {
		fmt.Println("i", i)
	}
	fmt.Println("\n For mannual")
	names := []string{"anggiat", "coki", "ciko"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Nama ke:", i, names[i])
	}
	fmt.Println("\n For range otomatis menggunakan range")
	for i, name := range names {
		fmt.Println("name ke", i, "=", name)
	}
	for _, name := range names {
		fmt.Println("name", name)
	}
}
