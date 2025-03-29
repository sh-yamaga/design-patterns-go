package processor

import "fmt"

func NewXmlProcessor() *processor {
	return &processor{
		process: &xmlProcess{},
	}
}

type xmlProcess struct{}

func (xp *xmlProcess) load() error {
	fmt.Println("Loading data from xml file")
	return nil
}

// handle processes the loaded data
func (xp *xmlProcess) handle() error {
	fmt.Println("Handling xml data")
	return nil
}

// SaveData saves the processed data
func (xp *xmlProcess) save() error {
	fmt.Println("Saving processed xml data")
	return nil
}
