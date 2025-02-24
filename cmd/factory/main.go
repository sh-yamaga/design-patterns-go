package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

func main() {
	factory := &factory.ProcessorFactory{}

	// csv processor
	cp := factory.CreateCsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data

	// xml processor
	xp := factory.CreateXmlProcessor()

	if err := xp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from xml file
	// Processing xml data
	// Saving processed xml data
}
