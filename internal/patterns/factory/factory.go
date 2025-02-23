package factory

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/factory/processor"

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	CreateCsvProcessor() processor.IDataProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// CreateCsvProcessor creates a new CsvProcessor instance
func (f *ProcessorFactory) CreateCsvProcessor() processor.IDataProcessor {
	return processor.NewCsvProcessor()
}
