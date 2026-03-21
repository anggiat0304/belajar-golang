package main

import "fmt"

func main() {
	names := [...]string{"Anggiat", "Coki", "Ciko", "Lola", "Loli", "Gerry", "Yanto", "Joko", "Budi"}
	slice := names[3:6]
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := names[:]
	fmt.Println(slice2)

	slice3 := names[:3]
	fmt.Println(slice3)

	slice4 := names[2:]
	fmt.Println(slice4)

	var slice5 []string = names[:]
	fmt.Println(slice5)

	var namaBulan = [...]string{"January", "February"}
	namaBulanTambahan := append(namaBulan[:], "Maret")
	fmt.Println(namaBulanTambahan[:])

	var newSLice []string = make([]string, 2, 5)
	newSLice = append(newSLice, "Anggiat")
	newSLice = append(newSLice, "Pangaribuan")
	newSLice = append(newSLice, "Soto")
	newSLice = append(newSLice, "LLLL")
	newSLice = append(newSLice, "POPOPOPO")
	newSLice = append(newSLice, "IUWIWIWI")
	newSLice = append(newSLice, "LALALALALA")
	newSLice = append(newSLice, "Pwjss")
	newSLice = append(newSLice, "jsjjjdjdj")
	fmt.Println(newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

}
