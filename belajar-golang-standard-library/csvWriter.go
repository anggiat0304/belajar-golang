package main

import (
	"encoding/csv"
	"os"
)

func main() {

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Anggiat", "Pangaribuan", "Sitampan"})
	_ = writer.Write([]string{"Siapapun", "Apaaja", "lalal"})
	_ = writer.Write([]string{"pukimak", "bapakmu", "dklsfjelsj"})
	writer.Flush()
}
