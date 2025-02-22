package template

import "fmt"

// IDataProcessor defines the template method and steps
type IDataProcessor interface {
	LoadData() error
	ProcessData() error
	SaveData() error
	Execute() error // Template method
}

// BaseProcessor provides a default implementation of Execute
type BaseProcessor struct {
	IDataProcessor
}

// This is a placeholder comment for the LoadData method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) LoadData() error {
	return fmt.Errorf("LoadData(): should be implemented in concrete struct")
}

// This is a placeholder comment for the ProcessData method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) ProcessData() error {
	return fmt.Errorf("ProcessData(): should be implemented in concrete struct")
}

// This is a placeholder comment for the SaveData method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) SaveData() error {
	return fmt.Errorf("SaveData(): should be implemented in concrete struct")
}

// Execute defines the skeleton of the algorithm
func (bp *BaseProcessor) Execute() error {
	if err := bp.IDataProcessor.LoadData(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := bp.IDataProcessor.ProcessData(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := bp.IDataProcessor.SaveData(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
