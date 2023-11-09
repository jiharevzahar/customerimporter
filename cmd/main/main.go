package main

import (
	"fmt"
	"github.com/jiharevzahar/customerimporter/importer"
	"log"
	"os"
)

func main() {
	reader := importer.NewCustomerImporter(os.Getenv("CSV_PATH"))
	records, err := reader.Import()
	if err != nil {
		log.Fatal("read failed: ", err)
	}

	fmt.Println(records)
	fmt.Println("success")
}
