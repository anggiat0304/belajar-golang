package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])t`)
	fmt.Println(regex.MatchString("aiat"))
	fmt.Println(regex.MatchString("aathhtyrtyrtyrtyrt"))
}
