package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/eksternal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
