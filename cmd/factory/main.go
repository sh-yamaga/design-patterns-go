package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

func main() {
	pf := &factory.ProcessorFactory{}

	// Create a new CsvProcessor instance using the factory method
	cp := pf.CsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data

	// Create a new XmlProcessor instance using the factory method
	xp := pf.XmlProcessor()

	if err := xp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from xml file
	// Processing xml data
	// Saving processed xml data
}
