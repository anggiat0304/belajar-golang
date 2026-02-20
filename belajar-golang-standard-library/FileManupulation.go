package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// createNewFile("Anggiat.txt", "Hai Aku anggiat")
	// result, _ := readFile("Anggiat.txt")
	_ = addToFile("Anggiat.txt", "\nTambahan abangku")
	// fmt.Println(result)
}
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}
func readFile(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}
