package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

func main() {
	factory := &factory.ProcessorFactory{}
	processor := factory.CreateCsvProcessor()

	if err := processor.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from CSV file
	// Processing CSV data
	// Saving processed CSV data
}
