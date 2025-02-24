package processor

import "fmt"

// XmlProcessor is a concrete implementation of IDataProcessor
type XmlProcessor struct {
	BaseProcessor
}

func (xp *XmlProcessor) Load() error {
	fmt.Println("Loading data from xml file")
	return nil
}

// ProcessData processes the loaded data
func (xp *XmlProcessor) Process() error {
	fmt.Println("Processing xml data")
	return nil
}

// SaveData saves the processed data
func (xp *XmlProcessor) Save() error {
	fmt.Println("Saving processed xml data")
	return nil
}
