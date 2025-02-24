package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/template/processor"
)

func main() {
	cp := &processor.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	if err := cp.Execute(); err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Loading data from csv file
	// Processing csv data
	// Saving processed csv data
}
