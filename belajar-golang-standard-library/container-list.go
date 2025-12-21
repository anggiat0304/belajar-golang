package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Angggiat")
	data.PushBack("Raju")
	data.PushBack("Bapak")
	// data.PushFront("Anjing")

	var head *list.Element = data.Front()
	head.Next().Next().Prev().Value = "Ganteng"
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}
}
