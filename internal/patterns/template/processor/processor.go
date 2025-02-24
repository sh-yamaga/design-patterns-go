package processor

import (
	"errors"
)

// IDataProcessor defines the template method and steps
type IDataProcessor interface {
	Load() error
	Process() error
	Save() error
	Execute() error // Template method
}

// BaseProcessor provides a default implementation of Execute
type BaseProcessor struct {
	IDataProcessor
}

// This is a placeholder comment for the LoadData method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) Load() error {
	return errors.New("Load(): should be implemented in concrete struct")
}

// This is a placeholder comment for the Process method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) Process() error {
	return errors.New("Process(): should be implemented in concrete struct")
}

// This is a placeholder comment for the Save method.
// The actual implementation should be provided in the concrete struct.
func (bp *BaseProcessor) Save() error {
	return errors.New("Save(): should be implemented in concrete struct")
}

// Execute defines the skeleton of the algorithm
func (bp *BaseProcessor) Execute() error {
	if err := bp.IDataProcessor.Load(); err != nil {
		return err
	}
	if err := bp.IDataProcessor.Process(); err != nil {
		return err
	}
	if err := bp.IDataProcessor.Save(); err != nil {
		return err
	}

	return nil
}
