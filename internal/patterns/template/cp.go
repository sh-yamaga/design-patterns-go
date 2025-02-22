package template

import "fmt"

// CSVProcessor is a concrete implementation of IDataProcessor
type CSVProcessor struct {
	BaseProcessor
}

// LoadData loads data from a CSV file
func (cp *CSVProcessor) LoadData() error {
	fmt.Println("Loading data from CSV file")
	return nil
}

// ProcessData processes the loaded data
func (cp *CSVProcessor) ProcessData() error {
	fmt.Println("Processing CSV data")
	return nil
}

// SaveData saves the processed data
func (cp *CSVProcessor) SaveData() error {
	fmt.Println("Saving processed CSV data")
	return nil
}
