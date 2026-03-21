package main

import (
	"fmt"
	"time"
)

func main() {
	var duration time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 10
	var duration3 time.Duration = duration - duration2

	fmt.Println(duration3)
	fmt.Printf("duration : %d\n", duration3)
}
