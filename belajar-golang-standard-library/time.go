package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2025, time.December, 20, 21, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse(time.RFC3339, "2025-12-20T21:20:21Z")
	fmt.Println(parse)

	formatter := "2006-01-02 15:04:05"
	value := "2025-12-20 21:20:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Hour())
}
