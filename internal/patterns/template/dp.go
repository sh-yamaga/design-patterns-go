package template

// DataProcessor defines the template method and steps
type DataProcessor interface {
	LoadData()
	ProcessData()
	SaveData()
	Execute() // Template method
}

// BaseProcessor provides a default implementation of Execute
type BaseProcessor struct {
	DataProcessor
}

// Execute defines the skeleton of the algorithm
func (bp *BaseProcessor) Execute() {
	bp.LoadData()
	bp.ProcessData()
	bp.SaveData()
}
