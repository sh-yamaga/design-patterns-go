package processor

import (
	"fmt"
)

func NewCsvProcessor() *processor {
	return &processor{
		process: &csvProcess{},
	}
}

type csvProcess struct{}

// Load loads data from a CSV file
func (cp *csvProcess) load() error {
	fmt.Println("Loading data from csv file")
	return nil
}

// handle processes the loaded data
func (cp *csvProcess) handle() error {
	fmt.Println("Handling csv data")
	return nil
}

// SaveData saves the processed data
func (cp *csvProcess) save() error {
	fmt.Println("Saving processed csv data")
	return nil
}
