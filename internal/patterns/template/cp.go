package template

import "fmt"

// CSVProcessor is a concrete implementation of DataProcessor
type CSVProcessor struct {
	BaseProcessor
}

// LoadData loads data from a CSV file
func (cp *CSVProcessor) LoadData() {
	fmt.Println("Loading data from CSV file")
}

// ProcessData processes the loaded data
func (cp *CSVProcessor) ProcessData() {
	fmt.Println("Processing CSV data")
}

// SaveData saves the processed data
func (cp *CSVProcessor) SaveData() {
	fmt.Println("Saving processed CSV data")
}
