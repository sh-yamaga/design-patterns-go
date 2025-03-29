package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

func main() {
	// csv processor
	cp := processor.NewCsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data

	// xml processor
	xp := processor.NewXmlProcessor()

	if err := xp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from xml file
	// Processing xml data
	// Saving processed xml data
}
