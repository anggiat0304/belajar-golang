package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var encode = base64.StdEncoding.EncodeToString([]byte("Anggiat Pangaribuan"))
	fmt.Println(encode)
	var decode, err = base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
