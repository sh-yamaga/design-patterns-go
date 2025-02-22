package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/template"

func main() {
	var cp *template.CSVProcessor = &template.CSVProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	cp.Execute()
	// Output:
	// Loading data from CSV file
	// Processing CSV data
	// Saving processed CSV data
}
