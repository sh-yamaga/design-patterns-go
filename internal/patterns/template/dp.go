package template

// DataProcessor defines the template method and steps
type DataProcessor interface {
	LoadData() error
	ProcessData() error
	SaveData() error
	Execute() error // Template method
}

// BaseProcessor provides a default implementation of Execute
type BaseProcessor struct {
	DataProcessor
}

// Execute defines the skeleton of the algorithm
func (bp *BaseProcessor) Execute() error {
	if err := bp.LoadData(); err != nil {
		return err
	}
	if err := bp.ProcessData(); err != nil {
		return err
	}
	if err := bp.SaveData(); err != nil {
		return err
	}

	return nil
}
