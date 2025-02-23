package factory

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/template"

// DataProcessorFactory defines the factory method for creating IDataProcessor instances
type DataProcessorFactory interface {
	CreateCsvProcessor() template.IDataProcessor
}

// ProcessorFactory is a concrete factory for creating CsvProcessor instances
type ProcessorFactory struct{}

// CreateCsvProcessor creates a new CsvProcessor instance
func (f *ProcessorFactory) CreateCsvProcessor() template.IDataProcessor {
	cp := &template.CsvProcessor{}
	cp.BaseProcessor.IDataProcessor = cp

	return cp
}
