package processor

import "fmt"

// CsvProcessor is a concrete implementation of IDataProcessor
type CsvProcessor struct {
	BaseProcessor
}

// NewCsvProcessor returns a new instance of CsvProcessor
func NewCsvProcessor() *CsvProcessor {
	cp := &CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}

// Load loads data from a CSV file
func (cp *CsvProcessor) Load() error {
	fmt.Println("Loading data from csv file")
	return nil
}

// Process processes the loaded data
func (cp *CsvProcessor) Process() error {
	fmt.Println("Processing csv data")
	return nil
}

// Save saves the processed data
func (cp *CsvProcessor) Save() error {
	fmt.Println("Saving processed csv data")
	return nil
}
