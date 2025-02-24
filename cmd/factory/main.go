package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/factory"
)

func main() {
	factory := &factory.ProcessorFactory{}
	cp := factory.CreateCsvProcessor()

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data
}
