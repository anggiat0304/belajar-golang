package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	// mux.Handle("/static/", fileserver)                              //Jika dengan ini akan menjadi page not found karena dia akan menjadi ./resources/static/
	mux.Handle("/static/", http.StripPrefix("/static", fileserver)) // dengan ini maka prefix static akan dilewatkan atau dihapus menjadi ./resources

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	// fileserver := http.FileServer(http.FS(resources)) // jika dengan hanya ini hasilnya akan page not found karena nama folder menjadi nama resouces atau url
	directory, _ := fs.Sub(resources, "resources")
	fileserver := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	// mux.Handle("/static/", fileserver)                              //Jika dengan ini akan menjadi page not found karena dia akan menjadi /static/
	mux.Handle("/static/", http.StripPrefix("/static", fileserver)) // dengan ini maka prefix static akan dilewatkan atau dihapus menjadi ./ akan tetapi di browser tetap menggunakan static/

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
