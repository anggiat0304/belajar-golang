package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "Anggiat,Coki,Ganteng\n" +
		"Anjing,Kucing,Babi\n" +
		"Kaya,Miskin, Apaaja"

	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

}
