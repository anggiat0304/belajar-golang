package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name  string
	Email string
	Age   int
}
type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	users := []User{
		{"Anggiat", "anggiatpangaribuan12@gmail.com", 24},
		{"Rojak", "rojak@gmail.com", 25},
		{"Royahdi", "rohyadi@gmail.com", 25},
		{"Jajajaj", "jajaaja@gmail.com", 19},
		{"Raja", "raja@gmail.com", 17},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
