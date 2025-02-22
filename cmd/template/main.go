package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template"
)

func main() {
	cp := template.NewCsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from CSV file
	// Processing CSV data
	// Saving processed CSV data
}
