package template

import "fmt"

// CsvProcessor is a concrete implementation of IDataProcessor
type CsvProcessor struct {
	BaseProcessor
}

// Load loads data from a CSV file
func (cp *CsvProcessor) Load() error {
	fmt.Println("Loading data from csv file")
	return nil
}

// ProcessData processes the loaded data
func (cp *CsvProcessor) Process() error {
	fmt.Println("Processing csv data")
	return nil
}

// SaveData saves the processed data
func (cp *CsvProcessor) Save() error {
	fmt.Println("Saving processed csv data")
	return nil
}
