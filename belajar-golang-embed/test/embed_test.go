package belajargolangembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed lee_jeehon.jpeg
var image []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func TestString(t *testing.T) {
	fmt.Println(version)
}

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("leejehon.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ", string(content))

		}
	}
}
