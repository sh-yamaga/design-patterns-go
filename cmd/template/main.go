package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/template"
)

func main() {
	// csv processor
	cp := template.NewCsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err)
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data

	// xml processor
	xp := template.NewXmlProcessor()

	if err := xp.Execute(); err != nil {
		fmt.Println(err)
	}
	// Output:
	// Loading data from xml file
	// Processing xml data
	// Saving processed xml data
}
